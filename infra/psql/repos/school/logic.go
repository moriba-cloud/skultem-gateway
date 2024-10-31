package school

import (
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/psql"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/school"
	"github.com/moriba-cloud/skultem-gateway/infra/psql/repos/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type model struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (m model) Save(args school.Domain) (*ddd.Response[school.Domain], error) {
	model := Model(&args)

	err := m.db.Transaction(func(tx *gorm.DB) error {

		// create school
		if err := tx.Save(&model).Error; err != nil {
			return err
		}

		// create user
		for _, u := range args.Users() {
			userModel := user.Model(&u)
			if err := tx.Save(&userModel).Error; err != nil {
				return err
			}
		}

		// create contact
		for _, phone := range args.Phones() {
			o := ToPhone(phone, args.ID())
			if err := tx.Save(&o).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	m.logger.Info(fmt.Sprintf("saved school with id: %s", model.ID))
	return m.FindById(model.ID)
}

func (m model) FindById(id string) (*ddd.Response[school.Domain], error) {
	var model School
	if err := m.db.Preload("Phones").
		Preload("Users.Role").
		Where(&School{ID: id}).
		First(&model).Error; err != nil {
		return nil, err
	}

	record, err := model.ToDomain()
	if err != nil {
		return nil, err
	}

	return ddd.NewResponse(ddd.ResponseArgs[school.Domain]{
		Record: record,
	}), nil
}

func (m model) Check(key string, value string) (*school.Domain, error) {
	var model School
	query := fmt.Sprintf("upper(%s) = upper(?)", key)

	if err := m.db.Where(query, value).First(&model).Error; err != nil {
		return nil, err
	}

	return model.ToDomain()
}

func (m model) List() (*ddd.Response[core.Option], error) {
	records := make([]*core.Option, 0)
	models := make([]*School, 0)

	m.db.Find(&models)
	for _, o := range models {
		record, err := o.ToDomain()
		if err != nil {
			return nil, err
		}

		records = append(records, record.Option())
	}

	m.logger.Info(fmt.Sprintf("fetch %d schools", len(records)))
	return ddd.NewResponse[core.Option](ddd.ResponseArgs[core.Option]{
		Records: records,
	}), nil
}

func (m model) ListByPage(args ddd.PaginationArgs) (*ddd.Response[school.Domain], error) {
	records := make([]*school.Domain, 0)
	models := make([]*School, 0)
	p := psql.NewPagination[School](psql.PaginationArgs{
		Limit: args.Limit,
		Page:  args.Page,
	})

	m.db.Preload("Phones").
		Scopes(p.Paginate(m.db)).
		Find(&models)

	for _, record := range models {
		d, err := record.ToDomain()
		if err != nil {
			return nil, err
		}

		records = append(records, d)
	}

	m.logger.Info(fmt.Sprintf("fetch %d school by limit: %d and page: %d", len(records), args.Limit, args.Page))
	return ddd.NewResponse(ddd.ResponseArgs[school.Domain]{
		Records: records,
		Pagination: ddd.PaginationArgs{
			Limit: p.Limit(),
			Page:  p.Page(),
			Size:  p.Size(),
			Pages: p.Pages(),
		},
	}), nil
}

func (m model) Remove(args school.Domain) (*school.Domain, error) {
	model := Model(&args)
	m.db.Delete(model)
	m.logger.Info(fmt.Sprintf("remove school by id: %s", model.ID))
	return model.ToDomain()
}

func New(db *gorm.DB, logger *zap.Logger) school.Repo {
	return &model{db: db, logger: logger}
}

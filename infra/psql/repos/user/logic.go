package feature

import (
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/psql"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type model struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (m *model) Save(args user.Domain) (*user.Domain, error) {
	model := Model(&args)
	if err := m.db.Save(model).Error; err != nil {
		return nil, err
	}
	m.logger.Info(fmt.Sprintf("saved user with id: %s", model.ID))
	return model.Domain()
}

func (m *model) Check(phone int) (*user.Domain, error) {
	var model User
	if err := m.db.Where(&User{Phone: phone}).First(&model).Error; err != nil {
		return nil, err
	}
	return model.Domain()
}

func (m *model) FindById(id string) (*user.Domain, error) {
	var model User
	if err := m.db.Where(&User{ID: id}).First(&model).Error; err != nil {
		return nil, err
	}
	return model.Domain()
}

func (m *model) ListByPage(args ddd.PaginationArgs) (*ddd.Response[user.Domain], error) {
	records := make([]*user.Domain, 0)
	models := make([]*User, 0)
	p := psql.NewPagination[User](psql.PaginationArgs{
		Limit: args.Limit,
		Page:  args.Page,
	})

	m.db.Scopes(p.Paginate(m.db)).Find(&models)
	for _, record := range models {
		d, err := record.Domain()
		if err != nil {
			return nil, err
		}

		records = append(records, d)
	}

	m.logger.Info(fmt.Sprintf("fetch %d user by limit: %d and page: %d", len(records), args.Limit, args.Page))
	return ddd.NewResponse(ddd.ResponseArgs[user.Domain]{
		Records: records,
		Pagination: ddd.PaginationArgs{
			Limit: p.Limit(),
			Page:  p.Page(),
			Size:  p.Size(),
			Pages: p.Pages(),
		},
	}), nil
}

func (m *model) List() (*ddd.Response[core.Option], error) {
	records := make([]*core.Option, 0)
	models := make([]*User, 0)

	m.db.Find(&models)
	for _, o := range models {
		record, err := o.Domain()
		if err != nil {
			return nil, err
		}

		records = append(records, record.Option())
	}

	m.logger.Info(fmt.Sprintf("fetch %d users", len(records)))
	return ddd.NewResponse[core.Option](ddd.ResponseArgs[core.Option]{
		Records: records,
	}), nil
}

func (m *model) Remove(args user.Domain) (*user.Domain, error) {
	model := Model(&args)
	m.db.Delete(model)
	m.logger.Info(fmt.Sprintf("remove user by id: %s", model.ID))
	return model.Domain()
}

func New(db *gorm.DB, logger *zap.Logger) user.Repo {
	return &model{db: db, logger: logger}
}

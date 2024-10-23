package role

import (
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/psql"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/role"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type model struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (m *model) Save(args role.Domain) (*role.Domain, error) {
	model := Model(&args)
	if err := m.db.Save(model).Error; err != nil {
		return nil, err
	}
	m.logger.Info(fmt.Sprintf("saved role with id: %s", model.ID))
	return model.Domain()
}

func (m *model) Check(name string) (*role.Domain, error) {
	var model Role
	if err := m.db.Where(&Role{Name: name}).First(&model).Error; err != nil {
		return nil, err
	}
	return model.Domain()
}

func (m *model) FindById(id string) (*role.Domain, error) {
	var model Role
	if err := m.db.Where(&Role{ID: id}).First(&model).Error; err != nil {
		return nil, err
	}
	return model.Domain()
}

func (m *model) ListByPage(args ddd.PaginationArgs) (*ddd.Response[role.Domain], error) {
	records := make([]*role.Domain, 0)
	models := make([]*Role, 0)
	p := psql.NewPagination[Role](psql.PaginationArgs{
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

	m.logger.Info(fmt.Sprintf("fetch %d role by limit: %d and page: %d", len(records), args.Limit, args.Page))
	return ddd.NewResponse(ddd.ResponseArgs[role.Domain]{
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
	models := make([]*Role, 0)

	m.db.Find(&models)
	for _, o := range models {
		record, err := o.Domain()
		if err != nil {
			return nil, err
		}

		records = append(records, &core.Option{
			Label: record.Name(),
			Value: record.ID(),
		})
	}

	m.logger.Info(fmt.Sprintf("fetch %d roles", len(records)))
	return ddd.NewResponse[core.Option](ddd.ResponseArgs[core.Option]{
		Records: records,
	}), nil
}

func (m *model) Remove(args role.Domain) (*role.Domain, error) {
	model := Model(&args)
	m.db.Delete(model)
	m.logger.Info(fmt.Sprintf("remove role by id: %s", model.ID))
	return model.Domain()
}

func New(db *gorm.DB, logger *zap.Logger) role.Repo {
	return &model{db: db, logger: logger}
}

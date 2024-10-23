package feature

import (
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/psql"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/feature"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type model struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (m *model) Save(args feature.Domain) (*feature.Domain, error) {
	model := Model(&args)
	if err := m.db.Save(model).Error; err != nil {
		return nil, err
	}
	m.logger.Info(fmt.Sprintf("saved feature with id: %s", model.ID))
	return model.Domain()
}

func (m *model) Check(name string) (*feature.Domain, error) {
	var model Feature
	if err := m.db.Where(&Feature{Name: name}).First(&model).Error; err != nil {
		return nil, err
	}
	return model.Domain()
}

func (m *model) FindById(id string) (*feature.Domain, error) {
	var model Feature
	if err := m.db.Where(&Feature{ID: id}).First(&model).Error; err != nil {
		return nil, err
	}
	return model.Domain()
}

func (m *model) ListByPage(args ddd.PaginationArgs) (*ddd.Response[feature.Domain], error) {
	records := make([]*feature.Domain, 0)
	models := make([]*Feature, 0)
	p := psql.NewPagination[Feature](psql.PaginationArgs{
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

	m.logger.Info(fmt.Sprintf("fetch %d feature by limit: %d and page: %d", len(records), args.Limit, args.Page))
	return ddd.NewResponse(ddd.ResponseArgs[feature.Domain]{
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
	models := make([]*Feature, 0)

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

	m.logger.Info(fmt.Sprintf("fetch %d features", len(records)))
	return ddd.NewResponse[core.Option](ddd.ResponseArgs[core.Option]{
		Records: records,
	}), nil
}

func (m *model) Remove(args feature.Domain) (*feature.Domain, error) {
	model := Model(&args)
	m.db.Delete(model)
	m.logger.Info(fmt.Sprintf("remove feature by id: %s", model.ID))
	return model.Domain()
}

func New(db *gorm.DB, logger *zap.Logger) feature.Repo {
	return &model{db: db, logger: logger}
}

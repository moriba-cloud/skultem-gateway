package year

import (
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-build/ose/ddd/psql"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/year"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type model struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (m *model) Save(args year.Domain) (*year.Domain, error) {
	model := Model(&args)
	if err := m.db.Save(model).Error; err != nil {
		return nil, err
	}
	m.logger.Info(fmt.Sprintf("saved academic year with id: %s", model.ID))
	return model.Domain()
}

func (m *model) OneById(id string) (*year.Domain, error) {
	var model Year
	if err := m.db.Where(&Year{ID: id}).First(&model).Error; err != nil {
		return nil, err
	}
	m.logger.Info(fmt.Sprintf("fetched academic year with id: %s", id))
	return model.Domain()
}

func (m *model) Check(start int32, end int32) bool {
	var model Year
	if err := m.db.Where(&Year{Start: start, End: end}).First(&model).Error; err != nil {
		return false
	}
	return true
}

func (m *model) ListByPage(args ddd.PaginationArgs) (*ddd.Response[year.Domain], error) {
	records := make([]*year.Domain, 0)
	models := make([]*Year, 0)
	p := psql.NewPagination[Year](psql.PaginationArgs{
		Limit: args.Limit,
		Page:  args.Page,
	})

	m.db.Scopes(p.Paginate(m.db)).Find(&models)
	for _, record := range models {
		vo, err := record.Domain()
		if err != nil {
			return nil, err
		}

		records = append(records, vo)
	}

	m.logger.Info(fmt.Sprintf("fetch %d academic year by limit: %d and page: %d", len(records), args.Limit, args.Page))
	return ddd.NewResponse(ddd.ResponseArgs[year.Domain]{
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
	models := make([]*Year, 0)

	m.db.Find(&models)
	for _, o := range models {
		record, err := o.Domain()
		if err != nil {
			return nil, err
		}

		records = append(records, &core.Option{
			Label: fmt.Sprintf("%s - %s", record.State(), record.End()),
			Value: record.ID(),
		})
	}

	m.logger.Info(fmt.Sprintf("fetch %d academic year", len(records)))
	return ddd.NewResponse[core.Option](ddd.ResponseArgs[core.Option]{
		Records: records,
	}), nil
}

func (m *model) Remove(args *year.Domain) (*year.Domain, error) {
	model := Model(args)
	m.db.Delete(model)
	m.logger.Info(fmt.Sprintf("remove values by id: %s", model.ID))
	return model.Domain()
}

func New(db *gorm.DB, logger *zap.Logger) year.Repo {
	return &model{db: db, logger: logger}
}

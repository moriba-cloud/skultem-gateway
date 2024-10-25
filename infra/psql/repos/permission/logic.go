package permission

import (
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/permission"
	"github.com/moriba-cloud/skultem-gateway/infra/psql/repos/feature"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type model struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (m *model) Save(args []*permission.Domain, role string) (*ddd.Response[permission.Domain], error) {
	err := m.db.Transaction(func(tx *gorm.DB) error {
		for _, arg := range args {
			model := Model(arg, role)
			if err := tx.Save(&model).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	m.logger.Info(fmt.Sprintf("saved permissions %d with id: %s", len(args), role))
	return m.RolePermissions(role)
}

func (m *model) RolePermissions(role string) (*ddd.Response[permission.Domain], error) {
	records := make([]*permission.Domain, 0)
	features := make([]*feature.Feature, 0)
	m.db.Find(&features)

	for _, o := range features {
		var model *Permission
		if err := m.db.Preload("Role").
			Preload("Feature").
			Where("role_id = ? and feature_id = ?", role, o.ID).
			First(&model).
			Error; err != nil {

			featureDomain, err := o.Domain()
			if err != nil {
				return nil, err
			}

			record, err := permission.New(permission.Args{
				Feature: core.Reference{
					Id:    featureDomain.ID(),
					Value: featureDomain.Name(),
				},
				Create:  false,
				Read:    false,
				ReadAll: false,
				Edit:    false,
				Delete:  false,
			})
			records = append(records, record)
		} else {
			record, err := model.Domain()
			if err != nil {
				return nil, err
			}
			records = append(records, record)
		}

	}

	m.logger.Info(fmt.Sprintf("fetch %d permissions by role: %s", len(records), role))
	return ddd.NewResponse(ddd.ResponseArgs[permission.Domain]{
		Records: records,
	}), nil
}

func (m *model) Check(feature string, role string) (*permission.Domain, error) {
	var model Permission
	if err := m.db.Preload("Feature").Preload("Role").Where(&Permission{FeatureId: feature, RoleId: role}).First(&model).Error; err != nil {
		return nil, err
	}

	return model.Domain()
}

func New(db *gorm.DB, logger *zap.Logger) permission.Repo {
	return &model{db: db, logger: logger}
}

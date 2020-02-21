// Package audited is used to log last UpdatedBy and CreatedBy for your models
//
// Github: http://GoTenancy/libs/audited
package audited

import "fmt"

// AuditedModel make Model Auditable, embed `audited.AuditedModel` into your model as anonymous field to make the model auditable
//    type User struct {
//      gorm.Model
//      audited.AuditedModel
//    }
type AuditedModel struct {
	CreatedBy string
	UpdatedBy string
}

// SetCreatedBy set created by
func (model *AuditedModel) SetCreatedBy(createdBy interface{}) {
	model.CreatedBy = fmt.Sprintf("%v", createdBy)
}

// GetCreatedBy get created by
func (model AuditedModel) GetCreatedBy() string {
	return model.CreatedBy
}

// SetUpdatedBy set updated by
func (model *AuditedModel) SetUpdatedBy(updatedBy interface{}) {
	model.UpdatedBy = fmt.Sprintf("%v", updatedBy)
}

// GetUpdatedBy get updated by
func (model AuditedModel) GetUpdatedBy() string {
	return model.UpdatedBy
}

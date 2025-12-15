package mappers

import (
	"softpharos/internal/core/domain/milestone"
	"softpharos/internal/infra/databases/models"
)

func MilestoneToDomain(model *models.MilestoneModel) *milestone.Milestone {
	if model == nil {
		return nil
	}

	return &milestone.Milestone{
		ID:          model.ID,
		ProjectID:   model.ProjectID,
		Project:     ProjectToDomain(model.Project),
		Title:       model.Title,
		Description: model.Description,
		ClassWeek:   model.ClassWeek,
		CreatedAt:   model.CreatedAt,
	}
}

func MilestoneToModel(domain *milestone.Milestone) *models.MilestoneModel {
	if domain == nil {
		return nil
	}

	return &models.MilestoneModel{
		ID:          domain.ID,
		ProjectID:   domain.ProjectID,
		Project:     ProjectToModel(domain.Project),
		Title:       domain.Title,
		Description: domain.Description,
		ClassWeek:   domain.ClassWeek,
		CreatedAt:   domain.CreatedAt,
	}
}

func MilestoneListToDomain(modelList []models.MilestoneModel) []milestone.Milestone {
	domainList := make([]milestone.Milestone, len(modelList))
	for i, model := range modelList {
		domainList[i] = *MilestoneToDomain(&model)
	}
	return domainList
}

package mappers

import (
	"softpharos/internal/core/domain/project"
	"softpharos/internal/infra/databases/models"
)

func ProjectToDomain(model *models.ProjectModel) *project.Project {
	if model == nil {
		return nil
	}

	return &project.Project{
		ID:        model.ID,
		Name:      model.Name,
		Objective: model.Objective,
		CreatedBy: model.CreatedBy,
		Creator:   UserToDomain(model.Creator),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func ProjectToModel(domain *project.Project) *models.ProjectModel {
	if domain == nil {
		return nil
	}

	return &models.ProjectModel{
		ID:        domain.ID,
		Name:      domain.Name,
		Objective: domain.Objective,
		CreatedBy: domain.CreatedBy,
		Creator:   UserToModel(domain.Creator),
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ProjectListToDomain(modelList []models.ProjectModel) []project.Project {
	domainList := make([]project.Project, len(modelList))
	for i, model := range modelList {
		domainList[i] = *ProjectToDomain(&model)
	}
	return domainList
}

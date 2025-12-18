package mappers

import (
	"softpharos/internal/core/domain/project_member"
	"softpharos/internal/infra/databases/models"
)

func ProjectMemberToDomain(model *models.ProjectMemberModel) *project_member.ProjectMember {
	if model == nil {
		return nil
	}

	return &project_member.ProjectMember{
		ID:        model.ID,
		ProjectID: model.ProjectID,
		Project:   ProjectToDomain(model.Project),
		UserID:    model.UserID,
		User:      UserToDomain(model.User),
		Role:      model.Role,
		JoinedAt:  model.JoinedAt,
	}
}

func ProjectMemberToModel(domain *project_member.ProjectMember) *models.ProjectMemberModel {
	if domain == nil {
		return nil
	}

	return &models.ProjectMemberModel{
		ID:        domain.ID,
		ProjectID: domain.ProjectID,
		Project:   ProjectToModel(domain.Project),
		UserID:    domain.UserID,
		User:      UserToModel(domain.User),
		Role:      domain.Role,
		JoinedAt:  domain.JoinedAt,
	}
}

func ProjectMemberListToDomain(modelList []models.ProjectMemberModel) []project_member.ProjectMember {
	domainList := make([]project_member.ProjectMember, len(modelList))
	for i, model := range modelList {
		domainList[i] = *ProjectMemberToDomain(&model)
	}
	return domainList
}

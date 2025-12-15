package mappers

import (
	"softpharos/internal/core/domain/comment"
	"softpharos/internal/infra/databases/models"
)

func CommentToDomain(model *models.CommentModel) *comment.Comment {
	if model == nil {
		return nil
	}

	return &comment.Comment{
		ID:          model.ID,
		MilestoneID: model.MilestoneID,
		Milestone:   MilestoneToDomain(model.Milestone),
		UserID:      model.UserID,
		User:        UserToDomain(model.User),
		Content:     model.Content,
		CreatedAt:   model.CreatedAt,
	}
}

func CommentToModel(domain *comment.Comment) *models.CommentModel {
	if domain == nil {
		return nil
	}

	return &models.CommentModel{
		ID:          domain.ID,
		MilestoneID: domain.MilestoneID,
		Milestone:   MilestoneToModel(domain.Milestone),
		UserID:      domain.UserID,
		User:        UserToModel(domain.User),
		Content:     domain.Content,
		CreatedAt:   domain.CreatedAt,
	}
}

func CommentListToDomain(modelList []models.CommentModel) []comment.Comment {
	domainList := make([]comment.Comment, len(modelList))
	for i, model := range modelList {
		domainList[i] = *CommentToDomain(&model)
	}
	return domainList
}

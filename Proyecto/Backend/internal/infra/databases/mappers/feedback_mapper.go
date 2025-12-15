package mappers

import (
	"softpharos/internal/core/domain/feedback"
	"softpharos/internal/infra/databases/models"
)

func FeedbackToDomain(model *models.FeedbackModel) *feedback.Feedback {
	if model == nil {
		return nil
	}

	return &feedback.Feedback{
		ID:          model.ID,
		MilestoneID: model.MilestoneID,
		Milestone:   MilestoneToDomain(model.Milestone),
		ProfessorID: model.ProfessorID,
		Professor:   UserToDomain(model.Professor),
		Content:     model.Content,
		CreatedAt:   model.CreatedAt,
	}
}

func FeedbackToModel(domain *feedback.Feedback) *models.FeedbackModel {
	if domain == nil {
		return nil
	}

	return &models.FeedbackModel{
		ID:          domain.ID,
		MilestoneID: domain.MilestoneID,
		Milestone:   MilestoneToModel(domain.Milestone),
		ProfessorID: domain.ProfessorID,
		Professor:   UserToModel(domain.Professor),
		Content:     domain.Content,
		CreatedAt:   domain.CreatedAt,
	}
}

func FeedbackListToDomain(modelList []models.FeedbackModel) []feedback.Feedback {
	domainList := make([]feedback.Feedback, len(modelList))
	for i, model := range modelList {
		domainList[i] = *FeedbackToDomain(&model)
	}
	return domainList
}

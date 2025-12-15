package mappers

import (
	"softpharos/internal/core/domain/reaction"
	"softpharos/internal/infra/databases/models"
)

func ReactionToDomain(model *models.ReactionModel) *reaction.Reaction {
	if model == nil {
		return nil
	}

	return &reaction.Reaction{
		ID:          model.ID,
		MilestoneID: model.MilestoneID,
		Milestone:   MilestoneToDomain(model.Milestone),
		UserID:      model.UserID,
		User:        UserToDomain(model.User),
		Type:        model.Type,
		CreatedAt:   model.CreatedAt,
	}
}

func ReactionToModel(domain *reaction.Reaction) *models.ReactionModel {
	if domain == nil {
		return nil
	}

	return &models.ReactionModel{
		ID:          domain.ID,
		MilestoneID: domain.MilestoneID,
		Milestone:   MilestoneToModel(domain.Milestone),
		UserID:      domain.UserID,
		User:        UserToModel(domain.User),
		Type:        domain.Type,
		CreatedAt:   domain.CreatedAt,
	}
}

func ReactionListToDomain(modelList []models.ReactionModel) []reaction.Reaction {
	domainList := make([]reaction.Reaction, len(modelList))
	for i, model := range modelList {
		domainList[i] = *ReactionToDomain(&model)
	}
	return domainList
}

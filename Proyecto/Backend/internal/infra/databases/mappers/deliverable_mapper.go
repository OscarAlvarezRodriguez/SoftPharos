package mappers

import (
	"softpharos/internal/core/domain/deliverable"
	"softpharos/internal/infra/databases/models"
)

func DeliverableToDomain(model *models.DeliverableModel) *deliverable.Deliverable {
	if model == nil {
		return nil
	}

	return &deliverable.Deliverable{
		ID:          model.ID,
		MilestoneID: model.MilestoneID,
		Milestone:   MilestoneToDomain(model.Milestone),
		URL:         model.URL,
		Type:        model.Type,
		CreatedAt:   model.CreatedAt,
	}
}

func DeliverableToModel(domain *deliverable.Deliverable) *models.DeliverableModel {
	if domain == nil {
		return nil
	}

	return &models.DeliverableModel{
		ID:          domain.ID,
		MilestoneID: domain.MilestoneID,
		Milestone:   MilestoneToModel(domain.Milestone),
		URL:         domain.URL,
		Type:        domain.Type,
		CreatedAt:   domain.CreatedAt,
	}
}

func DeliverableListToDomain(modelList []models.DeliverableModel) []deliverable.Deliverable {
	domainList := make([]deliverable.Deliverable, len(modelList))
	for i, model := range modelList {
		domainList[i] = *DeliverableToDomain(&model)
	}
	return domainList
}

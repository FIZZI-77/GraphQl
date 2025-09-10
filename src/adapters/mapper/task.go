package mapper

import (
	"GraphQL/graph/model"
	modelsService "GraphQL/src/models"
)

func TaskToGraphQlTask(tasks *modelsService.Tasks) *model.Task {
	return &model.Task{
		UserID:      tasks.UserID.String(),
		Title:       tasks.Title,
		Description: tasks.Description,
		Completed:   tasks.Completed,
	}
}

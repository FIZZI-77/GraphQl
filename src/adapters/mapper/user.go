package mapper

import (
	"GraphQL/graph/model"
	modelsService "GraphQL/src/models"
)

func UserToGraphQLUser(user *modelsService.Users) *model.User {
	return &model.User{
		ID:    user.UserID.String(),
		Name:  user.Name,
		Email: user.Email,
	}
}

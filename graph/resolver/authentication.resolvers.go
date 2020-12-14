package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/graph/graph_model"
	"example/models"
	"example/utils/convert"
	"example/utils/encrypt"
	jwt_service "example/utils/jwt"
	"example/utils/mysql_util"
)

func (r *mutationResolver) Register(ctx context.Context, input graph_model.Register) (*graph_model.User, error) {
	hashPassword, err := encrypt.HashPassword(input.Password)
	if err != nil {
		return &graph_model.User{}, err
	}
	newUser := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hashPassword,
	}
	if err := mysql_util.DB.Create(&newUser).Error; err != nil {
		return &graph_model.User{}, err
	}
	token := jwt_service.GenerateToken(newUser.ID)
	createdUser := &graph_model.User{
		Token:      &token,
		ID:         &newUser.ID,
		FirstName:  &newUser.FirstName,
		LastName:   &newUser.LastName,
		Email:      &newUser.Email,
		FacebookID: &newUser.FacebookId,
		GoogleID:   &newUser.GoogleId,
		Avatar:     &newUser.Avatar,
		CreatedAt:  convert.TimeToString(newUser.CreatedAt),
		UpdatedAt:  convert.TimeToString(newUser.UpdatedAt),
	}
	return createdUser, nil
}

func (r *mutationResolver) Login(ctx context.Context, input graph_model.Login) (*graph_model.User, error) {
	user := models.User{}
	err := mysql_util.DB.Model(&models.User{}).Where("email = ?", input.Email).First(&user).Error
	if err != nil {
		return &graph_model.User{}, err
	}
	if err = encrypt.CheckPassword(input.Password, user.Password); err != nil {
		return &graph_model.User{}, err
	}
	token := jwt_service.GenerateToken(user.ID)
	result := &graph_model.User{
		Token:      &token,
		ID:         &user.ID,
		FirstName:  &user.FirstName,
		LastName:   &user.LastName,
		Email:      &user.Email,
		FacebookID: &user.FacebookId,
		GoogleID:   &user.GoogleId,
		Avatar:     &user.Avatar,
		CreatedAt:  convert.TimeToString(user.CreatedAt),
		UpdatedAt:  convert.TimeToString(user.UpdatedAt),
	}
	return result, nil
}

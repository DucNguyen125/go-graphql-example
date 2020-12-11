package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/graph/custom_model"
	"example/graph/graph_model"
	"example/models"
	"example/utils/convert"
	"example/utils/mysql_util"
)

func (r *mutationResolver) OrderCreate(ctx context.Context, input graph_model.CreateOrder) (*graph_model.Order, error) {
	newOrder := models.Order{
		OrderCode:   input.OrderCode,
		OrderType:   input.OrderType,
		Products:    input.Products,
		OrderStatus: input.OrderStatus,
		Quantity:    input.Quantity,
		TotalPrice:  input.TotalPrice,
	}
	if err := mysql_util.DB.Create(&newOrder).Error; err != nil {
		return &graph_model.Order{}, err
	}
	createdOrder := &graph_model.Order{
		ID:          &newOrder.ID,
		OrderCode:   &newOrder.OrderCode,
		OrderType:   &newOrder.OrderType,
		Products:    &newOrder.Products,
		OrderStatus: &newOrder.OrderStatus,
		Quantity:    &newOrder.Quantity,
		TotalPrice:  &newOrder.TotalPrice,
		CreatedAt:   convert.TimeToString(newOrder.CreatedAt),
		UpdatedAt:   convert.TimeToString(newOrder.UpdatedAt),
	}
	return createdOrder, nil
}

func (r *mutationResolver) OrderUpdate(ctx context.Context, input custom_model.UpdateOrder) (*graph_model.Order, error) {
	order := &graph_model.Order{}
	err := mysql_util.DB.Updates(models.Order{
		ID:          input.ID,
		OrderCode:   input.OrderCode,
		OrderType:   input.OrderType,
		Products:    input.Products,
		OrderStatus: input.OrderStatus,
		Quantity:    input.Quantity,
		TotalPrice:  input.TotalPrice,
	}).First(order, input.ID).Error
	if err != nil {
		return &graph_model.Order{}, err
	}
	return order, nil
}

func (r *mutationResolver) OrderDelete(ctx context.Context, input graph_model.DeleteOrder) (*int, error) {
	order := graph_model.DeleteOrder{}
	err := mysql_util.DB.Model(&models.Order{}).Delete(&order, input.ID).Error
	if err != nil {
		return nil, err
	}
	return &input.ID, nil
}

func (r *queryResolver) Order(ctx context.Context, id int) (*graph_model.Order, error) {
	order := &graph_model.Order{}
	err := mysql_util.DB.Model(&models.Order{}).First(order, id).Error
	if err != nil {
		return &graph_model.Order{}, err
	}
	return order, nil
}

func (r *queryResolver) Orders(ctx context.Context, filter string, limit int, page int) ([]graph_model.Order, error) {
	listOrder := []graph_model.Order{}
	skip := (page - 1) * limit
	mysql_util.DB.Model(&models.Order{}).Offset(skip).Limit(limit).Find(&listOrder)
	return listOrder, nil
}

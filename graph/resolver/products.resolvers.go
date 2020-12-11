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

func (r *mutationResolver) ProductCreate(ctx context.Context, input graph_model.CreateProduct) (*graph_model.Product, error) {
	newProduct := models.Product{
		ProductCode: input.ProductCode,
		ProductName: input.ProductName,
		Price:       input.Price,
	}
	if err := mysql_util.DB.Create(&newProduct).Error; err != nil {
		return &graph_model.Product{}, err
	}
	createdProduct := &graph_model.Product{
		ID:          &newProduct.ID,
		ProductCode: &newProduct.ProductCode,
		ProductName: &newProduct.ProductName,
		Price:       &newProduct.Price,
		CreatedAt:   convert.TimeToString(newProduct.CreatedAt),
		UpdatedAt:   convert.TimeToString(newProduct.UpdatedAt),
	}
	return createdProduct, nil
}

func (r *mutationResolver) ProductUpdate(ctx context.Context, input custom_model.UpdateProduct) (*graph_model.Product, error) {
	order := &graph_model.Product{}
	err := mysql_util.DB.Updates(models.Product{
		ID:          input.ID,
		ProductCode: input.ProductCode,
		ProductName: input.ProductName,
		Price:       input.Price,
	}).First(order, input.ID).Error
	if err != nil {
		return &graph_model.Product{}, err
	}
	return order, nil
}

func (r *mutationResolver) ProductDelete(ctx context.Context, input graph_model.DeleteProduct) (*int, error) {
	order := graph_model.DeleteProduct{}
	err := mysql_util.DB.Model(&models.Product{}).Delete(&order, input.ID).Error
	if err != nil {
		return nil, err
	}
	return &input.ID, nil
}

func (r *queryResolver) Product(ctx context.Context, id int) (*graph_model.Product, error) {
	order := &graph_model.Product{}
	err := mysql_util.DB.Model(&models.Product{}).First(order, id).Error
	if err != nil {
		return &graph_model.Product{}, err
	}
	return order, nil
}

func (r *queryResolver) Products(ctx context.Context, filter string, limit int, page int) ([]graph_model.Product, error) {
	listProduct := []graph_model.Product{}
	skip := (page - 1) * limit
	mysql_util.DB.Model(&models.Product{}).Offset(skip).Limit(limit).Find(&listProduct)
	return listProduct, nil
}

package storage

import (
	"carsharing/internal/entity/car/model"
	"context"
)

type Storage interface {
	Create(ctx context.Context, model *model.Car) error
	FindAll(ctx context.Context) (car []model.Car, err error)
	FindOne(ctx context.Context, id string) (car model.Car, err error)
	Update(ctx context.Context, car *model.Car) error
	Delete(ctx context.Context, id string) error
}

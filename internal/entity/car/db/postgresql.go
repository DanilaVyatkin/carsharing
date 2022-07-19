package db

import (
	"carsharing/internal/entity/car/model"
	"carsharing/internal/entity/car/storage"
	"carsharing/pkg/postgres"
	"context"
)

type db struct {
	client *postgres.Client
}

func NewDB(client *postgres.Client) storage.Storage {
	return &db{client: client}
}

func (d *db) Create(ctx context.Context, model *model.Car) error {
	return nil
}

func (d *db) FindAll(ctx context.Context) (car []model.Car, err error) {
	return nil, nil
}

func (d *db) FindOne(ctx context.Context, id string) (car model.Car, err error) {
	return model.Car{}, nil
}

func (d *db) Update(ctx context.Context, car *model.Car) error {
	return nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	return nil
}

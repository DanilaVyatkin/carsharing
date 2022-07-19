package postgres

import (
	"carsharing/internal/config"
	"carsharing/internal/entity/car/model"
	model2 "carsharing/internal/entity/user/model"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func NewClient(logger *zap.SugaredLogger, cfg *config.Config) *Client {
	c := Client{logger, cfg}
	c.RegisterClient()

	return &c
}

func (c *Client) RegisterClient() {
	c.logger.Info("Create new client postgres")
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		c.cfg.Host, c.cfg.User, c.cfg.Password, c.cfg.DBname, c.cfg.Port)

	open, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		c.logger.Fatal(err)
	}

	db := open.Debug().AutoMigrate(&model.Car{}, &model2.User{})
	c.logger.Info(db, dbURL)
}

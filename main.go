package main

import (
	"context"
	"first/config"
	"first/handlers"
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	c   *mongo.Client
	db  *mongo.Database
	col *mongo.Collection
	cfg config.Properties
)

func init() {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Configuration cannot be read :%v", err)
	}

	connectURI := fmt.Sprintf("mongodb://%s:%s", cfg.DBHost, cfg.DBPort)
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectURI))
	if err != nil {
		log.Fatalf("Unable to connect to database : %v", err)
	}
	db = c.Database(cfg.DBName)
	col = db.Collection(cfg.CollectionName)

}

func main() {
	e := echo.New()
	h := handlers.MobileHandler{Col: col}

	e.POST("api/v1/mobiles", h.CreateMobiles)
	e.GET("api/v1/mobiles/:id", h.GetMobile)
	e.DELETE("api/v1/mobiles/:id", h.DeleteMobile)
	e.PUT("api/v1/mobiles/:id", h.UpdateMobile)
	e.Logger.Infof("Listerning on %s:%s", cfg.Host, cfg.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)))

}

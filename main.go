package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"social-todo-list/middleware"
	ginitem "social-todo-list/module/item/transport/gin"
)

func main() {
	dsn := os.Getenv("DB_CONN")
	//systemSecret := os.Getenv("SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug()

	log.Println("DB Connection:", db)
	//////////////////////////
	//authStore := storage.NewSQLStore(db)
	//tokenProvider := jwt.NewTokenJWTProvider("jwt", systemSecret)
	//middlewareAuth := middleware.RequiredAuth(authStore, tokenProvider)

	r := gin.Default()

	r.Use(middleware.Recover())

	r.Static("/static", "./static")

	r.Static("/static", "./static")

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ginitem.ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.PATCH("/:id", ginitem.UpdateItem(db))
			items.DELETE("/:id", ginitem.DeleteItem(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(":3000"); err != nil {
		log.Fatalln(err)
	}
}

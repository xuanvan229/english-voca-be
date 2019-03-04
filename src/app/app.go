package app 

import (
	// "github.com/gorilla/mux"
	// "encoding/json"
	"fmt"
	// "english-vocab/config"
	// "english-vocab/resource/page"
	"english-vocab/resource/vocab"
	"github.com/gin-gonic/gin"
)
type App struct {
	Router *gin.Engine
	// DB *config.PostGresConfig
}

// type User struct {
// 	Username string
// 	Password string
// }

func (a *App) Initialize() {
	// a.DB = config
	a.Router = gin.New()
	a.setRouters()
}

func AuthRequired() gin.HandlerFunc {
	return func( c *gin.Context ) {
		fmt.Println("middle ware")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func (app *App) setRouters() {
	app.Router.Use(gin.Logger())
	app.Router.Use(gin.Recovery())

	api := app.Router.Group("/api")
	{
		api.Use(AuthRequired())
		
		api.POST("/vocab",vocab.CreateVocab)
		api.OPTIONS("/vocab",vocab.CreateVocab)
		api.OPTIONS("/vocab/:id",vocab.DeleteVocab)
		api.DELETE("/vocab/:id",vocab.DeleteVocab)
		api.GET("/vocabs",vocab.GetVocabs)
	}
}

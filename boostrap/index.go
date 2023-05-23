package boostrap

import (
	"go-gin-api/config/appConfig"
	"go-gin-api/database"
	"go-gin-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func BoostrapApp(){
  database.ConnectDB()
  database.RunMigration()
  App := gin.Default()
  
  corsConfig := cors.DefaultConfig()
  corsConfig.AllowOrigins = []string{"*"}
  App.Use(cors.New(corsConfig))  

  routes.SetupRoute(App)
  
  App.Run(appconfig.PORT)
}


package routes

import (
	"go-gin-api/controllers/userController"

	"github.com/gin-gonic/gin"
)

func SetupRoute(app *gin.Engine){
  route := app
    route.GET("/user", userController.Reads)
    route.GET("/user/:id", userController.Read)
    route.POST("/create", userController.Create)
    route.PUT("/update/:id", userController.Update)
    route.DELETE("/delete/:id", userController.Delete)

  //route.GET("/book", bookController.GetAllBook)
 
}

package router

import (
	"go-test/api/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewServer(carController *controller.CarController) *gin.Engine {
	service := gin.Default()
	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})
	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")

	router.GET("car", carController.GetAllCar)
	router.POST("car", carController.CreateCars)
	router.DELETE("car/:id", carController.DeleteCarById)
	router.PUT("car/:id", carController.UpdateCarsById)

	return service
}

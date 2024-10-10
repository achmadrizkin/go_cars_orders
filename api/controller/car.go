package controller

import (
	"go-test/domain"
	"go-test/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CarController struct {
	carUseCase domain.CarUseCase
}

func NewCarController(carUseCase domain.CarUseCase) *CarController {
	return &CarController{carUseCase}
}

func (h *CarController) CreateCars(c *gin.Context) {
	var carRequest model.CarsRequest

	if err := c.ShouldBindJSON(&carRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			StatusCode: 400,
			Message:    err.Error(),
		})
		return
	}

	_, errResponse := h.carUseCase.CreateCars(carRequest)
	if errResponse != nil {
		c.AbortWithStatusJSON(500, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    errResponse.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Insert data cars success",
	})
}

func (h *CarController) UpdateCarsById(c *gin.Context) {
	var carRequest model.CarsRequest
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	if err := c.ShouldBindJSON(&carRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			StatusCode: 400,
			Message:    err.Error(),
		})
		return
	}

	_, errResponse := h.carUseCase.UpdateCarsById(carRequest, idInt)
	if errResponse != nil {
		c.AbortWithStatusJSON(500, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    errResponse.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Update data cars by " + id + " success!",
	})
}

func (h *CarController) GetAllCar(c *gin.Context) {
	var carListResponse []model.CarsResponse

	carList, err := h.carUseCase.GetAllCars()
	if err != nil {
		c.AbortWithStatusJSON(500, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	//
	HandlerCarsResponse(c, convertCarListToResponse(carList, carListResponse))
}

func (h *CarController) DeleteCarById(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			StatusCode: 400,
			Message:    "need /chapter/:id",
		})

		// gunakan return untuk tidak melanjutkan yang dibawah
		return
	}

	err := h.carUseCase.DeleteCarsById(idInt)
	if err != nil {
		c.AbortWithStatusJSON(500, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Delete data cars by " + id + " is success!",
	})
}

func convertCarListToResponse(carList []model.Cars, carListResponse []model.CarsResponse) []model.CarsResponse {
	for _, b := range carList {
		carListResponsee := converToAllCarsResponse(b)
		carListResponse = append(carListResponse, carListResponsee)
	}
	return carListResponse
}

func converToAllCarsResponse(b model.Cars) model.CarsResponse {
	return model.CarsResponse{
		CarId:     b.CarId,
		CarName:   b.CarName,
		DayRate:   b.DayRate,
		MonthRate: b.MonthRate,
		Image:     b.Image,
	}
}

func HandlerCarsResponse(c *gin.Context, carsResponse []model.CarsResponse) {
	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Get data success",
		Data:       carsResponse,
	})
}

package routes

import (
	"golang_basic_gin/config"
	"golang_basic_gin/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPositions(c *gin.Context) {
	positions := []models.Position{}
	// config.DB.Find(&positions)
	config.DB.Preload("Department").Find(&positions)

	GetPositionResponse := []models.GetPositionResponse{}

	for _, p := range positions {
		pos := models.GetPositionResponse{
			ID:   p.ID,
			Name: p.Name,
			Code: p.Code,
			Department: models.DepartmentResponse{
				Name: p.Department.Name,
				Code: p.Department.Code,
			},
		}

		GetPositionResponse = append(GetPositionResponse, pos)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get position",
		// "data":    positions,
		"data": GetPositionResponse,
	})
}

func GetPositionByID(c *gin.Context) {
	id := c.Param("id")
	var position models.Position

	data := config.DB.Preload("Department").First(&position, "id = ?", id)
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	GetPositionResponse := models.GetPositionResponse{
		ID:   position.ID,
		Name: position.Name,
		Code: position.Code,
		Department: models.DepartmentResponse{
			Name: position.Department.Name,
			Code: position.Department.Code,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get position",
		"data":    GetPositionResponse,
	})
}

func Postpositions(c *gin.Context) {
	var position models.Position
	c.BindJSON(&position)

	insert := config.DB.Create(&position)
	if insert.Error != nil {
		log.Printf(insert.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": insert.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{ // 201
		"data":    position,
		"message": "success post position",
	})
}

func Putpositions(c *gin.Context) {
	id := c.Param("id")
	var position models.Position

	data := config.DB.First(&position, "id = ?", id)
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	// cara json
	c.BindJSON(&position)

	update := config.DB.Model(&position).Where("id = ?", id).Updates(&position)
	if update.Error != nil {
		log.Printf(update.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": update.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
		"data":    position,
	})
}

func Deletepositions(c *gin.Context) {
	id := c.Param("id")
	var position models.Position

	data := config.DB.First(&position, "id = ?", id)
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	config.DB.Delete(&position, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "delete success",
	})
}

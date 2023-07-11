package routes

import (
	"golang_basic_gin/config"
	"golang_basic_gin/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetDepartments(c *gin.Context) {
	departments := []models.Department{}
	// config.DB.Find(&departments) // department aja
	// config.DB.Preload("Positions").Find(&departments)
	config.DB.Preload(clause.Associations).Find(&departments)

	GetDepartmentResponse := []models.GetDepartmentResponse{}

	for _, d := range departments {
		positions := []models.PositionResponse{}
		for _, p := range d.Positions {
			pos := models.PositionResponse{
				ID:   p.ID,
				Name: p.Name,
				Code: p.Code,
			}

			positions = append(positions, pos)
		}

		dept := models.GetDepartmentResponse{
			ID:        d.ID,
			Name:      d.Name,
			Code:      d.Code,
			Positions: positions,
		}

		GetDepartmentResponse = append(GetDepartmentResponse, dept)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get department",
		// "data":    departments,
		"data": GetDepartmentResponse,
	})
}

func GetDepartmentByID(c *gin.Context) {
	id := c.Param("id")
	var department models.Department
	data := config.DB.Preload("Positions").First(&department, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	positions := []models.PositionResponse{}
	for _, p := range department.Positions {
		pos := models.PositionResponse{
			ID:   p.ID,
			Name: p.Name,
			Code: p.Code,
		}
		positions = append(positions, pos)
	}

	GetDepartmentResponse := models.GetDepartmentResponse{
		ID:        department.ID,
		Name:      department.Name,
		Code:      department.Code,
		Positions: positions,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get department",
		"data":    GetDepartmentResponse,
	})
}

func PostDepartments(c *gin.Context) {
	// c.PostForm pake x-www-form-urlencoded
	// department := models.Department{
	// 	Name: c.PostForm("name"),
	// 	Code: c.PostForm("code"),
	// }

	// cara post dengan json
	var department models.Department
	c.BindJSON(&department)

	config.DB.Create(&department)

	c.JSON(http.StatusCreated, gin.H{ // 201
		"data":    department,
		"message": "success post department",
	})
}

func PutDepartments(c *gin.Context) {
	id := c.Param("id")
	var department models.Department

	data := config.DB.First(&department, "id = ?", id)
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	// c.PostForm pake x-www-form-urlencoded
	// config.DB.Model(&department).Updates(models.Department{
	// 	Name: c.PostForm("name"),
	// 	Code: c.PostForm("code"),
	// })

	// cara json
	c.BindJSON(&department)
	config.DB.Model(&department).Where("id = ?", id).Updates(&department)

	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
		"data":    department,
	})
}

func DeleteDepartments(c *gin.Context) {
	id := c.Param("id")
	var department models.Department

	data := config.DB.First(&department, "id = ?", id)
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		return
	}

	config.DB.Delete(&department, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "delete success",
	})
}

package admin

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"point-of-sale/app/model"
	"point-of-sale/config"
	"point-of-sale/utils/dto"
	"point-of-sale/utils/res"
)

func IndexCategory(c echo.Context) error {
	var category []model.Category
	if err := config.Db.Find(&category).Error; err != nil {
		format := res.Response(http.StatusInternalServerError, "error", "error retried data", err.Error())
		return c.JSON(http.StatusInternalServerError, format)
	}
	categories := res.TransformCategory(category)
	format := res.Response(http.StatusOK, "success", "successfully retried data", categories)
	return c.JSON(http.StatusOK, format)
}

func CreateCategory(c echo.Context) error {
	request := dto.CreateCategoryRequest{}
	if err := c.Bind(&request); err != nil {
		format := res.Response(http.StatusInternalServerError, "error", "error request body", err.Error())
		return c.JSON(http.StatusInternalServerError, format)
	}

	category := model.Category{
		Name: request.Name,
	}

	if err := config.Db.Create(&category).Error; err != nil {
		format := res.Response(http.StatusInternalServerError, "error", "error create data", err.Error())
		return c.JSON(http.StatusInternalServerError, format)
	}
	format := res.Response(http.StatusCreated, "success", "Category created successfully", category)
	return c.JSON(http.StatusCreated, format)
}

func DeleteCategory(c echo.Context) error {
	id := c.QueryParam("id")

	if err := config.Db.Delete(&model.Category{}, id).Error; err != nil {
		format := res.Response(http.StatusInternalServerError, "error", "error delete data", err.Error())
		return c.JSON(http.StatusInternalServerError, format)
	}

	format := res.Response(http.StatusOK, "success", "CategoryID deleted successfully", nil)
	return c.JSON(http.StatusOK, format)
}

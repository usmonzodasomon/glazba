package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usmonzodasomon/glazba/logger"
	"github.com/usmonzodasomon/glazba/models"
)

type CategoryData struct {
	Name string `json:"name" binding:"required"`
}

func (h *handler) CreateCategory(c *gin.Context) {
	logger.GetLogger().Info("Creating category")
	var input CategoryData

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var category models.Category
	category.Name = input.Name

	id, err := h.services.CreateCategory(&category)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"id":      id,
	})
	logger.GetLogger().Infof("Category with Id %v created succesfully", id)
}

func (h *handler) ReadCategoryByName(c *gin.Context) {
	categoryName := c.Param("category")

	logger.GetLogger().Infof("Reading Category with name %s", categoryName)

	category, err := h.services.ReadCategoryByName(categoryName)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    category,
	})
	logger.GetLogger().Infof("Category with name %s readed succesfully", categoryName)
}

func (h *handler) ReadCategory(c *gin.Context) {
	logger.GetLogger().Info("Reading Category")

	var categories []models.Category
	if err := h.services.ReadCategory(&categories); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    categories,
	})
	logger.GetLogger().Info("Categories read succesfully")

}

func (h *handler) UpdateCategory(c *gin.Context) {
	logger.GetLogger().Info("Updating Category")
	categoryName := c.Param("category")

	var category CategoryData
	if err := c.BindJSON(&category); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateCategory(categoryName, category.Name); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
	logger.GetLogger().Infof("Category with name %s updated succesfully", categoryName)
}

func (h *handler) DeleteCategory(c *gin.Context) {
	logger.GetLogger().Info("Deletating Category")
	categoryName := c.Param("id")

	if err := h.services.DeleteCategory(categoryName); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
	logger.GetLogger().Infof("Category with name %s deleted succesfully", categoryName)
}

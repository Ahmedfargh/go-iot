package controllers

import (
	"net/http"
	"strconv"

	Autoloader "go-dashboard/internals/AutoLoader"
	models "go-dashboard/internals/Models"
	crudServices "go-dashboard/internals/Services/Crud"

	"github.com/gin-gonic/gin"
)

type SectorController struct {
	sector_service crudServices.SectorService
}

func NewSectorController() SectorController {
	return SectorController{
		sector_service: *Autoloader.Services_providers["sector.crud"].(*crudServices.SectorService),
	}
}

func (ctrl *SectorController) Create(c *gin.Context) {
	var sector models.Sector
	if err := c.ShouldBindJSON(&sector); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := ctrl.sector_service.Create(&sector)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (ctrl *SectorController) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	sectors, total, err := ctrl.sector_service.GetAll(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  sectors,
		"total": total,
		"page":  page,
		"limit": pageSize,
	})
}

func (ctrl *SectorController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	sector, err := ctrl.sector_service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if sector == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sector not found"})
		return
	}

	c.JSON(http.StatusOK, sector)
}

func (ctrl *SectorController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var sector models.Sector
	if err := c.ShouldBindJSON(&sector); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := ctrl.sector_service.Update(uint(id), &sector)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (ctrl *SectorController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	deleted, err := ctrl.sector_service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sector not found or already deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sector deleted successfully"})
}

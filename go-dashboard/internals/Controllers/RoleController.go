package controllers

import (
	"net/http"
	"strconv"

	Autoloader "go-dashboard/internals/AutoLoader"
	interfaces "go-dashboard/internals/Interfaces"
	models "go-dashboard/internals/Models"
	crudServices "go-dashboard/internals/Services/Crud"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleService       crudServices.RoleService
	permissionService interfaces.PermissionInterface
}

func NewRoleController() *RoleController {
	return &RoleController{
		roleService:       *Autoloader.Services_providers["role.crud"].(*crudServices.RoleService),
		permissionService: Autoloader.Services_providers["permission"].(interfaces.PermissionInterface),
	}
}

func (ctrl *RoleController) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "100"))

	roles, total, err := ctrl.roleService.GetAll(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":  roles,
		"total": total,
		"page":  page,
		"limit": pageSize,
	})
}

func (ctrl *RoleController) Create(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := ctrl.roleService.Create(&role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (ctrl *RoleController) AssignPermissions(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	var input struct {
		Permissions []string `json:"permissions" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := ctrl.roleService.GetByID(uint(id))
	if err != nil || role == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	err = ctrl.permissionService.SyncPermissionsForRole(role, input.Permissions...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Permissions synced successfully"})
}

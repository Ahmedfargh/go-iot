package controllers

import (
	"net/http"
	"strconv"

	Autoloader "go-dashboard/internals/AutoLoader"
	interfaces "go-dashboard/internals/Interfaces"
	models "go-dashboard/internals/Models"
	authServices "go-dashboard/internals/Services/Auth"
	crudServices "go-dashboard/internals/Services/Crud"
	utils "go-dashboard/internals/Utils"
	"github.com/gin-gonic/gin"
)

type AdminCrudController struct {
	admin_crud_service crudServices.AdminService
	admin_auth_service authServices.AdminAuthService
	permission_service interfaces.PermissionInterface
}

func NewAdminCrudController() AdminCrudController {
	return AdminCrudController{
		admin_crud_service: *Autoloader.Services_providers["admin.crud"].(*crudServices.AdminService),
		admin_auth_service: *Autoloader.Services_providers["admin.auth"].(*authServices.AdminAuthService),
		permission_service: Autoloader.Services_providers["permission"].(interfaces.PermissionInterface),
	}
}

func (ctrl *AdminCrudController) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.admin_auth_service.Login(loginData.Email, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ctrl *AdminCrudController) Create(c *gin.Context) {
	var admin models.Admin
	if err := c.ShouldBind(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Extract roles and permissions from form
	roles := c.PostFormArray("roles")
	if len(roles) == 0 {
		singleRole := c.PostForm("role")
		if singleRole != "" {
			roles = []string{singleRole}
		}
	}
	permissions := c.PostFormArray("permissions")

	// Handle file upload
	file, err := c.FormFile("avatar")
	if err == nil {
		path, err := utils.SaveUploadedFile(file, "avatars")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save avatar"})
			return
		}
		admin.Avatar = path
	}

	createdAdmin, err := ctrl.admin_crud_service.Create(&admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Sync roles and permissions
	if len(roles) > 0 {
		ctrl.permission_service.SyncRoles(createdAdmin, roles...)
	}
	if len(permissions) > 0 {
		ctrl.permission_service.SyncPermissions(createdAdmin, permissions...)
	}

	c.JSON(http.StatusCreated, createdAdmin)
}

func (ctrl *AdminCrudController) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	admins, total, err := ctrl.admin_crud_service.GetAll(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  admins,
		"total": total,
		"page":  page,
		"limit": pageSize,
	})
}

func (ctrl *AdminCrudController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	admin, err := ctrl.admin_crud_service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if admin == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	c.JSON(http.StatusOK, admin)
}

func (ctrl *AdminCrudController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var admin models.Admin
	admin.Email = c.PostForm("email")
	admin.Name = c.PostForm("name")
	admin.Phone = c.PostForm("phone")

	// Extract roles and permissions from form
	roles := c.PostFormArray("roles")
	if len(roles) == 0 {
		singleRole := c.PostForm("role")
		if singleRole != "" {
			roles = []string{singleRole}
		}
	}
	permissions := c.PostFormArray("permissions")

	// Handle file upload
	file, err := c.FormFile("avatar")
	if err == nil {
		path, err := utils.SaveUploadedFile(file, "avatars")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save avatar"})
			return
		}
		admin.Avatar = path
	}
	updatedAdmin, err := ctrl.admin_crud_service.Update(uint(id), &admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Sync roles and permissions
	if len(roles) > 0 {
		ctrl.permission_service.SyncRoles(updatedAdmin, roles...)
	}
	if len(permissions) > 0 {
		ctrl.permission_service.SyncPermissions(updatedAdmin, permissions...)
	}

	c.JSON(http.StatusOK, updatedAdmin)
}

func (ctrl *AdminCrudController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	deleted, err := ctrl.admin_crud_service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found or already deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin deleted successfully"})
}
func (controller *AdminCrudController) index() {

}

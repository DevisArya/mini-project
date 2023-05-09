package repository

import (
	"miniproject/config"
	"miniproject/models"
	"miniproject/utils"
	"net/http"

	"github.com/labstack/echo"
)

type IAdminService interface {
	CreateAdmin(admin *models.Admin) error
	GetAdmin(id int) (error, interface{})
	GetAdminEmail(email string) error
	GetAdmins() (error, interface{})
	DeleteAdmin(id int) error
	UpdateAdmin(dataUpdate *models.Admin, id int) error
	LoginAdmin(admin *models.Admin) error
}

type AdminRepository struct {
	Func IAdminService
}

var adminRepository IAdminService

func init() {
	bg := &AdminRepository{}
	bg.Func = bg

	adminRepository = bg
}
func GetAdminRepository() IAdminService {
	return adminRepository
}
func SetAdminRepository(ur IAdminService) {
	adminRepository = ur
}

func (u *AdminRepository) CreateAdmin(admin *models.Admin) error {
	admin.Role = true

	password := admin.Password

	hash, err := utils.HashPassword(password)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	admin.Password = hash
	if err := config.DB.Save(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (u *AdminRepository) GetAdmin(id int) (err error, res interface{}) {
	var admin models.Admin
	if err := config.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "admin not found",
		}), nil
	}
	result := models.AdminResponse{
		Id:      admin.Id,
		Name:    admin.Name,
		Address: admin.Address,
		Phone:   admin.Phone,
		Email:   admin.Email,
		Role:    admin.Role,
	}
	return nil, result
}
func (u *AdminRepository) GetAdminEmail(email string) (err error) {
	var admin models.Admin
	if err := config.DB.Where("email = ?", email).First(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "admin not found",
		})
	}
	return nil
}

func (u *AdminRepository) GetAdmins() (err error, res interface{}) {
	var admins []models.Admin

	if err := config.DB.Find(&admins).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()), nil
	}

	var adminResponse []models.AdminResponse

	for _, val := range admins {
		result := models.AdminResponse{
			Id:      val.Id,
			Name:    val.Name,
			Address: val.Address,
			Phone:   val.Phone,
			Email:   val.Email,
			Role:    val.Role,
		}
		adminResponse = append(adminResponse, result)
	}
	return nil, adminResponse
}

func (u *AdminRepository) DeleteAdmin(id int) error {
	result := config.DB.Delete(&models.Admin{}, id)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}
	return nil
}

func (u *AdminRepository) UpdateAdmin(updateData *models.Admin, id int) error {
	result := config.DB.Model(&models.Admin{}).Where("id = ?", id).Updates(&updateData)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}
	return nil
}
func (u *AdminRepository) LoginAdmin(admin *models.Admin) error {
	password := admin.Password

	if err := config.DB.Where("email = ?", admin.Email).First(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "user not found",
			"error":   err.Error(),
		})
	}
	if match := utils.CheckPasswordHash(password, admin.Password); match == false {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "mismatch password",
		})
	}
	return nil
}

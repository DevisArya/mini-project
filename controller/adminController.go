package controller

import (
	"miniproject/config"
	md "miniproject/middleware"
	m "miniproject/models"
	u "miniproject/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAdmin(c echo.Context) error {

	var admin m.Admin

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id",
		})
	}

	if err := config.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "admin not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get admin",
		"admin":   admin,
	})

}

func GetAdmins(c echo.Context) error {

	var admins []m.Admin

	if err := config.DB.Find(&admins).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all admins",
		"admins":  admins,
	})
}

func CreateAdmin(c echo.Context) error {

	admin := m.Admin{}
	c.Bind(&admin)
	admin.Role = true

	password := admin.Password

	hash, err := u.HashPassword(password)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	admin.Password = hash
	valid := u.PostAdminValidation(admin)
	if valid != nil {
		return echo.NewHTTPError(http.StatusBadRequest, valid.Error())
	}

	if err := config.DB.Save(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes create new admin",
		"admin":   admin,
	})
}

func DeleteAdmin(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id",
		})
	}

	result := config.DB.Delete(&m.Admin{}, id)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete admin",
	})
}

func UpdateAdmin(c echo.Context) error {
	updateData := m.Admin{}
	c.Bind(&updateData)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	if updateData.Password != "" {
		hash, err := u.HashPassword(updateData.Password)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		updateData.Password = hash
	}

	if updateData.Email != "" {
		valid := u.EmailValidation(updateData.Email)
		if valid != nil {
			return echo.NewHTTPError(http.StatusBadRequest, valid.Error())
		}
	}

	result := config.DB.Model(&m.Admin{}).Where("id = ?", id).Updates(&updateData)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update admin",
	})
}

func LoginAdmin(c echo.Context) error {
	admin := m.Admin{}
	c.Bind(&admin)

	password := admin.Password

	if err := config.DB.Where("email = ?", admin.Email).First(&admin).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "user not found",
			"error":   err.Error(),
		})
	}

	if match := u.CheckPasswordHash(password, admin.Password); match == false {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "mismatch password",
		})
	}

	token, err := md.CreateToken(int(admin.ID), admin.Name, admin.Role)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	adminResponse := m.AdminResponse{ID: int(admin.ID), Name: admin.Name, Email: admin.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"admin":   adminResponse,
	})
}

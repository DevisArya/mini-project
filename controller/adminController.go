package controller

import (
	md "miniproject/middleware"
	m "miniproject/models"
	"miniproject/repository"
	u "miniproject/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAdmin(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	err, res := repository.GetAdminRepository().GetAdmin(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success get admin",
		"Admin":   res,
	})

}

func GetAdmins(c echo.Context) error {

	err, res := repository.GetAdminRepository().GetAdmins()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success get all admins",
		"Admins":  res,
	})
}

func CreateAdmin(c echo.Context) error {

	admin := m.Admin{}
	c.Bind(&admin)

	if err := repository.GetAdminRepository().GetAdminEmail(admin.Email); err == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "Email Account Already Exists",
		})
	}
	valid := md.PostAdminValidation(admin)
	if valid != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": valid.Error(),
		})
	}

	if err := repository.GetAdminRepository().CreateAdmin(&admin); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "succes create new admin",
		"Admin":   admin,
	})
}

func DeleteAdmin(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	if err := repository.GetAdminRepository().DeleteAdmin(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success delete admin",
	})
}

func UpdateAdmin(c echo.Context) error {
	updateData := m.Admin{}
	c.Bind(&updateData)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	if updateData.Password != "" {
		hash, err := u.HashPassword(updateData.Password)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Status":  "400",
				"Message": err.Error(),
			})
		}

		updateData.Password = hash
	}

	if updateData.Email != "" {
		valid := md.EmailValidation(updateData.Email)
		if valid != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Status":  "400",
				"Message": valid.Error(),
			})
		}
	}

	if err := repository.GetAdminRepository().UpdateAdmin(&updateData, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success update admin",
	})
}

func LoginAdmin(c echo.Context) error {
	admin := m.Admin{}
	c.Bind(&admin)

	if err := repository.GetAdminRepository().LoginAdmin(&admin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	token, err := md.CreateToken(int(admin.Id), admin.Name, admin.Role)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	adminResponse := m.AdminResponse{Id: int(admin.Id), Name: admin.Name, Email: admin.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success login",
		"Admin":   adminResponse,
	})
}

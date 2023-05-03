package controller

import (
	"miniproject/config"
	m "miniproject/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetTransactionDetail(c echo.Context) error {

	var transactionDetail m.TransactionDetail

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id",
		})
	}

	if err := config.DB.Where("id = ?", id).First(&transactionDetail).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "transaction detail not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "success get transaction detail",
		"transaction detail": transactionDetail,
	})

}

func GetTransactionDetails(c echo.Context) error {

	var transactionDetails []m.TransactionDetail

	if err := config.DB.Find(&transactionDetails).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":             "success get all transaction details",
		"transaction details": transactionDetails,
	})
}

func CreateTransactionDetail(c echo.Context) error {

	transactionDetail := m.TransactionDetail{}
	c.Bind(&transactionDetail)

	if err := config.DB.Save(&transactionDetail).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "succes create new transaction detail",
		"transaction detail": transactionDetail,
	})
}

func DeleteTransactionDetail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id",
		})
	}

	result := config.DB.Delete(&m.TransactionDetail{}, id)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete transaction detail",
	})
}

func UpdateTransactionDetail(c echo.Context) error {
	updateData := m.TransactionDetail{}
	c.Bind(&updateData)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	result := config.DB.Model(&m.TransactionDetail{}).Where("id = ?", id).Updates(&updateData)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update transaction detail",
	})
}

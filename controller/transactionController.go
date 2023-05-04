package controller

import (
	"miniproject/config"
	m "miniproject/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetTransaction(c echo.Context) error {

	var transaction m.Transaction

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id",
		})
	}

	if err := config.DB.Where("id = ?", id).First(&transaction).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "transaction not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success get transaction",
		"transaction": transaction,
	})

}

func GetTransactions(c echo.Context) error {

	var transactions []m.Transaction

	if err := config.DB.Find(&transactions).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success get all transactions",
		"transactions": transactions,
	})
}

func CreateTransaction(c echo.Context) error {

	transaction := m.Transaction{}
	c.Bind(&transaction)

	if err := config.DB.Save(&transaction).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "succes create new transaction",
		"transaction": transaction,
	})
}

func DeleteTransaction(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id",
		})
	}

	result := config.DB.Delete(&m.Transaction{}, id)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete transaction",
	})
}

func UpdateTransaction(c echo.Context) error {
	updateData := m.Transaction{}
	c.Bind(&updateData)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	result := config.DB.Model(&m.Transaction{}).Where("id = ?", id).Updates(&updateData)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update transaction",
	})
}
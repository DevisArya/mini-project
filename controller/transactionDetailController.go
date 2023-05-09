package controller

import (
	md "miniproject/middleware"
	m "miniproject/models"
	"miniproject/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetTransactionDetail(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}
	err, res := repository.GetTransactionDetailRepository().GetTransactionDetail(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":            "200",
		"Message":           "success get transaction detail",
		"TransactionDetail": res,
	})

}

func GetTransactionDetails(c echo.Context) error {

	err, res := repository.GetTransactionDetailRepository().GetTransactionDetails()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":             "200",
		"Message":            "success get all transaction details",
		"TransactionDetails": res,
	})
}

func CreateTransactionDetail(c echo.Context) error {
	transactionDetail := m.TransactionDetail{}
	c.Bind(&transactionDetail)

	valid := md.PostTransactionDetailValidation(transactionDetail)
	if valid != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": valid.Error(),
		})
	}

	if err := repository.GetTransactionDetailRepository().CreateTransactionDetail(&transactionDetail); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":            "200",
		"Message":           "succes create new transaction detail",
		"TransactionDetail": transactionDetail,
	})
}

func DeleteTransactionDetail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	if err := repository.GetTransactionDetailRepository().DeleteTransactionDetail(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success delete transaction detail",
	})
}

func UpdateTransactionDetail(c echo.Context) error {
	updateData := m.TransactionDetail{}
	c.Bind(&updateData)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	if err := repository.GetTransactionDetailRepository().UpdateTransactionDetail(&updateData, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success update transaction detail",
	})
}

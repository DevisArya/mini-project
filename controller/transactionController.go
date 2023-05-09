package controller

import (
	md "miniproject/middleware"
	m "miniproject/models"
	"miniproject/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetTransaction(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	err, res := repository.GetTransactionRepository().GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":      "200",
		"Message":     "success get transaction",
		"Transaction": res,
	})

}

func GetTransactions(c echo.Context) error {

	err, res := repository.GetTransactionRepository().GetTransactions()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":       "200",
		"Message":      "success get all transactions",
		"Transactions": res,
	})
}

func CreateTransaction(c echo.Context) error {

	transaction := m.Transaction{}
	c.Bind(&transaction)

	valid := md.PostTransactionValidation(transaction)
	if valid != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": valid.Error(),
		})
	}

	if transaction.Rating != 0 || transaction.Status != false {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"Status":  "401",
			"Message": "can't change Status or rating",
		})
	}

	if err := repository.GetTransactionRepository().CreateTransaction(&transaction); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}
	result := m.TransactionResponse{
		Id:         transaction.CustomerId,
		CustomerId: transaction.CustomerId,
		TeamId:     transaction.TeamId,
		PaymentId:  transaction.PaymentId,
		Location:   transaction.Location,
		AreaId:     transaction.AreaId,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":      "200",
		"Message":     "succes create new transaction",
		"Transaction": result,
	})
}

func DeleteTransaction(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	if err := repository.GetTransactionRepository().DeleteTransaction(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success delete transaction",
	})
}

func UpdateTransaction(c echo.Context) error {
	updateData := m.Transaction{}
	c.Bind(&updateData)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	if updateData.Rating != 0 || updateData.Status != false {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"Status":  "401",
			"Message": "can't change Status or rating",
		})
	}

	if err := repository.GetTransactionRepository().UpdateTransaction(&updateData, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success update transaction",
	})
}

func UpdateRating(c echo.Context) error {
	updateRating := m.TransactionUpdateRating{}
	c.Bind(&updateRating)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	valid := md.PostRatingValidation(updateRating)
	if valid != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": valid.Error(),
		})
	}

	if err := repository.GetTransactionRepository().UpdateRating(&updateRating, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success update rating",
	})
}

func UpdateStatus(c echo.Context) error {
	updateStatus := m.TransactionUpdateStatus{}
	c.Bind(&updateStatus)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	valid := md.PostStatusValidation(updateStatus)
	if valid != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": valid.Error(),
		})
	}

	if err := repository.GetTransactionRepository().UpdateStatus(&updateStatus, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success update status",
	})
}

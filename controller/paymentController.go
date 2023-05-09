package controller

import (
	md "miniproject/middleware"
	m "miniproject/models"
	"miniproject/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetPayment(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	err, res := repository.GetPaymentRepository().GetPayment(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success get payment",
		"Payment": res,
	})

}

func GetPayments(c echo.Context) error {

	err, res := repository.GetPaymentRepository().GetPayments()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":   "200",
		"Message":  "success get all payments",
		"Payments": res,
	})
}

func CreatePayment(c echo.Context) error {

	payment := m.Payment{}
	c.Bind(&payment)

	valid := md.PostPaymentValidation(payment)
	if valid != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": valid.Error(),
		})
	}

	if err := repository.GetPaymentRepository().CreatePayment(&payment); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}
	result := m.PaymentResponse{
		Id:            payment.Id,
		Name:          payment.Name,
		PaymentType:   payment.PaymentType,
		PaymentNumber: payment.PaymentNumber,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "succes create new payment",
		"Payment": result,
	})
}

func DeletePayment(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	if err := repository.GetPaymentRepository().DeletePayment(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success delete payment",
	})
}

func UpdatePayment(c echo.Context) error {
	updateData := m.Payment{}
	c.Bind(&updateData)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	if err := repository.GetPaymentRepository().UpdatePayment(&updateData, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success update payment",
	})
}

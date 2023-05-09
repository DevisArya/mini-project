package controller

import (
	"miniproject/config"
	md "miniproject/middleware"
	m "miniproject/models"
	"miniproject/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetChart(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}
	err, res := repository.GetChartRepository().GetChart(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success get chart",
		"chart":   res,
	})

}

func GetCharts(c echo.Context) error {

	err, res := repository.GetChartRepository().GetCharts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success get all charts",
		"Charts":  res,
	})
}

func CreateChart(c echo.Context) error {
	serviceType := m.ServiceType{}
	chart := m.Chart{}
	c.Bind(&chart)

	valid := md.PostChartValidation(chart)
	if valid != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": valid.Error(),
		})
	}

	if err := config.DB.First(&serviceType, chart.ServiceTypeID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	totalPrice := serviceType.Price * uint64(chart.Qty)
	chart.TotalPrice = uint(totalPrice)

	if err := repository.GetChartRepository().CreateChart(&chart); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}
	chartRes := m.ChartResponse{
		Id:            chart.ID,
		CustomerID:    chart.CustomerID,
		ServiceTypeID: chart.ServiceTypeID,
		Qty:           chart.Qty,
		TotalPrice:    chart.TotalPrice,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "succes create new chart",
		"Chart":   chartRes,
	})
}

func DeleteChart(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	if err := repository.GetChartRepository().DeleteChart(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success delete chart",
	})
}

func UpdateChart(c echo.Context) error {
	updateData := m.Chart{}
	c.Bind(&updateData)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	if err := repository.GetChartRepository().UpdateChart(&updateData, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success update chart",
	})
}

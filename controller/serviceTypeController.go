package controller

import (
	md "miniproject/middleware"
	m "miniproject/models"
	"miniproject/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetServiceType(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}
	err, res := repository.GetServiceTypeRepository().GetServiceType(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":      "200",
		"Message":     "success get service type",
		"ServiceType": res,
	})

}

func GetServiceTypes(c echo.Context) error {

	err, res := repository.GetServiceTypeRepository().GetServiceTypes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":       "200",
		"Message":      "success get all service types",
		"ServiceTypes": res,
	})
}

func CreateServiceType(c echo.Context) error {

	serviceType := m.ServiceType{}
	c.Bind(&serviceType)

	valid := md.PostServiceTypeValidation(serviceType)
	if valid != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": valid.Error(),
		})
	}

	if err := repository.GetServiceTypeRepository().CreateServiceType(&serviceType); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}
	result := m.ServiceTypeResponse{
		Id:    serviceType.Id,
		Name:  serviceType.Name,
		Price: serviceType.Price,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":      "200",
		"Message":     "succes create new service type",
		"ServiceType": result,
	})
}

func DeleteServiceType(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	if err := repository.GetServiceTypeRepository().DeleteServiceType(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success delete service type",
	})
}

func UpdateServiceType(c echo.Context) error {
	updateData := m.ServiceType{}
	c.Bind(&updateData)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	if err := repository.GetServiceTypeRepository().UpdateServiceType(&updateData, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success update service type",
	})
}

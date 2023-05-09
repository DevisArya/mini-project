package controller

import (
	md "miniproject/middleware"
	m "miniproject/models"
	"miniproject/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetStore(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	err, res := repository.GetStoreRepository().GetStore(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success get store",
		"Store":   res,
	})

}

func GetStores(c echo.Context) error {

	err, res := repository.GetStoreRepository().GetStores()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success get all stores",
		"Stores":  res,
	})
}

func CreateStore(c echo.Context) error {

	store := m.Store{}
	c.Bind(&store)

	valid := md.PostStoreValidation(store)
	if valid != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": valid.Error(),
		})
	}
	if err := repository.GetStoreRepository().CreateStore(&store); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}
	result := m.StoreResponse{
		Id:      store.Id,
		AreaId:  store.AreaId,
		Address: store.Address,
		Phone:   store.Phone,
		Email:   store.Email,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "succes create new store",
		"Store":   result,
	})
}

func DeleteStore(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	if err := repository.GetStoreRepository().DeleteStore(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success delete store",
	})
}

func UpdateStore(c echo.Context) error {
	updateData := m.Store{}
	c.Bind(&updateData)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	if err := repository.GetStoreRepository().UpdateStore(&updateData, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success update store",
	})
}

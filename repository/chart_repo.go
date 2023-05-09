package repository

import (
	"miniproject/config"
	"miniproject/models"
	"net/http"

	"github.com/labstack/echo"
)

type IChartService interface {
	CreateChart(chart *models.Chart) error
	GetChart(id int) (error, interface{})
	GetCharts() (error, interface{})
	DeleteChart(id int) error
	UpdateChart(dataUpdate *models.Chart, id int) error
}

type ChartRepository struct {
	Func IChartService
}

var chartRepository IChartService

func init() {
	bg := &ChartRepository{}
	bg.Func = bg

	chartRepository = bg
}
func GetChartRepository() IChartService {
	return chartRepository
}
func SetChartRepository(ur IChartService) {
	chartRepository = ur
}

func (u *ChartRepository) CreateChart(chart *models.Chart) error {
	if err := config.DB.Save(&chart).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (u *ChartRepository) GetChart(id int) (err error, res interface{}) {
	var chart models.Chart
	if err := config.DB.Where("id = ?", id).First(&chart).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "chart not found",
		}), nil
	}
	chartRes := models.ChartResponse{
		Id:            chart.ID,
		CustomerID:    chart.CustomerID,
		ServiceTypeID: chart.ServiceTypeID,
		Qty:           chart.Qty,
		TotalPrice:    chart.TotalPrice,
	}
	return nil, chartRes
}
func (u *ChartRepository) GetCharts() (err error, res interface{}) {
	var charts []models.Chart

	if err := config.DB.Find(&charts).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()), nil
	}

	var chartsResponse []models.ChartResponse

	for _, val := range charts {
		chartRes := models.ChartResponse{
			Id:            val.ID,
			CustomerID:    val.CustomerID,
			ServiceTypeID: val.ServiceTypeID,
			Qty:           val.Qty,
			TotalPrice:    val.TotalPrice,
		}
		chartsResponse = append(chartsResponse, chartRes)
	}
	return nil, chartsResponse
}

func (u *ChartRepository) DeleteChart(id int) error {
	result := config.DB.Delete(&models.Chart{}, id)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}
	return nil
}

func (u *ChartRepository) UpdateChart(updateData *models.Chart, id int) error {
	result := config.DB.Model(&models.Chart{}).Where("id = ?", id).Updates(&updateData)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}
	return nil
}

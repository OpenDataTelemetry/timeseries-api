package main

import (
	"github.com/OpenDataTelemetry/timeseries-api/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // Create a new gin router instance
	api := r.Group("/api/timeseries/v0.3/IMT/LNS/")
	{
		api.GET("SmartLight/all", controller.GetAllSmartLight)
		api.GET("SmartLight/deviceId/:deviceId", controller.GetSmartLightByDeviceId)

		api.GET("WaterTankLevel/all", controller.GetAllWaterTankLevel)
		api.GET("WaterTankLevel/deviceId/:deviceId", controller.GetWaterTankLevelByDeviceId)

		api.GET("GaugePressure/all", controller.GetAllGaugePressure)
		api.GET("GaugePressure/deviceId/:deviceId", controller.GetGaugePressureByDeviceId)

		api.GET("Hydrometer/all", controller.GetAllHydrometer)
		api.GET("Hydrometer/deviceId/:deviceId", controller.GetHydrometerByDeviceId)

		api.GET("EnergyMeter/all", controller.GetAllEnergyMeter)
		api.GET("EnergyMeter/deviceId/:deviceId", controller.GetEnergyMeterByDeviceId)

		api.GET("WeatherStation/all", controller.GetAllWeatherStation)
		api.GET("WeatherStation/deviceId/:deviceId", controller.GetWeatherStationByDeviceId)
	}

	r.Run(":8888")
}

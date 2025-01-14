package main

import (
	"github.com/OpenDataTelemetry/timeseries-api/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // Create a new gin router instance

	r.Use(cors.Default())

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

		api.GET("WeatherStation/all", controller.GetAllWeatherStation)
		api.GET("WeatherStation/deviceId/:deviceId", controller.GetWeatherStationByDeviceId)

		api.GET("Sprinkler/all", controller.GetAllSprinkler)
		api.GET("Sprinkler/deviceId/:deviceId", controller.GetSprinklerByDeviceId)

		// api.GET("SoilMoisture3DepthLevels/all", controller.GetAllSoilMoisture3DepthLevels)
		// api.GET("SoilMoisture3DepthLevels/deviceId/:deviceId", controller.GetSoilMoisture3DepthLevelsByDeviceId)

		// api.GET("LnsDownlink/all", controller.GetAllLnsDownlink)
		// api.GET("LnsDownlink/deviceId/:deviceId", controller.GetLnsDownlinkByDeviceId)

		// api.GET("Alert/all", controller.GetAllAlert)
		// api.GET("Alert/deviceId/:deviceId", controller.GetAlertByDeviceId)

	}

	r.Run(":8888")
}

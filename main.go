package main

import (
	"github.com/OpenDataTelemetry/timeseries-api/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // Create a new gin router instance

	r.Use(cors.Default())

	api := r.Group("/api/timeseries/v0.3/IMT/")
	{
		api.GET("LNS/SmartLight/all", controller.GetAllLnsSmartLight)
		api.GET("LNS/SmartLight/deviceId/:deviceId", controller.GetLnsSmartLightByDeviceId)

		api.GET("LNS/WaterTankLevel/all", controller.GetAllLnsWaterTankLevel)
		api.GET("LNS/WaterTankLevel/deviceId/:deviceId", controller.GetLnsWaterTankLevelByDeviceId)

		api.GET("LNS/GaugePressure/all", controller.GetAllLnsGaugePressure)
		api.GET("LNS/GaugePressure/deviceId/:deviceId", controller.GetLnsGaugePressureByDeviceId)

		api.GET("LNS/Hydrometer/all", controller.GetAllLnsHydrometer)
		api.GET("LNS/Hydrometer/deviceId/:deviceId", controller.GetLnsHydrometerByDeviceId)

		api.GET("LNS/EnergyMeter/all", controller.GetAllLnsEnergyMeter)
		api.GET("LNS/EnergyMeter/deviceId/:deviceId", controller.GetLnsEnergyMeterByDeviceId)

		api.GET("LNS/WeatherStation/all", controller.GetAllLnsWeatherStation)
		api.GET("LNS/WeatherStation/deviceId/:deviceId", controller.GetLnsWeatherStationByDeviceId)

		api.GET("LNS/Sprinkler/all", controller.GetAllLnsSprinkler)
		api.GET("LNS/Sprinkler/deviceId/:deviceId", controller.GetLnsSprinklerByDeviceId)

		api.GET("LNS/SoilMoisture3DepthLevels/all", controller.GetAllLnsSoilMoisture3DepthLevels)
		api.GET("LNS/SoilMoisture3DepthLevels/deviceId/:deviceId", controller.GetLnsSoilMoisture3DepthLevelsByDeviceId)

		api.GET("LNS/Downlink/all", controller.GetAllLnsDownlink)
		api.GET("LNS/Downlink/deviceId/:deviceId", controller.GetLnsDownlinkByDeviceId)

		api.GET("LNS/Alert/all", controller.GetAllLnsAlert)
		api.GET("LNS/Alert/deviceId/:deviceId", controller.GetLnsAlertByDeviceId)

		api.GET("NSPI/GenericJson/all", controller.GetAllNspiGenericJson)
		api.GET("NSPI/GenericJson/deviceId/:deviceId", controller.GetNspiGenericJsonByDeviceId)

		api.GET("NSPI/Alert/all", controller.GetAllNspiAlert)
		api.GET("NSPI/Alert/deviceId/:deviceId", controller.GetNspiAlertByDeviceId)

		api.GET("HealthPack/Vital/all", controller.GetAllHealthPackVital)
		api.GET("HealthPack/Vital/deviceId/:deviceId", controller.GetHealthPackVitalByDeviceId)

		api.GET("EVSE/MeterValues/all", controller.GetAllEvseMeterValues)
		api.GET("EVSE/MeterValues/deviceId/:deviceId", controller.GetEvseMeterValuesByDeviceId)

		api.GET("EVSE/StatusNotification/all", controller.GetAllEvseStatusNotification)
		api.GET("EVSE/StatusNotification/deviceId/:deviceId", controller.GetEvseStatusNotificationByDeviceId)

		api.GET("EVSE/StartTransaction/all", controller.GetAllEvseStartTransaction)
		api.GET("EVSE/StartTransaction/deviceId/:deviceId", controller.GetEvseStartTransactionByDeviceId)

		api.GET("EVSE/StopTransaction/all", controller.GetAllEvseStopTransaction)
		api.GET("EVSE/StopTransaction/deviceId/:deviceId", controller.GetEvseStopTransactionByDeviceId)

	}

	r.Run(":8888")
}

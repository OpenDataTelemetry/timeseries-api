package main

import (
	"github.com/OpenDataTelemetry/timeseries-api/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // Create a new gin router instance

	r.Use(cors.Default())

	api := r.Group("/api/timeseries/v0.3/")
	{
		api.GET("IMT/all/Alert/all", controller.GetAllAlert)
		api.GET("IMT/all/Alert/deviceId/:deviceId", controller.GetAlertByDeviceId)

		api.GET("IMT/LNS/SmartLight/all", controller.GetAllLnsSmartLight)
		api.GET("IMT/LNS/SmartLight/deviceId/:deviceId", controller.GetLnsSmartLightByDeviceId)

		api.GET("IMT/LNS/WaterTankLevel/all", controller.GetAllLnsWaterTankLevel)
		api.GET("IMT/LNS/WaterTankLevel/deviceId/:deviceId", controller.GetLnsWaterTankLevelByDeviceId)

		api.GET("IMT/LNS/GaugePressure/all", controller.GetAllLnsGaugePressure)
		api.GET("IMT/LNS/GaugePressure/deviceId/:deviceId", controller.GetLnsGaugePressureByDeviceId)

		api.GET("IMT/LNS/Hydrometer/all", controller.GetAllLnsHydrometer)
		api.GET("IMT/LNS/Hydrometer/deviceId/:deviceId", controller.GetLnsHydrometerByDeviceId)

		api.GET("IMT/LNS/EnergyMeter/all", controller.GetAllLnsEnergyMeter)
		api.GET("IMT/LNS/EnergyMeter/deviceId/:deviceId", controller.GetLnsEnergyMeterByDeviceId)

		api.GET("IMT/LNS/WeatherStation/all", controller.GetAllLnsWeatherStation)
		api.GET("IMT/LNS/WeatherStation/deviceId/:deviceId", controller.GetLnsWeatherStationByDeviceId)

		api.GET("IMT/LNS/Sprinkler/all", controller.GetAllLnsSprinkler)
		api.GET("IMT/LNS/Sprinkler/deviceId/:deviceId", controller.GetLnsSprinklerByDeviceId)

		api.GET("IMT/LNS/SoilMoisture3DepthLevels/all", controller.GetAllLnsSoilMoisture3DepthLevels)
		api.GET("IMT/LNS/SoilMoisture3DepthLevels/deviceId/:deviceId", controller.GetLnsSoilMoisture3DepthLevelsByDeviceId)

		api.GET("IMT/LNS/Command/all", controller.GetAllLnsCommand)
		api.GET("IMT/LNS/Command/deviceId/:deviceId", controller.GetLnsCommandByDeviceId)

		api.GET("IMT/LNS/Alert/all", controller.GetAllLnsAlert)
		api.GET("IMT/LNS/Alert/deviceId/:deviceId", controller.GetLnsAlertByDeviceId)

		api.GET("IMT/NSPI/GenericJson/all", controller.GetAllNspiGenericJson)
		api.GET("IMT/NSPI/GenericJson/deviceId/:deviceId", controller.GetNspiGenericJsonByDeviceId)

		api.GET("IMT/NSPI/Alert/all", controller.GetAllNspiAlert)
		api.GET("IMT/NSPI/Alert/deviceId/:deviceId", controller.GetNspiAlertByDeviceId)

		api.GET("IMT/EVSE/MeterValues/all", controller.GetAllEvseMeterValues)
		api.GET("IMT/EVSE/MeterValues/deviceId/:deviceId", controller.GetEvseMeterValuesByDeviceId)

		api.GET("IMT/EVSE/StatusNotification/all", controller.GetAllEvseStatusNotification)
		api.GET("IMT/EVSE/StatusNotification/deviceId/:deviceId", controller.GetEvseStatusNotificationByDeviceId)

		api.GET("IMT/EVSE/StartTransaction/all", controller.GetAllEvseStartTransaction)
		api.GET("IMT/EVSE/StartTransaction/deviceId/:deviceId", controller.GetEvseStartTransactionByDeviceId)

		api.GET("IMT/EVSE/StopTransaction/all", controller.GetAllEvseStopTransaction)
		api.GET("IMT/EVSE/StopTransaction/deviceId/:deviceId", controller.GetEvseStopTransactionByDeviceId)

		api.GET("IMT/HealthPack/Inertias/all", controller.GetAllHealthPackInertias)
		api.GET("IMT/HealthPack/Inertias/deviceId/:deviceId", controller.GetHealthPackInertiasByDeviceId)

		api.GET("IMT/HealthPack/Tracking/all", controller.GetAllHealthPackTracking)
		api.GET("IMT/HealthPack/Tracking/deviceId/:deviceId", controller.GetHealthPackTrackingByDeviceId)

		api.GET("IMT/HealthPack/Status/all", controller.GetAllHealthPackStatus)
		api.GET("IMT/HealthPack/Status/deviceId/:deviceId", controller.GetHealthPackStatusByDeviceId)

		api.GET("IMT/HealthPack/Ischemia/all", controller.GetAllHealthPackIschemia)
		api.GET("IMT/HealthPack/Ischemia/deviceId/:deviceId", controller.GetHealthPackIschemiaByDeviceId)

		// api.GET("IMT/HealthPack/Alarms/all", controller.GetAllHealthPackAlarms)
		// api.GET("IMT/HealthPack/Alarms/deviceId/:deviceId", controller.GetHealthPackAlarmsByDeviceId)
	}

	r.Run(":8888")
}

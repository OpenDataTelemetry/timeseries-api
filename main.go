package main

import (
	"github.com/gin-gonic/gin"
	"github.com/OpenDataTelemetry/timeseries-api/controllers"
)

func main() {

	r := gin.Default() // Create a new gin router instance
	api := r.Group("/api/timeseries/v0.1/smartcampusmaua/SmartLights")
	{
		api.GET("", controllers.GetSmartLights)
		api.GET("deviceName/:nodename", controllers.GetSmartLightbyNodeName)
		api.GET("deviceId/:devEUI", controllers.GetSmartLightbyDevEUI)
	}

	r.Run(":8888")
}

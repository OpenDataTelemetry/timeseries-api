package main

import (
	"github.com/gin-gonic/gin"
	"github.com/OpenDataTelemetry/timeseries-api/controller"
)

func main() {

	r := gin.Default() // Create a new gin router instance
	api := r.Group("/api/timeseries/v0.1/smartcampusmaua/SmartLights")
	{
		api.GET("", controller.GetSmartLights)
		api.GET("deviceName/:nodename", controller.GetSmartLightbyNodeName)
		api.GET("deviceId/:devEUI", controller.GetSmartLightbyDevEUI)
	}

	r.Run(":8888")
}

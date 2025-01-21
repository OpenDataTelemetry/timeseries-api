package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/OpenDataTelemetry/timeseries-api/database"
	"github.com/gin-gonic/gin"
)

// Uplink
// EnergyMeter
func GetAllLnsEnergyMeter(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "EnergyMeter"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		ORDER BY time DESC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"forwardEnergy":  value["forwardEnergy"],
				"reverseEnergy":  value["reverseEnergy"],
				"boardVoltage":   value["boardVoltage"],
				"data":           value["data"],
				"fCnt":           value["fCnt"],
				"fPort":          value["fPort"],
				"rxAlt_0":        value["rxAlt_0"],
				"rxLat_0":        value["rxLat_0"],
				"rxLon_0":        value["rxLon_0"],
				"rxRssi_0":       value["rxRssi_0"],
				"rxSnr_0":        value["rxSnr_0"],
				"txBandWidth":    value["txBandWidth"],
				"txFrequency":    value["txFrequency"],
				"txSpreadFactor": value["txSpreadFactor"],
			},
			"name": "EnergyMeter",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"host":       value["host"],
				"origin":     value["origin"],
				"rxMac_0":    value["rxMac_0"],
				// "txCodeRate":   value["txCodeRate"],
				// "txModulation": value["txModulation"],
				"type": value["type"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetLnsEnergyMeterByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	deviceId := c.Param("deviceId") // Parameter to query
	var objs = []gin.H{}            // Slice to store the query response in a list
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
		SELECT *
		FROM "EnergyMeter"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceId" IN ('` + deviceId + `')
		ORDER BY time DESC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() {
		value := iterator.Value()
		obj := gin.H{
			"fields": gin.H{
				"forwardEnergy":  value["forwardEnergy"],
				"reverseEnergy":  value["reverseEnergy"],
				"boardVoltage":   value["boardVoltage"],
				"data":           value["data"],
				"fCnt":           value["fCnt"],
				"fPort":          value["fPort"],
				"rxAlt_0":        value["rxAlt_0"],
				"rxLat_0":        value["rxLat_0"],
				"rxLon_0":        value["rxLon_0"],
				"rxRssi_0":       value["rxRssi_0"],
				"rxSnr_0":        value["rxSnr_0"],
				"txBandWidth":    value["txBandWidth"],
				"txFrequency":    value["txFrequency"],
				"txSpreadFactor": value["txSpreadFactor"],
			},
			"name": "EnergyMeter",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"host":       value["host"],
				"origin":     value["origin"],
				"rxMac_0":    value["rxMac_0"],
				// "txCodeRate":   value["txCodeRate"],
				// "txModulation": value["txModulation"],
				"type": value["type"],
			},
			"timestamp": value["time"],
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs)
}

// GaugePressure
func GetAllLnsGaugePressure(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "GaugePressure"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		ORDER BY time DESC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"inletPressure":  value["inletPressure"],
				"outletPressure": value["outletPressure"],
				"boardVoltage":   value["boardVoltage"],
				"data":           value["data"],
				"fCnt":           value["fCnt"],
				"fPort":          value["fPort"],
				"rxAlt_0":        value["rxAlt_0"],
				"rxLat_0":        value["rxLat_0"],
				"rxLon_0":        value["rxLon_0"],
				"rxRssi_0":       value["rxRssi_0"],
				"rxSnr_0":        value["rxSnr_0"],
				"txBandWidth":    value["txBandWidth"],
				"txFrequency":    value["txFrequency"],
				"txSpreadFactor": value["txSpreadFactor"],
			},
			"name": "GaugePressure",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"host":       value["host"],
				"origin":     value["origin"],
				"rxMac_0":    value["rxMac_0"],
				// "txCodeRate":   value["txCodeRate"],
				// "txModulation": value["txModulation"],
				"type": value["type"],
			},
			"timestamp": value["time"],
		} // Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetLnsGaugePressureByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	deviceId := c.Param("deviceId") // Parameter to query
	var objs = []gin.H{}            // Slice to store the query response in a list
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
		SELECT *
		FROM "GaugePressure"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceId" IN ('` + deviceId + `')
		ORDER BY time DESC;
	`
	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() {
		value := iterator.Value()
		obj := gin.H{
			"fields": gin.H{
				"inletPressure":  value["inletPressure"],
				"outletPressure": value["outletPressure"],
				"boardVoltage":   value["boardVoltage"],
				"data":           value["data"],
				"fCnt":           value["fCnt"],
				"fPort":          value["fPort"],
				"rxAlt_0":        value["rxAlt_0"],
				"rxLat_0":        value["rxLat_0"],
				"rxLon_0":        value["rxLon_0"],
				"rxRssi_0":       value["rxRssi_0"],
				"rxSnr_0":        value["rxSnr_0"],
				"txBandWidth":    value["txBandWidth"],
				"txFrequency":    value["txFrequency"],
				"txSpreadFactor": value["txSpreadFactor"],
			},
			"name": "GaugePressure",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"host":       value["host"],
				"origin":     value["origin"],
				"rxMac_0":    value["rxMac_0"],
				// "txCodeRate":   value["txCodeRate"],
				// "txModulation": value["txModulation"],
				"type": value["type"],
			},
			"timestamp": value["time"],
		} // Convert the row to a gin.H map (JSON)
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs)
}

// Hydrometer
func GetAllLnsHydrometer(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "Hydrometer"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		ORDER BY time DESC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"counter":        value["counter"],
				"boardVoltage":   value["boardVoltage"],
				"data":           value["data"],
				"fCnt":           value["fCnt"],
				"fPort":          value["fPort"],
				"rxAlt_0":        value["rxAlt_0"],
				"rxLat_0":        value["rxLat_0"],
				"rxLon_0":        value["rxLon_0"],
				"rxRssi_0":       value["rxRssi_0"],
				"rxSnr_0":        value["rxSnr_0"],
				"txBandWidth":    value["txBandWidth"],
				"txFrequency":    value["txFrequency"],
				"txSpreadFactor": value["txSpreadFactor"],
			},
			"name": "Hydrometer",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"host":       value["host"],
				"origin":     value["origin"],
				"rxMac_0":    value["rxMac_0"],
				// "txCodeRate":   value["txCodeRate"],
				// "txModulation": value["txModulation"],
				"type": value["type"],
			},
			"timestamp": value["time"],
		} // Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetLnsHydrometerByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	deviceId := c.Param("deviceId") // Parameter to query
	var objs = []gin.H{}            // Slice to store the query response in a list
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
		SELECT *
		FROM "Hydrometer"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceId" IN ('` + deviceId + `')
		ORDER BY time DESC;
	`
	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() {
		value := iterator.Value()
		obj := gin.H{
			"fields": gin.H{
				"counter":        value["counter"],
				"boardVoltage":   value["boardVoltage"],
				"data":           value["data"],
				"fCnt":           value["fCnt"],
				"fPort":          value["fPort"],
				"rxAlt_0":        value["rxAlt_0"],
				"rxLat_0":        value["rxLat_0"],
				"rxLon_0":        value["rxLon_0"],
				"rxRssi_0":       value["rxRssi_0"],
				"rxSnr_0":        value["rxSnr_0"],
				"txBandWidth":    value["txBandWidth"],
				"txFrequency":    value["txFrequency"],
				"txSpreadFactor": value["txSpreadFactor"],
			},
			"name": "Hydrometer",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"host":       value["host"],
				"origin":     value["origin"],
				"rxMac_0":    value["rxMac_0"],
				// "txCodeRate":   value["txCodeRate"],
				// "txModulation": value["txModulation"],
				"type": value["type"],
			},
			"timestamp": value["time"],
		} // Convert the row to a gin.H map (JSON)
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs)
}

// SmartLight
func GetAllLnsSmartLight(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "SmartLight"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		ORDER BY time DESC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"temperature":    value["temperature"],
				"humidity":       value["humidity"],
				"movement":       value["movement"],
				"luminosity":     value["luminosity"],
				"batteryVoltage": value["batteryVoltage"],
				"boardVoltage":   value["boardVoltage"],
				"data":           value["data"],
				"fCnt":           value["fCnt"],
				"fPort":          value["fPort"],
				"rxAlt_0":        value["rxAlt_0"],
				"rxLat_0":        value["rxLat_0"],
				"rxLon_0":        value["rxLon_0"],
				"rxRssi_0":       value["rxRssi_0"],
				"rxSnr_0":        value["rxSnr_0"],
				"txBandWidth":    value["txBandWidth"],
				"txFrequency":    value["txFrequency"],
				"txSpreadFactor": value["txSpreadFactor"],
			},
			"name": "SmartLight",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"host":       value["host"],
				"origin":     value["origin"],
				"rxMac_0":    value["rxMac_0"],
				// "txCodeRate":   value["txCodeRate"],
				// "txModulation": value["txModulation"],
				"type": value["type"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetLnsSmartLightByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	deviceId := c.Param("deviceId") // Parameter to query
	var objs = []gin.H{}            // Slice to store the query response in a list
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
		SELECT *
		FROM "SmartLight"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceId" IN ('` + deviceId + `')
		ORDER BY time DESC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() {
		value := iterator.Value()
		obj := gin.H{
			"fields": gin.H{
				"temperature":    value["temperature"],
				"humidity":       value["humidity"],
				"movement":       value["movement"],
				"luminosity":     value["luminosity"],
				"batteryVoltage": value["batteryVoltage"],
				"boardVoltage":   value["boardVoltage"],
				"data":           value["data"],
				"fCnt":           value["fCnt"],
				"fPort":          value["fPort"],
				"rxAlt_0":        value["rxAlt_0"],
				"rxLat_0":        value["rxLat_0"],
				"rxLon_0":        value["rxLon_0"],
				"rxRssi_0":       value["rxRssi_0"],
				"rxSnr_0":        value["rxSnr_0"],
				"txBandWidth":    value["txBandWidth"],
				"txFrequency":    value["txFrequency"],
				"txSpreadFactor": value["txSpreadFactor"],
			},
			"name": "SmartLight",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"host":       value["host"],
				"origin":     value["origin"],
				"rxMac_0":    value["rxMac_0"],
				// "txCodeRate":   value["txCodeRate"],
				// "txModulation": value["txModulation"],
				"type": value["type"],
			},
			"timestamp": value["time"],
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs)
}

// SoilMoisture3DepthLevels
func GetAllLnsSoilMoisture3DepthLevels(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "SoilMoisture3DepthLevels"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		ORDER BY time DESC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"soilMoistureDepthLevel1": value["soilMoistureDepthLevel1"],
				"soilMoistureDepthLevel2": value["soilMoistureDepthLevel2"],
				"soilMoistureDepthLevel3": value["soilMoistureDepthLevel3"],
				"boardVoltage":            value["boardVoltage"],
				"data":                    value["data"],
				"fCnt":                    value["fCnt"],
				"fPort":                   value["fPort"],
				"rxAlt_0":                 value["rxAlt_0"],
				"rxLat_0":                 value["rxLat_0"],
				"rxLon_0":                 value["rxLon_0"],
				"rxRssi_0":                value["rxRssi_0"],
				"rxSnr_0":                 value["rxSnr_0"],
				"txBandWidth":             value["txBandWidth"],
				"txFrequency":             value["txFrequency"],
				"txSpreadFactor":          value["txSpreadFactor"],
			},
			"name": "SoilMoisture3DepthLevels",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"host":       value["host"],
				"origin":     value["origin"],
				"rxMac_0":    value["rxMac_0"],
				// "txCodeRate":   value["txCodeRate"],
				// "txModulation": value["txModulation"],
				"type": value["type"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetLnsSoilMoisture3DepthLevelsByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	deviceId := c.Param("deviceId") // Parameter to query
	var objs = []gin.H{}            // Slice to store the query response in a list
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
		SELECT *
		FROM "SoilMoisture3DepthLevels"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceId" IN ('` + deviceId + `')
		ORDER BY time DESC;
	`
	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() {
		value := iterator.Value()
		obj := gin.H{
			"fields": gin.H{
				"soilMoistureDepthLevel1": value["soilMoistureDepthLevel1"],
				"soilMoistureDepthLevel2": value["soilMoistureDepthLevel2"],
				"soilMoistureDepthLevel3": value["soilMoistureDepthLevel3"],
				"boardVoltage":            value["boardVoltage"],
				"data":                    value["data"],
				"fCnt":                    value["fCnt"],
				"fPort":                   value["fPort"],
				"rxAlt_0":                 value["rxAlt_0"],
				"rxLat_0":                 value["rxLat_0"],
				"rxLon_0":                 value["rxLon_0"],
				"rxRssi_0":                value["rxRssi_0"],
				"rxSnr_0":                 value["rxSnr_0"],
				"txBandWidth":             value["txBandWidth"],
				"txFrequency":             value["txFrequency"],
				"txSpreadFactor":          value["txSpreadFactor"],
			},
			"name": "SoilMoisture3DepthLevels",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"host":       value["host"],
				"origin":     value["origin"],
				"rxMac_0":    value["rxMac_0"],
				// "txCodeRate":   value["txCodeRate"],
				// "txModulation": value["txModulation"],
				"type": value["type"],
			},
			"timestamp": value["time"],
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs)
}

// Sprinkler
func GetAllLnsSprinkler(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "Sprinkler"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		ORDER BY time DESC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"solenoid1":      value["solenoid1"],
				"solenoid2":      value["solenoid2"],
				"solenoid3":      value["solenoid3"],
				"counter":        value["counter"],
				"boardVoltage":   value["boardVoltage"],
				"data":           value["data"],
				"fCnt":           value["fCnt"],
				"fPort":          value["fPort"],
				"rxAlt_0":        value["rxAlt_0"],
				"rxLat_0":        value["rxLat_0"],
				"rxLon_0":        value["rxLon_0"],
				"rxRssi_0":       value["rxRssi_0"],
				"rxSnr_0":        value["rxSnr_0"],
				"txBandWidth":    value["txBandWidth"],
				"txFrequency":    value["txFrequency"],
				"txSpreadFactor": value["txSpreadFactor"],
			},
			"name": "Sprinkler",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"host":       value["host"],
				"origin":     value["origin"],
				"rxMac_0":    value["rxMac_0"],
				// "txCodeRate":   value["txCodeRate"],
				// "txModulation": value["txModulation"],
				"type": value["type"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetLnsSprinklerByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	deviceId := c.Param("deviceId") // Parameter to query
	var objs = []gin.H{}            // Slice to store the query response in a list
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
		SELECT *
		FROM "Sprinkler"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceId" IN ('` + deviceId + `')
		ORDER BY time DESC;
	`
	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() {
		value := iterator.Value()
		obj := gin.H{
			"fields": gin.H{
				"solenoid1":      value["solenoid1"],
				"solenoid2":      value["solenoid2"],
				"solenoid3":      value["solenoid3"],
				"counter":        value["counter"],
				"boardVoltage":   value["boardVoltage"],
				"data":           value["data"],
				"fCnt":           value["fCnt"],
				"fPort":          value["fPort"],
				"rxAlt_0":        value["rxAlt_0"],
				"rxLat_0":        value["rxLat_0"],
				"rxLon_0":        value["rxLon_0"],
				"rxRssi_0":       value["rxRssi_0"],
				"rxSnr_0":        value["rxSnr_0"],
				"txBandWidth":    value["txBandWidth"],
				"txFrequency":    value["txFrequency"],
				"txSpreadFactor": value["txSpreadFactor"],
			},
			"name": "Sprinkler",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"host":       value["host"],
				"origin":     value["origin"],
				"rxMac_0":    value["rxMac_0"],
				// "txCodeRate":   value["txCodeRate"],
				// "txModulation": value["txModulation"],
				"type": value["type"],
			},
			"timestamp": value["time"],
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs)
}

// WaterTankLevel
func GetAllLnsWaterTankLevel(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "WaterTankLevel"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		ORDER BY time DESC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"distance":       value["distance"],
				"boardVoltage":   value["boardVoltage"],
				"data":           value["data"],
				"fCnt":           value["fCnt"],
				"fPort":          value["fPort"],
				"rxAlt_0":        value["rxAlt_0"],
				"rxLat_0":        value["rxLat_0"],
				"rxLon_0":        value["rxLon_0"],
				"rxRssi_0":       value["rxRssi_0"],
				"rxSnr_0":        value["rxSnr_0"],
				"txBandWidth":    value["txBandWidth"],
				"txFrequency":    value["txFrequency"],
				"txSpreadFactor": value["txSpreadFactor"],
			},
			"name": "WaterTankLevel",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"host":       value["host"],
				"origin":     value["origin"],
				"rxMac_0":    value["rxMac_0"],
				// "txCodeRate":   value["txCodeRate"],
				// "txModulation": value["txModulation"],
				"type": value["type"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetLnsWaterTankLevelByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	deviceId := c.Param("deviceId") // Parameter to query
	var objs = []gin.H{}            // Slice to store the query response in a list
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
		SELECT *
		FROM "WaterTankLevel"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceId" IN ('` + deviceId + `')
		ORDER BY time DESC;
	`
	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() {
		value := iterator.Value()
		obj := gin.H{
			"fields": gin.H{
				"distance":       value["distance"],
				"boardVoltage":   value["boardVoltage"],
				"data":           value["data"],
				"fCnt":           value["fCnt"],
				"fPort":          value["fPort"],
				"rxAlt_0":        value["rxAlt_0"],
				"rxLat_0":        value["rxLat_0"],
				"rxLon_0":        value["rxLon_0"],
				"rxRssi_0":       value["rxRssi_0"],
				"rxSnr_0":        value["rxSnr_0"],
				"txBandWidth":    value["txBandWidth"],
				"txFrequency":    value["txFrequency"],
				"txSpreadFactor": value["txSpreadFactor"],
			},
			"name": "WaterTankLevel",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"host":       value["host"],
				"origin":     value["origin"],
				"rxMac_0":    value["rxMac_0"],
				// "txCodeRate":   value["txCodeRate"],
				// "txModulation": value["txModulation"],
				"type": value["type"],
			},
			"timestamp": value["time"],
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs)
}

// WeatherStation
func GetAllLnsWeatherStation(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "WeatherStation"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		ORDER BY time DESC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"powerSource":            value["powerSource"],
				"envSensorFailStatus":    value["envSensorFailStatus"],
				"internalBatteryVoltage": value["internalBatteryVoltage"],
				"firmwareVersion":        value["firmwareVersion"],
				"c1State":                value["c1State"],
				"c1Count":                value["c1Count"],
				"c2State":                value["c2State"],
				"c2Count":                value["c2Count"],
				"internalTemperature":    value["internalTemperature"],
				"internalHumidity":       value["internalHumidity"],
				"emwRainLevel":           value["emwRainLevel"],
				"emwAvgWindSpeed":        value["emwAvgWindSpeed"],
				"emwGustWindSpeed":       value["emwGustWindSpeed"],
				"emwWindDirection":       value["emwWindDirection"],
				"emwTemperature":         value["emwTemperature"],
				"emwHumidity":            value["emwHumidity"],
				"emwLuminosity":          value["emwLuminosity"],
				"emwUv":                  value["emwUv"],
				"emwSolarRadiation":      value["emwSolarRadiation"],
				"emwAtmPres":             value["emwAtmPres"],
				"data":                   value["data"],
				"fCnt":                   value["fCnt"],
				"fPort":                  value["fPort"],
				"rxAlt_0":                value["rxAlt_0"],
				"rxLat_0":                value["rxLat_0"],
				"rxLon_0":                value["rxLon_0"],
				"rxRssi_0":               value["rxRssi_0"],
				"rxSnr_0":                value["rxSnr_0"],
				"txBandWidth":            value["txBandWidth"],
				"txFrequency":            value["txFrequency"],
				"txSpreadFactor":         value["txSpreadFactor"],
			},
			"name": "WeatherStation",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"host":       value["host"],
				"origin":     value["origin"],
				"rxMac_0":    value["rxMac_0"],
				// "txCodeRate":   value["txCodeRate"],
				// "txModulation": value["txModulation"],
				"type": value["type"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetLnsWeatherStationByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	deviceId := c.Param("deviceId") // Parameter to query
	var objs = []gin.H{}            // Slice to store the query response in a list
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
		SELECT *
		FROM "WeatherStation"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceId" IN ('` + deviceId + `')
		ORDER BY time DESC;
	`
	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() {
		value := iterator.Value()
		obj := gin.H{
			"fields": gin.H{
				"powerSource":            value["powerSource"],
				"envSensorFailStatus":    value["envSensorFailStatus"],
				"internalBatteryVoltage": value["internalBatteryVoltage"],
				"firmwareVersion":        value["firmwareVersion"],
				"c1State":                value["c1State"],
				"c1Count":                value["c1Count"],
				"c2State":                value["c2State"],
				"c2Count":                value["c2Count"],
				"internalTemperature":    value["internalTemperature"],
				"internalHumidity":       value["internalHumidity"],
				"emwRainLevel":           value["emwRainLevel"],
				"emwAvgWindSpeed":        value["emwAvgWindSpeed"],
				"emwGustWindSpeed":       value["emwGustWindSpeed"],
				"emwWindDirection":       value["emwWindDirection"],
				"emwTemperature":         value["emwTemperature"],
				"emwHumidity":            value["emwHumidity"],
				"emwLuminosity":          value["emwLuminosity"],
				"emwUv":                  value["emwUv"],
				"emwSolarRadiation":      value["emwSolarRadiation"],
				"emwAtmPres":             value["emwAtmPres"],
				"data":                   value["data"],
				"fCnt":                   value["fCnt"],
				"fPort":                  value["fPort"],
				"rxAlt_0":                value["rxAlt_0"],
				"rxLat_0":                value["rxLat_0"],
				"rxLon_0":                value["rxLon_0"],
				"rxRssi_0":               value["rxRssi_0"],
				"rxSnr_0":                value["rxSnr_0"],
				"txBandWidth":            value["txBandWidth"],
				"txFrequency":            value["txFrequency"],
				"txSpreadFactor":         value["txSpreadFactor"],
			},
			"name": "WeatherStation",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"host":       value["host"],
				"origin":     value["origin"],
				"rxMac_0":    value["rxMac_0"],
				// "txCodeRate":   value["txCodeRate"],
				// "txModulation": value["txModulation"],
				"type": value["type"],
			},
			"timestamp": value["time"],
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs)
}

// Command
func GetAllLnsCommand(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "Command"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('LNS')
		ORDER BY time DESC;
	`
	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"confirmed": value["confirmed"],
				"data":      value["data"],
				"fPort":     value["fPort"],
			},
			"name": "Command",
			"tags": gin.H{
				"application": value["application"],
				"deviceId":    value["deviceId"],
				"deviceType":  value["deviceType"],
				"direction":   value["direction"],
				"host":        value["host"],
				"origin":      value["origin"],
				"reference":   value["reference"],
				"type":        value["type"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetLnsCommandByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	deviceId := c.Param("deviceId") // Parameter to query
	var objs = []gin.H{}            // Slice to store the query response in a list
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
		SELECT *
		FROM "Command"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('LNS')
		AND
		"deviceId" IN ('` + deviceId + `')
		ORDER BY time DESC;
	`
	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() {
		value := iterator.Value()
		obj := gin.H{
			"fields": gin.H{
				"confirmed": value["confirmed"],
				"data":      value["data"],
				"fPort":     value["fPort"],
			},
			"name": "Command",
			"tags": gin.H{
				"application": value["application"],
				"deviceId":    value["deviceId"],
				"deviceType":  value["deviceType"],
				"direction":   value["direction"],
				"host":        value["host"],
				"origin":      value["origin"],
				"reference":   value["reference"],
				"type":        value["type"],
			},
			"timestamp": value["time"],
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs)
}

// Alert
func GetAllLnsAlert(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "Alert"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('LNS')
		ORDER BY time DESC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"data": value["data"],
			},
			"name": "Alert",
			"tags": gin.H{
				"deviceId": value["deviceId"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetLnsAlertByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	deviceId := c.Param("deviceId") // Parameter to query
	var objs = []gin.H{}            // Slice to store the query response in a list
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
		SELECT *
		FROM "Alert"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('LNS')
		AND
		"deviceId" IN ('` + deviceId + `')
		ORDER BY time DESC;
	`
	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() {
		value := iterator.Value()
		obj := gin.H{
			"fields": gin.H{
				"data": value["data"],
			},
			"name": "Alert",
			"tags": gin.H{
				"deviceId": value["deviceId"],
			},
			"timestamp": value["time"],
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs)
}

func GetAllAlert(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "Alert"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		ORDER BY time DESC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"data":         value["data"],
				"trigger":      value["trigger"],
				"triggerAt":    value["triggerAt"],
				"triggerType":  value["triggerType"],
				"lastPlayed":   value["lastPlayed"],
				"actionSensor": value["actionSensor"],
				"currentValue": value["currentValue"],
			},
			"name": "Alert",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"etc":        value["origin"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetAlertByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	deviceId := c.Param("deviceId") // Parameter to query
	var objs = []gin.H{}            // Slice to store the query response in a list
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
		SELECT *
		FROM "Alert"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceId" IN ('` + deviceId + `')
		ORDER BY time DESC;
	`
	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() {
		value := iterator.Value()
		obj := gin.H{
			"fields": gin.H{
				"data":         value["data"],
				"trigger":      value["trigger"],
				"triggerAt":    value["triggerAt"],
				"triggerType":  value["triggerType"],
				"lastPlayed":   value["lastPlayed"],
				"actionSensor": value["actionSensor"],
				"currentValue": value["currentValue"],
			},
			"name": "Alert",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"etc":        value["origin"],
			},
			"timestamp": value["time"],
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs)
}

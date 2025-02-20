package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/OpenDataTelemetry/timeseries-api/database"
	"github.com/gin-gonic/gin"
)

func GetAllEvseMeterValues(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 43200 {
		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
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
		FROM "MeterValues"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('EVSE')
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
				"forwardEnergy": value["forwardEnergy"],
			},
			"name": "EvseMeterValues",
			"tags": gin.H{
				"chargePointId": value["chargePointId"],
				"connectorId":   value["connectorId"],
				"deviceType":    value["deviceType"],
				"deviceId":      value["deviceId"],
				"direction":     value["direction"],
				"origin":        value["origin"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetEvseMeterValuesByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 43200 {
		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
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
		FROM "MeterValues"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('EVSE')
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
				"forwardEnergy": value["forwardEnergy"],
			},
			"name": "EvseMeterValues",
			"tags": gin.H{
				"chargePointId": value["chargePointId"],
				"connectorId":   value["connectorId"],
				"deviceType":    value["deviceType"],
				"deviceId":      value["deviceId"],
				"direction":     value["direction"],
				"origin":        value["origin"],
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

func GetAllEvseStatusNotification(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 43200 {
		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
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
		FROM "StatusNotification"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('EVSE')
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
				"info":            value["info"],
				"errorCode":       value["errorCode"],
				"VendorErrorCode": value["VendorErrorCode"],
			},
			"name": "EvseStatusNotification",
			"tags": gin.H{
				"chargePointId": value["chargePointId"],
				"connectorId":   value["connectorId"],
				"deviceType":    value["deviceType"],
				"deviceId":      value["deviceId"],
				"direction":     value["direction"],
				"origin":        value["origin"],
				"vendorId":      value["vendorId"],
				"status":        value["status"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetEvseStatusNotificationByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 43200 {
		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
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
		FROM "StatusNotification"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('EVSE')
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
				"info":            value["info"],
				"errorCode":       value["errorCode"],
				"VendorErrorCode": value["VendorErrorCode"],
			},
			"name": "EvseStatusNotification",
			"tags": gin.H{
				"chargePointId": value["chargePointId"],
				"connectorId":   value["connectorId"],
				"deviceType":    value["deviceType"],
				"deviceId":      value["deviceId"],
				"direction":     value["direction"],
				"origin":        value["origin"],
				"vendorId":      value["vendorId"],
				"status":        value["status"],
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

func GetAllEvseStartTransaction(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 43200 {
		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
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
		FROM "StartTransaction"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('EVSE')
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
				"startMeter":    value["startMeter"],
				"transactionId": value["transactionId"],
				"startTime":     value["startTime"],
				"idTag":         value["idTag"],
			},
			"name": "EvseStartTransaction",
			"tags": gin.H{
				"chargePointId": value["chargePointId"],
				"connectorId":   value["connectorId"],
				"deviceType":    value["deviceType"],
				"deviceId":      value["deviceId"],
				"direction":     value["direction"],
				"origin":        value["origin"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetEvseStartTransactionByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 43200 {
		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
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
		FROM "StartTransaction"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('EVSE')
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
				"startMeter":    value["startMeter"],
				"transactionId": value["transactionId"],
				"startTime":     value["startTime"],
				"idTag":         value["idTag"],
			},
			"name": "EvseStartTransaction",
			"tags": gin.H{
				"chargePointId": value["chargePointId"],
				"connectorId":   value["connectorId"],
				"deviceType":    value["deviceType"],
				"deviceId":      value["deviceId"],
				"direction":     value["direction"],
				"origin":        value["origin"],
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

func GetAllEvseStopTransaction(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 43200 {
		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
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
		FROM "StopTransaction"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('EVSE')
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
				"transactionId": value["transactionId"],
				"MeterStop":     value["MeterStop"],
				"stopTime":      value["stopTime"],
			},
			"name": "EvseStopTransaction",
			"tags": gin.H{
				"chargePointId": value["chargePointId"],
				"connectorId":   value["connectorId"],
				"deviceType":    value["deviceType"],
				"deviceId":      value["deviceId"],
				"direction":     value["direction"],
				"origin":        value["origin"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetEvseStopTransactionByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 43200 {
		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
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
		FROM "StopTransaction"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('EVSE')
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
				"transactionId": value["transactionId"],
				"MeterStop":     value["MeterStop"],
				"stopTime":      value["stopTime"],
			},
			"name": "EvseStopTransaction",
			"tags": gin.H{
				"chargePointId": value["chargePointId"],
				"connectorId":   value["connectorId"],
				"deviceType":    value["deviceType"],
				"deviceId":      value["deviceId"],
				"direction":     value["direction"],
				"origin":        value["origin"],
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
func GetAllEvseAlert(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 43200 {
		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
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
		"deviceType" IN ('EVSE')
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

func GetEvseAlertByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 43200 {
		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
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
		"deviceType" IN ('EVSE')
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

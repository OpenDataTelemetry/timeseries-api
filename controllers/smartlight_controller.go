package controllers

import (
	"context"
	"net/http"
	"github.com/OpenDataTelemetry/timeseries-api/initializers"
	"github.com/gin-gonic/gin"
)

func GetSmartLights(c *gin.Context) {
	interval := c.Query("interval")

	
	var objs = []gin.H{}
	influxDB, err := initializers.ConnectToDB()
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
SELECT *
FROM "SmartLights"
WHERE "time" >= now() - interval '`+ interval +` minutes'
ORDER BY time DESC;
`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H(value)       // Convert the row to a gin.H map (JSON)
		objs = append(objs, obj)  // Append the row to the objs slice
	}

	c.IndentedJSON(http.StatusOK, objs)

}

func GetSmartLightbyNodeName(c *gin.Context) {
	nodename := c.Param("nodename") // Parameter to query

	var objs = []gin.H{} // Slice to store the query response in a list
	influxDB, err := initializers.ConnectToDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
	SELECT *
	FROM "SmartLights"
	WHERE "nodeName" = '` + nodename + `'
	ORDER BY time DESC;
	`
	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H(value)       // Convert the row to a gin.H map (JSON)
		objs = append(objs, obj)  // Append the row to the objs slice
	}

	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceName not found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, objs) // Return the objs slice as a JSON response
}

func GetSmartLightbyDevEUI(c *gin.Context) {
	devEUI := c.Param("devEUI") // Parameter to query
	var objs = []gin.H{}        // Slice to store the query response in a list
	influxDB, err := initializers.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
	SELECT *
	FROM "SmartLights"
	WHERE "devEUI" = '` + devEUI + `'
	ORDER BY time DESC;
	`
	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H(value)       // Convert the row to a gin.H map (JSON)
		objs = append(objs, obj)  // Append the row to the objs slice
	}
	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs) // Return the objs slice as a JSON response
}

package main

import (
	"context"
	// "fmt"
	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)


func getSmartLights(c *gin.Context) {

	var objs = []gin.H{}
	url := os.Getenv("INFLUXDB_URL")
	token := os.Getenv("INFLUXDB_TOKEN")
	database := os.Getenv("INFLUXDB_DATABASE")

	influxdb3Client, err := influxdb3.New(influxdb3.ClientConfig{
		Host:     url,
		Token:    token,
		Database: database,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxdb3Client.Close()
	query := `
SELECT *
FROM "SmartLights"
WHERE time >= now() - interval '2 hour'
ORDER BY time DESC;
`

	iterator, err := influxdb3Client.Query(context.Background(), query)// Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H(value) // Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}

	c.IndentedJSON(http.StatusOK, objs)

}





func getSmartLightbyNodeName(c *gin.Context) {
	nodename := c.Param("nodename") // Parameter to query
	var objs = []gin.H{} 					// Slice to store the query response in a list
	url := os.Getenv("INFLUXDB_URL") 
	token := os.Getenv("INFLUXDB_TOKEN")
	database := os.Getenv("INFLUXDB_DATABASE")

	influxdb3Client, err := influxdb3.New(influxdb3.ClientConfig{ // Create a new client connection
		Host:     url,
		Token:    token,
		Database: database,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxdb3Client.Close() // Close the client connection after the function ends
	query := `
	SELECT *
	FROM "SmartLights"
	WHERE "nodeName" = '` + nodename + `'
	ORDER BY time DESC;
	`
	iterator, err := influxdb3Client.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H(value) 		 // Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}

	c.IndentedJSON(http.StatusOK, objs) // Return the objs slice as a JSON response
}



func getSmartLightbyDevEUI(c *gin.Context) {
	devEUI := c.Param("devEUI") // Parameter to query
	var objs = []gin.H{} 					// Slice to store the query response in a list
	url := os.Getenv("INFLUXDB_URL") 
	token := os.Getenv("INFLUXDB_TOKEN")
	database := os.Getenv("INFLUXDB_DATABASE")

	influxdb3Client, err := influxdb3.New(influxdb3.ClientConfig{ // Create a new client connection
		Host:     url,
		Token:    token,
		Database: database,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxdb3Client.Close() // Close the client connection after the function ends
	query := `
	SELECT *
	FROM "SmartLights"
	WHERE "devEUI" = '` + devEUI + `'
	ORDER BY time DESC;
	`
	iterator, err := influxdb3Client.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H(value) 		 // Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}

	c.IndentedJSON(http.StatusOK, objs) // Return the objs slice as a JSON response
}


func main() {

	r := gin.Default() // Create a new gin router instance
	api := r.Group("/api/v0.1/smartcampusmaua/SmartLights")
	{
		api.GET("", getSmartLights)
		api.GET("nodeName/:nodename", getSmartLightbyNodeName)
		api.GET("devEUI/:devEUI", getSmartLightbyDevEUI)
	}

	r.Run(":8888")
}

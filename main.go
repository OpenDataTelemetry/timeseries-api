package main

import (
	"context"
	"log"
	// "strconv"
	// "fmt"
	"net/http"
	"os"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/gin-gonic/gin"
)

func ConnectToDB() (*influxdb3.Client, error) {
	url := os.Getenv("INFLUXDB_URL")
	token := os.Getenv("INFLUXDB_TOKEN")
	database := os.Getenv("INFLUXDB_DATABASE")

	influxdb3Client, err := influxdb3.New(influxdb3.ClientConfig{
		Host:     url,
		Token:    token,
		Database: database,
	})

	if err != nil {
		log.Fatal("Failed to connect to database")
		return &influxdb3.Client{}, err
	} 

	return influxdb3Client, nil
}
func getSmartLights(c *gin.Context) {

	var objs = []gin.H{}
	
	influxDB, err := ConnectToDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()
	query := `
SELECT *
FROM "SmartLights"
WHERE time >= now() - interval '2 hour'
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

func getSmartLightbyNodeName(c *gin.Context) {
	nodename := c.Param("nodename") // Parameter to query
	// temperature := c.Param("temperature")

	var objs = []gin.H{} // Slice to store the query response in a list
	influxDB, err := ConnectToDB()
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

	c.IndentedJSON(http.StatusOK, objs) // Return the objs slice as a JSON response
}

func getSmartLightbyDevEUI(c *gin.Context) {
	devEUI := c.Param("devEUI") // Parameter to query
	var objs = []gin.H{}        // Slice to store the query response in a list
	influxDB, err := ConnectToDB()

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

	c.IndentedJSON(http.StatusOK, objs) // Return the objs slice as a JSON response
}

func main() {

	r := gin.Default() // Create a new gin router instance
	api := r.Group("/api/v0.1/smartcampusmaua/SmartLights")
	{
		api.GET("", getSmartLights)
		api.GET("nodeName/:nodename/tempera", getSmartLightbyNodeName)
		api.GET("devEUI/:devEUI", getSmartLightbyDevEUI)
	}

	r.Run(":8888")
}

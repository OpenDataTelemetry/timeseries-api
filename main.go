package main

import (
	"context"
	// "fmt"
	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	// "reflect"
)

var influxdb3Client *influxdb3.Client

func getAllSmartLights(c *gin.Context) {

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
ORDER BY time DESC;
`

	iterator, err := influxdb3Client.Query(context.Background(), query)

	if err != nil {
		panic(err)
	}

	for iterator.Next() {
		value := iterator.Value()
		obj := gin.H(value)
		objs = append(objs, obj)
	}

	c.IndentedJSON(http.StatusOK, objs)

}

func getSmartLightbyID(c *gin.Context) {
	nodename := c.Param("nodename")
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
WHERE "nodename" = '` + nodename + `'
ORDER BY time DESC;
`
	iterator, err := influxdb3Client.Query(context.Background(), query)

	if err != nil {
		panic(err)
	}

	for iterator.Next() {
		value := iterator.Value()
		obj := gin.H(value)
		objs = append(objs, obj)
	}

	c.IndentedJSON(http.StatusOK, objs)
}

func main() {

	r := gin.Default()
	api := r.Group("/api/v0.1/smartcampusmaua")
	{
		api.GET("/SmartLights", getAllSmartLights)
		// api.GET("/SmartLights/:nodename", getSmartLightbyID)

	}

	r.Run(":8888")
}

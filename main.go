package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/gin-gonic/gin"
)

func main() {
	// Use env variables to initialize client
	url := os.Getenv("INFLUXDB_URL")
	token := os.Getenv("INFLUXDB_TOKEN")
	database := os.Getenv("INFLUXDB_DATABASE")

	// Create a new client using an InfluxDB server base URL and an authentication token
	influxdb3Client, err := influxdb3.New(influxdb3.ClientConfig{
		Host:     url,
		Token:    token,
		Database: database,
	})

	if err != nil {
		panic(err)
	}

	// Close influxdb3Client at the end and escalate error if present
	defer func(influxdb3Client *influxdb3.Client) {
		err := influxdb3Client.Close()
		if err != nil {
			panic(err)
		}
	}(influxdb3Client)

	// API
	r := gin.Default()
	r.GET("/api/v0.1/smartcampusmaua/SmartLights", func(c *gin.Context) {

		// Prepare FlightSQL query
		query := `
		SELECT *
		FROM "SmartLights"
		WHERE
		time >= now() - interval '1 hour'
		AND
		("data_ad" IS NOT NULL OR "data_boardVoltage" IS NOT NULL)
		ORDER BY time DESC;
	`
		
		iterator, err := influxdb3Client.Query(context.Background(), query)
		fmt.Printf("\niterator: %v\n", iterator)

		if err != nil {
			panic(err)
		}
		
		
		objs := []gin.H{}
		
		for iterator.Next() {
			value := iterator.Value()
			obj := gin.H(value)
			objs = append(objs, obj)
		}
		c.JSON(http.StatusOK, objs)

	})
	r.Run(":8888")
}


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
	r.GET("/api/timeseries", func(c *gin.Context) {

		// Prepare FlightSQL query
		query := `
					SELECT *
						FROM "APPDET"
						WHERE
						time >= now() - interval '7 days'
						AND
						("fCnt" IS NOT NULL)
						AND
						"nodeName" IN ('DET-03') AND "rxInfo_mac_0" IN ('b827ebfffe5ddbb6')
						ORDER BY time DESC
	`
		// queryOptions := influxdb3.QueryOptions{
		// 	QueryType: influxdb3.InfluxQL,
		// }

		// iterator, err := influxdb3Client.QueryWithOptions(context.Background(), &queryOptions, query)
		iterator, err := influxdb3Client.Query(context.Background(), query)
		fmt.Printf("\niterator: %v\n", iterator)

		if err != nil {
			panic(err)
		}

		for iterator.Next() {
			value := iterator.Value()
			fmt.Printf("\nvalue: %v\n", value)

			fCnt := value["fCnt"]
			fmt.Printf("Frame Count: %v\n", fCnt)

			// API
			c.JSON(http.StatusOK, gin.H{
				"fCnt": fCnt,
			})
		}

	})
	r.Run(":8888")
}

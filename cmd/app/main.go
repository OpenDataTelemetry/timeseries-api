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
			data_ad := value["data_ad"]
			data_boardVoltage := value["data_boardVoltage"]
			data_counter := value["data_counter"]
			data_counter_0d_0 := value["data_counter_0d_0"]
			data_counter_0d_1 := value["data_counter_0d_1"]
			data_counter_0d_2 := value["data_counter_0d_2"]
			data_counter_4_20 := value["data_counter_4_20"]
			data_distance := value["data_distance"]
			data_humidity := value["data_humidity"]
			data_lat := value["data_lat"]
			data_lon := value["data_lon"]
			data_temperature := value["data_temperature"]
			rxInfo_altitude_0 := value["rxInfo.altitude_0"]
			rxInfo_altitude_1 := value["rxInfo.altitude_1"]
			rxInfo_altitude_2 := value["rxInfo.altitude_2"]
			rxInfo_latitude_0 := value["rxInfo.latitude_0"]
			rxInfo_latitude_1 := value["rxInfo.latitude_1"]
			rxInfo_latitude_2 := value["rxInfo.latitude_2"]
			rxInfo_loRaSNR_0 := value["rxInfo.loRaSNR_0"]
			rxInfo_loRaSNR_1 := value["rxInfo.loRaSNR_1"]
			rxInfo_loRaSNR_2 := value["rxInfo.loRaSNR_2"]
			rxInfo_longitude_0 := value["rxInfo.longitude_0"]
			rxInfo_longitude_1 := value["rxInfo.longitude_1"]
			rxInfo_longitude_2 := value["rxInfo.longitude_2"]
			rxInfo_rssi_0 := value["rxInfo.rssi_0"]
			rxInfo_rssi_1 := value["rxInfo.rssi_1"]
			rxInfo_rssi_2 := value["rxInfo.rssi_2"]
			txInfo_dataRate_spreadFactor := value["txInfo_dataRate_spreadFactor"]
			txInfo_frequency := value["txInfo_frequency"]

			fmt.Printf("Frame Count: %v\n", fCnt)

			// API
			c.JSON(http.StatusOK, gin.H{
				"fCnt": fCnt,
				"data_ad": data_ad,
				"data_boardVoltage": data_boardVoltage,
				"data_counter" : data_counter,
				"data_counter_0d_0" : data_counter_0d_0,
				"data_counter_0d_1" : data_counter_0d_1,
				"data_counter_0d_2" : data_counter_0d_2,
				"data_counter_4_20": data_counter_4_20,
				"data_distance": data_distance,
				"data_humidity" : data_humidity,
				"data_lat" : data_lat,
				"data_lon" : data_lon,
				"data_temperature" : data_temperature,
				"rxInfo_altitude_0" : rxInfo_altitude_0,
				"rxInfo_altitude_1" : rxInfo_altitude_1,
				"rxInfo_altitude_2" : rxInfo_altitude_2,
				"rxInfo_latitude_0" : rxInfo_latitude_0,
				"rxInfo_latitude_1" : rxInfo_latitude_1,
				"rxInfo_latitude_2" : rxInfo_latitude_2,
				"rxInfo_loRaSNR_0" : rxInfo_loRaSNR_0,
				"rxInfo_loRaSNR_1" : rxInfo_loRaSNR_1,
				"rxInfo_loRaSNR_2" : rxInfo_loRaSNR_2,
				"rxInfo_longitude_0" : rxInfo_longitude_0,
				"rxInfo_longitude_1" : rxInfo_longitude_1,
				"rxInfo_longitude_2" : rxInfo_longitude_2,
				"rxInfo_rssi_0" : rxInfo_rssi_0,
				"rxInfo_rssi_1" : rxInfo_rssi_1,
				"rxInfo_rssi_2" : rxInfo_rssi_2,
				"txInfo_dataRate_spreadFactor" : txInfo_dataRate_spreadFactor,
				"txInfo_frequency" : txInfo_frequency,
			})
		}

	})
	r.Run(":8888")
}

package database

import (
	"log"
	"os"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
)

func ConnectToDB() (*influxdb3.Client, error) {
	url := os.Getenv("INFLUXDB_URL")
	token := os.Getenv("INFLUXDB_TOKEN")
	bucket := os.Getenv("INFLUXDB_BUCKET")

	influxdb3Client, err := influxdb3.New(influxdb3.ClientConfig{
		Host:     url,
		Token:    token,
		Database: bucket,
	})

	if err != nil {
		log.Fatal("Failed to connect to bucket")
		return &influxdb3.Client{}, err
	}

	return influxdb3Client, nil
}

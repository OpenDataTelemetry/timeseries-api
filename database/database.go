package database

import (
	"log"
	"os"
	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
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
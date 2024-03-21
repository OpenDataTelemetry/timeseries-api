package controller

import (
	"context"
	"net/http"

	"github.com/OpenDataTelemetry/timeseries-api/database"
	"github.com/gin-gonic/gin"
)

type SmartLight struct {
	Fields struct {
			DataCounter0d0               float64 `json:"data_counter_0d_0"`
			DataCounter0d1               float64 `json:"data_counter_0d_1"`
			DataEnergy0                  float64 `json:"data_energy_0"`
			DataEnergy1                  float64 `json:"data_energy_1"`
			FCnt                         int     `json:"fCnt"`
			RxInfoAltitude0              float64 `json:"rxInfo_altitude_0"`
			RxInfoLatitude0              float64 `json:"rxInfo_latitude_0"`
			RxInfoLoRaSNR0               int     `json:"rxInfo_loRaSNR_0"`
			RxInfoLongitude0             float64 `json:"rxInfo_longitude_0"`
			RxInfoRssi0                  int     `json:"rxInfo_rssi_0"`
			TxInfoDataRateSpreadFactor   int     `json:"txInfo_dataRate_spreadFactor"`
			TxInfoFrequency              int     `json:"txInfo_frequency"`
	} `json:"fields"`
	Name string `json:"name"`
	Tags struct {
			ApplicationID            string  `json:"applicationID"`
			DevEUI                   string  `json:"devEUI"`
			FPort                    string  `json:"fPort"`
			NodeName                 string  `json:"nodeName"`
			RxInfoMac0               string  `json:"rxInfo_mac_0"`
			RxInfoName0              string  `json:"rxInfo_name_0"`
			TxInfoAdr                string  `json:"txInfo_adr"`
			TxInfoCodeRate           string  `json:"txInfo_codeRate"`
			TxInfoDataRateBandwidth  string  `json:"txInfo_dataRate_bandwidth"`
			TxInfoDataRateModulation string  `json:"txInfo_dataRate_modulation"`
	} `json:"tags"`
	Timestamp int64 `json:"timestamp"`
}

func GetSmartLights(c *gin.Context) {
	interval := c.Query("interval")

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "SmartLights"
		WHERE "time" >= now() - interval '` + interval + ` minutes'
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
					"data_counter_0d_0": value["data_counter_0d_0"],
					"data_counter_0d_1": value["data_counter_0d_1"],
					"data_energy_0": value["data_energy_0"],
					"data_energy_1": value["data_energy_1"],
					"fCnt": value["fCnt"],
					"rxInfo_altitude_0": value["rxInfo_altitude_0"],
					"rxInfo_latitude_0": value["rxInfo_latitude_0"],
					"rxInfo_loRaSNR_0": value["rxInfo_loRaSNR_0"],
					"rxInfo_longitude_0": value["rxInfo_longitude_0"],
					"rxInfo_rssi_0": value["rxInfo_rssi_0"],
					"txInfo_dataRate_spreadFactor": value["txInfo_dataRate_spreadFactor"],
					"txInfo_frequency": value["txInfo_frequency"],
			},
			"name": "SmartLights",
			"tags": gin.H{
					// Substitua as chaves e valores abaixo pelos valores apropriados
					"applicationID": value["applicationID"],
					"devEUI": value["devEUI"],
					"fPort": value["fPort"],
					"nodeName": value["nodeName"],
					"rxInfo_mac_0": value["rxInfo_mac_0"],
					"rxInfo_name_0": value["rxInfo_name_0"],
					"txInfo_adr": "true",
					"txInfo_codeRate": "4/5",
					"txInfo_dataRate_bandwidth": "125",
					"txInfo_dataRate_modulation": "LORA",
			},
			"timestamp": value["time"], // Substitua por value["timestamp"] se o timestamp estiver disponível no valor
	}       
		// "data_counter_0d_0" : value["data_counter_0d_0"]}       // Convert the row to a gin.H map (JSON)
		objs = append(objs, obj)  // Append the row to the objs slice
	}
	c.IndentedJSON(http.StatusOK, objs)


}

func GetSmartLightbyNodeName(c *gin.Context) {
	nodename := c.Param("nodename") // Parameter to query

	var objs = []gin.H{} // Slice to store the query response in a list
	influxDB, err := database.ConnectToDB()
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
		obj := gin.H{
			"fields": gin.H{
					"data_counter_0d_0": value["data_counter_0d_0"],
					"data_counter_0d_1": value["data_counter_0d_1"],
					"data_energy_0": value["data_energy_0"],
					"data_energy_1": value["data_energy_1"],
					"fCnt": value["fCnt"],
					"rxInfo_altitude_0": value["rxInfo_altitude_0"],
					"rxInfo_latitude_0": value["rxInfo_latitude_0"],
					"rxInfo_loRaSNR_0": value["rxInfo_loRaSNR_0"],
					"rxInfo_longitude_0": value["rxInfo_longitude_0"],
					"rxInfo_rssi_0": value["rxInfo_rssi_0"],
					"txInfo_dataRate_spreadFactor": value["txInfo_dataRate_spreadFactor"],
					"txInfo_frequency": value["txInfo_frequency"],
			},
			"name": "SmartLights",
			"tags": gin.H{
					// Substitua as chaves e valores abaixo pelos valores apropriados
					"applicationID": value["applicationID"],
					"devEUI": value["devEUI"],
					"fPort": value["fPort"],
					"nodeName": value["nodeName"],
					"rxInfo_mac_0": value["rxInfo_mac_0"],
					"rxInfo_name_0": value["rxInfo_name_0"],
					"txInfo_adr": "true",
					"txInfo_codeRate": "4/5",
					"txInfo_dataRate_bandwidth": "125",
					"txInfo_dataRate_modulation": "LORA",
			},
			"timestamp": value["time"], // Substitua por value["timestamp"] se o timestamp estiver disponível no valor
	}       // Convert the row to a gin.H map (JSON)
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
	influxDB, err := database.ConnectToDB()

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
		obj := gin.H{
			"fields": gin.H{
					"data_counter_0d_0": value["data_counter_0d_0"],
					"data_counter_0d_1": value["data_counter_0d_1"],
					"data_energy_0": value["data_energy_0"],
					"data_energy_1": value["data_energy_1"],
					"fCnt": value["fCnt"],
					"rxInfo_altitude_0": value["rxInfo_altitude_0"],
					"rxInfo_latitude_0": value["rxInfo_latitude_0"],
					"rxInfo_loRaSNR_0": value["rxInfo_loRaSNR_0"],
					"rxInfo_longitude_0": value["rxInfo_longitude_0"],
					"rxInfo_rssi_0": value["rxInfo_rssi_0"],
					"txInfo_dataRate_spreadFactor": value["txInfo_dataRate_spreadFactor"],
					"txInfo_frequency": value["txInfo_frequency"],
			},
			"name": "SmartLights",
			"tags": gin.H{
					// Substitua as chaves e valores abaixo pelos valores apropriados
					"applicationID": value["applicationID"],
					"devEUI": value["devEUI"],
					"fPort": value["fPort"],
					"nodeName": value["nodeName"],
					"rxInfo_mac_0": value["rxInfo_mac_0"],
					"rxInfo_name_0": value["rxInfo_name_0"],
					"txInfo_adr": "true",
					"txInfo_codeRate": "4/5",
					"txInfo_dataRate_bandwidth": "125",
					"txInfo_dataRate_modulation": "LORA",
			},
			"timestamp": value["time"], // Substitua por value["timestamp"] se o timestamp estiver disponível no valor
	}       // Convert the row to a gin.H map (JSON)
		objs = append(objs, obj)  // Append the row to the objs slice
	}
	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs) // Return the objs slice as a JSON response
}

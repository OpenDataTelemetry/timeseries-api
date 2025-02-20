package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/OpenDataTelemetry/timeseries-api/database"
	"github.com/gin-gonic/gin"
)

func GetAllHealthPackInertias(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 43200 {
		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "Inertias"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('HealthPack')
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
				"fAccX":        value["fAccX"],
				"fAccY":        value["fAccY"],
				"fAccZ":        value["fAccZ"],
				"accX":         value["accX"],
				"accY":         value["accY"],
				"accZ":         value["accZ"],
				"gyrX":         value["gyrX"],
				"gyrY":         value["gyrY"],
				"gyrZ":         value["gyrZ"],
				"contimpactoX": value["contimpactoX"],
				"contimpactoY": value["contimpactoY"],
				"contimpactoZ": value["contimpactoZ"],
				"pitch":        value["pitch"],
				"roll":         value["roll"],
				"yaw":          value["yaw"],
			},
			"name": "Inertias",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"macAddress": value["macAddress"],
				"deviceIp":   value["deviceIp"],
				"direction":  value["direction"],
				"origin":     value["origin"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetHealthPackInertiasByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 43200 {
		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
		return
	}

	deviceId := c.Param("deviceId") // Parameter to query
	var objs = []gin.H{}            // Slice to store the query response in a list
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
		SELECT *
		FROM "Inertias"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('HealthPack')
		AND
		"deviceId" IN ('` + deviceId + `')
		ORDER BY time DESC;
	`
	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() {
		value := iterator.Value()
		obj := gin.H{
			"fields": gin.H{
				"fAccX":        value["fAccX"],
				"fAccY":        value["fAccY"],
				"fAccZ":        value["fAccZ"],
				"accX":         value["accX"],
				"accY":         value["accY"],
				"accZ":         value["accZ"],
				"gyrX":         value["gyrX"],
				"gyrY":         value["gyrY"],
				"gyrZ":         value["gyrZ"],
				"contimpactoX": value["contimpactoX"],
				"contimpactoY": value["contimpactoY"],
				"contimpactoZ": value["contimpactoZ"],
				"pitch":        value["pitch"],
				"roll":         value["roll"],
				"yaw":          value["yaw"],
			},
			"name": "Inertias",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"macAddress": value["macAddress"],
				"deviceIp":   value["deviceIp"],
				"direction":  value["direction"],
				"origin":     value["origin"],
			},
			"timestamp": value["time"],
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs)
}

func GetAllHealthPackTracking(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 43200 {
		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "Tracking"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('HealthPack')
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
				"latitude":               value["latitude"],
				"longitude":              value["longitude"],
				"tempbateriasecundaria":  value["tempbateriasecundaria"],
				"tempbateriaprincipal":   value["tempbateriaprincipal"],
				"temperaturacondensador": value["temperaturacondensador"],
				"temperaturacuba1":       value["temperaturacuba1"],
				"temperaturacuba2":       value["temperaturacuba2"],
				"temperaturaexternaLL":   value["temperaturaexternaLL"],
				"temperaturaexternaLS":   value["temperaturaexternaLS"],
				"temperaturaexterna":     value["temperaturaexterna"],
				"temperaturadissipador":  value["temperaturadissipador"],
				"correntebateria":        value["correntebateria"],
				"correntecompressor":     value["correntecompressor"],
				"correntepeltier":        value["correntepeltier"],
				"correntecooler":         value["correntecooler"],
				"correnteexaustor":       value["correnteexaustor"],
				"temperaturacompressor":  value["temperaturacompressor"],
				"setpoint_pid1":          value["setpoint_pid1"],
				"valor_pid1_atual":       value["valor_pid1_atual"],
				"esforco_pid1":           value["esforco_pid1"],
				"setpoint_pid2":          value["setpoint_pid2"],
				"valor_pid2_atual":       value["valor_pid2_atual"],
				"esforco_pid2":           value["esforco_pid2"],
			},
			"name": "Tracking",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"macAddress": value["macAddress"],
				"deviceIp":   value["deviceIp"],
				"direction":  value["direction"],
				"origin":     value["origin"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetHealthPackTrackingByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 43200 {
		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
		return
	}

	deviceId := c.Param("deviceId") // Parameter to query
	var objs = []gin.H{}            // Slice to store the query response in a list
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
		SELECT *
		FROM "Tracking"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('HealthPack')
		AND
		"deviceId" IN ('` + deviceId + `')
		ORDER BY time DESC;
	`
	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() {
		value := iterator.Value()
		obj := gin.H{
			"fields": gin.H{
				"latitude":               value["latitude"],
				"longitude":              value["longitude"],
				"tempbateriasecundaria":  value["tempbateriasecundaria"],
				"tempbateriaprincipal":   value["tempbateriaprincipal"],
				"temperaturacondensador": value["temperaturacondensador"],
				"temperaturacuba1":       value["temperaturacuba1"],
				"temperaturacuba2":       value["temperaturacuba2"],
				"temperaturaexternaLL":   value["temperaturaexternaLL"],
				"temperaturaexternaLS":   value["temperaturaexternaLS"],
				"temperaturaexterna":     value["temperaturaexterna"],
				"temperaturadissipador":  value["temperaturadissipador"],
				"correntebateria":        value["correntebateria"],
				"correntecompressor":     value["correntecompressor"],
				"correntepeltier":        value["correntepeltier"],
				"correntecooler":         value["correntecooler"],
				"correnteexaustor":       value["correnteexaustor"],
				"temperaturacompressor":  value["temperaturacompressor"],
				"setpoint_pid1":          value["setpoint_pid1"],
				"valor_pid1_atual":       value["valor_pid1_atual"],
				"esforco_pid1":           value["esforco_pid1"],
				"setpoint_pid2":          value["setpoint_pid2"],
				"valor_pid2_atual":       value["valor_pid2_atual"],
				"esforco_pid2":           value["esforco_pid2"],
			},
			"name": "Tracking",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"macAddress": value["macAddress"],
				"deviceIp":   value["deviceIp"],
				"direction":  value["direction"],
				"origin":     value["origin"],
			},
			"timestamp": value["time"],
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs)
}

func GetAllHealthPackStatus(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 43200 {
		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "Status"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('HealthPack')
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
				"vbateriaprincipal":           value["vbateriaprincipal"],
				"vbateriasecundaria":          value["vbateriasecundaria"],
				"ventradafonteexterna":        value["ventradafonteexterna"],
				"numerocaixa":                 value["numerocaixa"],
				"estadomaquina":               value["estadomaquina"],
				"timestamp":                   value["timestamp"],
				"idFalha":                     value["idFalha"],
				"porcentagemsinalcomunicacao": value["porcentagemsinalcomunicacao"],
				"fatorRH":                     value["fatorRH"],
				"btdown":                      value["btdown"],
				"btselect":                    value["btselect"],
				"btup":                        value["btup"],
				"tecla_enter":                 value["tecla_enter"],
				"statustampaprincipal":        value["statustampaprincipal"],
				"statusserialprincipal":       value["statusserialprincipal"],
				"statusserialsecundaria":      value["statusserialsecundaria"],
				"statustampacasamaq":          value["statustampacasamaq"],
				"controle_peltier":            value["controle_peltier"],
				"porcentagem_bat":             value["porcentagem_bat"],
			},
			"name": "Status",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"macAddress": value["macAddress"],
				"deviceIp":   value["deviceIp"],
				"direction":  value["direction"],
				"origin":     value["origin"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetHealthPackStatusByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 43200 {
		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
		return
	}

	deviceId := c.Param("deviceId") // Parameter to query
	var objs = []gin.H{}            // Slice to store the query response in a list
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
		SELECT *
		FROM "Status"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('HealthPack')
		AND
		"deviceId" IN ('` + deviceId + `')
		ORDER BY time DESC;
	`
	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() {
		value := iterator.Value()
		obj := gin.H{
			"fields": gin.H{
				"vbateriaprincipal":           value["vbateriaprincipal"],
				"vbateriasecundaria":          value["vbateriasecundaria"],
				"ventradafonteexterna":        value["ventradafonteexterna"],
				"numerocaixa":                 value["numerocaixa"],
				"estadomaquina":               value["estadomaquina"],
				"timestamp":                   value["timestamp"],
				"idFalha":                     value["idFalha"],
				"porcentagemsinalcomunicacao": value["porcentagemsinalcomunicacao"],
				"fatorRH":                     value["fatorRH"],
				"btdown":                      value["btdown"],
				"btselect":                    value["btselect"],
				"btup":                        value["btup"],
				"tecla_enter":                 value["tecla_enter"],
				"statustampaprincipal":        value["statustampaprincipal"],
				"statusserialprincipal":       value["statusserialprincipal"],
				"statusserialsecundaria":      value["statusserialsecundaria"],
				"statustampacasamaq":          value["statustampacasamaq"],
				"controle_peltier":            value["controle_peltier"],
				"porcentagem_bat":             value["porcentagem_bat"],
			},
			"name": "Status",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"macAddress": value["macAddress"],
				"deviceIp":   value["deviceIp"],
				"direction":  value["direction"],
				"origin":     value["origin"],
			},
			"timestamp": value["time"],
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs)
}

func GetAllHealthPackIschemia(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 43200 {
		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "Ischemia"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('HealthPack')
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
				"idModal":                 value["idModal"],
				"IdOperador":              value["IdOperador"],
				"niveldepermissao":        value["niveldepermissao"],
				"nome":                    value["nome"],
				"numtransplante":          value["numtransplante"],
				"numeroempresa":           value["numeroempresa"],
				"orgao":                   value["orgao"],
				"tempo_total_isquemia":    value["tempo_total_isquemia"],
				"tempo_restante_isquemia": value["tempo_restante_isquemia"],
				"hora_isquemia":           value["hora_isquemia"],
				"timeinfo_sp2":            value["timeinfo_sp2"],
			},
			"name": "Ischemia",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"macAddress": value["macAddress"],
				"deviceIp":   value["deviceIp"],
				"direction":  value["direction"],
				"origin":     value["origin"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetHealthPackIschemiaByDeviceId(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 43200 {
		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
		return
	}

	deviceId := c.Param("deviceId") // Parameter to query
	var objs = []gin.H{}            // Slice to store the query response in a list
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close() // Close the client connection after the function ends
	query := `
		SELECT *
		FROM "Ischemia"
		WHERE 
		time >= now() - interval '` + intervalStr + ` minutes'
		AND
		"deviceType" IN ('HealthPack')
		AND
		"deviceId" IN ('` + deviceId + `')
		ORDER BY time DESC;
	`
	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() {
		value := iterator.Value()
		obj := gin.H{
			"fields": gin.H{
				"idModal":                 value["idModal"],
				"IdOperador":              value["IdOperador"],
				"niveldepermissao":        value["niveldepermissao"],
				"nome":                    value["nome"],
				"numtransplante":          value["numtransplante"],
				"numeroempresa":           value["numeroempresa"],
				"orgao":                   value["orgao"],
				"tempo_total_isquemia":    value["tempo_total_isquemia"],
				"tempo_restante_isquemia": value["tempo_restante_isquemia"],
				"hora_isquemia":           value["hora_isquemia"],
				"timeinfo_sp2":            value["timeinfo_sp2"],
			},
			"name": "Ischemia",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"macAddress": value["macAddress"],
				"deviceIp":   value["deviceIp"],
				"direction":  value["direction"],
				"origin":     value["origin"],
			},
			"timestamp": value["time"],
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs)
}

// func GetAllHealthPackAlarms(c *gin.Context) {
// 	intervalStr := c.Query("interval")
// 	interval, err := strconv.Atoi(intervalStr)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
// 		return
// 	}

// 	if interval > 43200 {
// 		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
// 		return
// 	}

// 	var objs = []gin.H{}
// 	influxDB, err := database.ConnectToDB()

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer influxDB.Close()

// 	query := `
// 		SELECT *
// 		FROM "Alarms"
// 		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
// 		AND
// 		"deviceType" IN ('HealthPack')
// 		ORDER BY time DESC;
// 	`

// 	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

// 	if err != nil {
// 		panic(err)
// 	}

// 	for iterator.Next() { // Iterate over query response
// 		value := iterator.Value() // Value of the current row
// 		obj := gin.H{
// 			"fields": gin.H{
// 				"alarms": value["alarms"],
// 			},
// 			"name": "Alarms",
// 			"tags": gin.H{
// 				"deviceId":   value["deviceId"],
// 				"deviceType": value["deviceType"],
// 				"macAddress": value["macAddress"],
// 				"deviceIp":   value["deviceIp"],
// 				"direction":  value["direction"],
// 				"origin":     value["origin"],
// 			},
// 			"timestamp": value["time"],
// 		}
// 		// Convert the row to a gin.H map (JSON)
// 		objs = append(objs, obj) // Append the row to the objs slice
// 	}
// 	fmt.Println(len(objs))
// 	c.IndentedJSON(http.StatusOK, objs)
// }

// func GetHealthPackAlarmsByDeviceId(c *gin.Context) {
// 	intervalStr := c.Query("interval")
// 	interval, err := strconv.Atoi(intervalStr)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
// 		return
// 	}

// 	if interval > 43200 {
// 		c.JSON(400, gin.H{"error": "Interval must be less than 43200"})
// 		return
// 	}

// 	deviceId := c.Param("deviceId") // Parameter to query
// 	var objs = []gin.H{}            // Slice to store the query response in a list
// 	influxDB, err := database.ConnectToDB()

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer influxDB.Close() // Close the client connection after the function ends
// 	query := `
// 		SELECT *
// 		FROM "Alarms"
// 		WHERE
// 		time >= now() - interval '` + intervalStr + ` minutes'
// 		AND
// 		"deviceType" IN ('HealthPack')
// 		AND
// 		"deviceId" IN ('` + deviceId + `')
// 		ORDER BY time DESC;
// 	`
// 	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

// 	if err != nil {
// 		panic(err)
// 	}

// 	for iterator.Next() {
// 		value := iterator.Value()
// 		obj := gin.H{
// 			"fields": gin.H{
// 				"alarms": value["alarms"],
// 			},
// 			"name": "Alarms",
// 			"tags": gin.H{
// 				"deviceId":   value["deviceId"],
// 				"deviceType": value["deviceType"],
// 			},
// 			"timestamp": value["time"],
// 		}
// 		objs = append(objs, obj)
// 	}
// 	if len(objs) == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "deviceId not found!"})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, objs)
// }

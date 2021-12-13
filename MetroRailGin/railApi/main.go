package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Vishnukvsvk/MetroRailGin/dbutils"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// StationResource holds information about locations
type StationResource struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

//CURL -X GET "http://localhost:8000/v1/stations/1"
func GetStation(c *gin.Context) {
	var station StationResource
	id := c.Param("station_id")
	err := DB.QueryRow("select ID,NAME,CAST(OPENING_TIME AS CHAR),CAST(CLOSING_TIME AS CHAR) FROM station where ID=?", id).Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"result": station,
		})
	}
}

//curl -X POST http://localhost:8000/v1/stations -H 'cache-control: no-cache' -H 'content-type: application/json' -d '{"name":"Brooklyn", "opening_time":"8:12:00", "closing_time":"18:23:00"}'
func CreateStation(c *gin.Context) {
	var station StationResource
	// Parse the body into our resrource
	if err := c.BindJSON(&station); err == nil {
		// Format Time to Go time format
		statement, _ := DB.Prepare("insert into station (NAME, OPENING_TIME, CLOSING_TIME) values (?, ?, ?)")
		result, _ := statement.Exec(station.Name, station.OpeningTime, station.ClosingTime)
		if err == nil {
			newID, _ := result.LastInsertId()
			station.ID = int(newID)
			c.JSON(http.StatusOK, gin.H{
				"result": station,
			})
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

//curl -X DELETE "http://localhost:8000/v1/stations/1"
func DeleteStation(c *gin.Context) {
	id := c.Param("station_id")
	statement, _ := DB.Prepare("delete from station where id=?")
	_, err := statement.Exec(id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.String(http.StatusOK, "")
	}
}

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "./railapi.db") // No :=, beacuse already declared in global level
	if err != nil {
		log.Println("Driver creation failed!")
	}
	dbutils.Initialize(DB)
	r := gin.Default()
	//Add routes to REST verbs
	r.GET("/v1/stations/:station_id", GetStation)
	r.POST("/v1/stations", CreateStation) //No "/" after stations
	r.DELETE("/v1/stations/:station_id", DeleteStation)
	r.Run(":8000")

}

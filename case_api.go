package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var records = readCsvFile("./full_data.csv")

func main() {
	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	fmt.Println(records)

}

func setupRoutes(r *gin.Engine) {
	r.GET("/cases/new/country/:country", route1)
	r.GET("/cases/total/country/:from_date", route2)

}

//Dummy function
func route1(c *gin.Context) {
	country, ok := c.Params.Get("country")
	date, ok := c.GetQuery("date")
	cases := getNewCases(records, country, date)
	if ok == false {
		res := gin.H{
			"error": "country is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}
	/*
		city := ""
	*/
	res := gin.H{
		"new_cases": cases,
		"country":   country,
		"date":      date,
		"count":     len(cases),
	}
	c.JSON(http.StatusOK, res)
}

//Dummy function
func route2(c *gin.Context) {

	date, ok := c.Params.Get("from_date")

	total_cases := getTotalCases(records, date)
	if ok == false {
		res := gin.H{
			"error": "date is missing",
			"date":  date,
		}
		c.JSON(http.StatusOK, res)
		return
	}
	/*
		city := ""
	*/
	res := gin.H{
		"total_cases": total_cases,
		"date":        date,
		"count":       len(total_cases),
	}
	c.JSON(http.StatusOK, res)
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func getNewCases(records [][]string, country string, date string) []string {

	var newCases = []string{}
	for i := 1; i < len(records); i++ {

		//fmt.Println(records[0][0], i)
		if records[i][1] == country {
			if records[i][0] == date {
				newCases = append(newCases, records[i][2])
			}

		}

	}
	return newCases
}

func getTotalCases(records [][]string, date string) []string {

	var total_cases = []string{}
	for i := 1; i < len(records); i++ {

		//fmt.Println(records[0][0], i)
		if records[i][0] == date {

			total_cases = append(total_cases, records[i][4])

		}

	}
	return total_cases
}

package main

// Import the libraries
import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Covid Structure
type CovidData struct {
	State        string
	Positive     int
	Negative     int
	Hospitalized int
	Recovered    int
	Deaths       int
}

// FirstRow adds the first row in a CSV file
func FirstRow() {
	rows := [][]string{
		{"State", "Positive", "Negative", "Hospitalized", "Recovered", "Deaths"},
	}

	// It opens, creates, and update a file
	file, err := os.OpenFile("pandemic.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// It writes into the file
	csvWriter := csv.NewWriter(file)
	for _, row := range rows {
		_ = csvWriter.Write(row)
	}

	csvWriter.Flush()
	file.Close()
}

// CSVData writes the pandemic data in a file
func CSVData(value1 string, value2, value3, value4, value5, value6 int) {
	rows := [][]string{
		{value1, strconv.Itoa(value2), strconv.Itoa(value3), strconv.Itoa(value4), strconv.Itoa(value5), strconv.Itoa(value6)},
	}

	file, err := os.OpenFile("pandemic.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	csvWriter := csv.NewWriter(file)
	for _, row := range rows {
		_ = csvWriter.Write(row)
	}

	csvWriter.Flush()
	file.Close()
}

// main gets the data from a website
// It transfers that data to Google maps
func main() {
	stateNames1 := []string{
		"AK", "AL", "AR", "AZ", "CA", "CO", "CT", "DE", "FL", "GA", "HI", "IA", "ID", "IL",
		"IN", "KS", "KY", "LA", "MA", "MD", "ME", "MI", "MN", "MO", "MS", "MT", "NC", "ND",
		"NE", "NH", "NJ", "NM", "NV", "NY", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN",
		"TX", "UT", "VA", "VT", "WA", "WI", "WV", "WY"}

	stateNames2 := []string{
		"Alaska", "Alabama", "Arkansas", "Arizona", "California", "Colorado", "Connecticut",
		"Delaware", "Florida", "Georgia", "Hawaii", "Iowa", "Idaho", "Illinois", "Indiana",
		"Kansas", "Kentucky", "Louisiana", "Massachusetts", "Maryland", "Maine", "Michigan",
		"Minnesota", "Missouri", "Mississippi", "Montana", "North Carolina", "North Dakota",
		"Nebraska", "New Hampshire", "New Jersey", "New Mexico", "Nevada", "New York", "Ohio",
		"Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota",
		"Tennessee", "Texas", "Utah", "Virginia", "Vermont", "Washington", "Wisconsin",
		"West Virginia", "Wyoming"}

	url := "https://covidtracking.com/api/states"
	method := "GET"

	// It sends the request to get the data from the above URL and method
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Do() sends the request and returns an HTTP response
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// It reads the data from the website
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(body))

	var specific []CovidData

	// It converts json body data into a struct
	// It fetches the data based on CovidData struct
	err = json.Unmarshal([]byte(body), &specific)
	if err != nil {
		log.Fatal(err)
	}

	// FirstRow() prints the first row of the data in a CSV file
	FirstRow()
	// It fetches all the states data and prints them
	for values := range specific {
		for i := 0; i < 50; i++ {
			if specific[values].State == stateNames1[i] {
				CSVData(stateNames2[i], specific[values].Positive, specific[values].Negative, specific[values].Hospitalized, specific[values].Recovered, specific[values].Deaths)
			}
		}
	}
}

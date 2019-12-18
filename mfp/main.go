package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var user string
var date string

// MacroData stores the parsed MFP macro data
type MacroData struct {
	Calories int
	Carbs    int
	Fat      int
	Protein  int
	Date     string
}

func main() {

	userName := flag.String("u", "default", "MFP Username")
	dateFlag := flag.String("d", "2019-12-12", "Request date (YYYY-MM-DD)")

	flag.Parse()

	makeRequest(*userName, *dateFlag)

}

// makeRequest will take the parsed flags and use them in a request to MFP.
// Following the request, the response body is parsed for the table and sorted into
// the MacroData struct.
func makeRequest(user string, date string) {

	// Create a new instance of the MacroData struct
	macros := new(MacroData)

	// Craft the URL
	mfpURL := "http://www.myfitnesspal.com/reports/printable_diary/" + user + "?from=" + date + "&to=" + date

	// Prepare new request
	req, err := http.NewRequest("GET", mfpURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Set User-Agent in the header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")

	// Make the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Close the response body when finished with it
	defer resp.Body.Close()

	// Load the response body with GoQuery
	contents, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Use GoQuery to parse and search the table for our relevent data selections
	contents.Find("tfoot").Each(func(i int, s *goquery.Selection) {
		s.Find("td").Each(func(indexstr int, rowHtml *goquery.Selection) {
			tdData := rowHtml.Text()

			// Format Calories
			if indexstr == 1 {
				iformat, err := strconv.Atoi(tdData)
				if err != nil {
					log.Fatal(err)
				}
				macros.Calories = iformat
			}

			// Format Carbs
			if indexstr == 2 {
				sf := strings.TrimRight(tdData, "g")
				iformat, err := strconv.Atoi(sf)
				if err != nil {
					log.Fatal(err)
				}
				macros.Carbs = iformat
			}

			// Format Fat
			if indexstr == 3 {
				sf := strings.TrimRight(tdData, "g")
				iformat, err := strconv.Atoi(sf)
				if err != nil {
					log.Fatal(err)
				}
				macros.Fat = iformat
			}

			// Format Protein
			if indexstr == 4 {
				sf := strings.TrimRight(tdData, "g")
				iformat, err := strconv.Atoi(sf)
				if err != nil {
					log.Fatal(err)
				}
				macros.Protein = iformat
			}

			// Include the date in the object
			macros.Date = date

		})
	})

	// Marshal our macros struct to JSON
	format, err := json.Marshal(macros)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(format))

}

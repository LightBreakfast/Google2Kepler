package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Records struct {
	TimelineObjects []struct {
		ActivitySegment struct {
			StartLocation struct {
				LatitudeE7  int `json:"latitudeE7"`
				LongitudeE7 int `json:"longitudeE7"`
				SourceInfo  struct {
					DeviceTag int `json:"deviceTag"`
				} `json:"sourceInfo"`
			} `json:"startLocation"`
			EndLocation struct {
				LatitudeE7  int `json:"latitudeE7"`
				LongitudeE7 int `json:"longitudeE7"`
				SourceInfo  struct {
					DeviceTag int `json:"deviceTag"`
				} `json:"sourceInfo"`
			} `json:"endLocation"`
			Duration struct {
				StartTimestampMs string `json:"startTimestampMs"`
				EndTimestampMs   string `json:"endTimestampMs"`
			} `json:"duration"`
			Distance     int    `json:"distance"`
			ActivityType string `json:"activityType"`
			Confidence   string `json:"confidence"`
			Activities   []struct {
				ActivityType string  `json:"activityType"`
				Probability  float64 `json:"probability"`
			} `json:"activities"`
			WaypointPath struct {
				Waypoints []struct {
					LatE7 int `json:"latE7"`
					LngE7 int `json:"lngE7"`
				} `json:"waypoints"`
			} `json:"waypointPath"`
		} `json:"activitySegment,omitempty"`
		PlaceVisit struct {
			Location struct {
				LatitudeE7  int    `json:"latitudeE7"`
				LongitudeE7 int    `json:"longitudeE7"`
				PlaceID     string `json:"placeId"`
				Address     string `json:"address"`
				Name        string `json:"name"`
				SourceInfo  struct {
					DeviceTag int `json:"deviceTag"`
				} `json:"sourceInfo"`
			} `json:"location"`
			Duration struct {
				StartTimestampMs string `json:"startTimestampMs"`
				EndTimestampMs   string `json:"endTimestampMs"`
			} `json:"duration"`
			PlaceConfidence string `json:"placeConfidence"`
			CenterLatE7     int    `json:"centerLatE7"`
			CenterLngE7     int    `json:"centerLngE7"`
		} `json:"placeVisit,omitempty"`
	} `json:"timelineObjects"`
}

func main() {

	searchDir := "Semantic Location History"

	fileList := []string{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.OpenFile("results.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := csv.NewWriter(f)

	w.Write([]string{"StartLat", "StartLong", "EndLat", "EndLong", "StartTime", "EndTime"})

	var totalrecords int

	for _, file := range fileList {
		if strings.Contains(file, ".json") == true {

			// Open our jsonFile
			jsonFile, err := os.Open(file)
			// if we os.Open returns an error then handle it
			if err != nil {
				fmt.Println(err)
			}
			// defer the closing of our jsonFile so that we can parse it later on
			defer jsonFile.Close()

			// read our opened xmlFile as a byte array.
			byteValue, _ := ioutil.ReadAll(jsonFile)

			var records Records

			// we unmarshal our byteArray which contains our
			// jsonFile's content into 'records' which we defined above
			json.Unmarshal(byteValue, &records)

			//fmt.Printf("Number Of Records to Parse: %v \n", len(records.TimelineObjects))
			totalrecords = totalrecords + len(records.TimelineObjects)

			// we iterate through every object in the records array
			// then write it out to the CSV File
			for i := 0; i < len(records.TimelineObjects); i++ {
				if records.TimelineObjects[i].ActivitySegment.StartLocation.LatitudeE7 != 0 {
					if records.TimelineObjects[i].ActivitySegment.StartLocation.LongitudeE7 != 0 {
						if records.TimelineObjects[i].ActivitySegment.EndLocation.LatitudeE7 != 0 {
							if records.TimelineObjects[i].ActivitySegment.EndLocation.LongitudeE7 != 0 {

								stimes64, err := strconv.ParseInt(records.TimelineObjects[i].ActivitySegment.Duration.StartTimestampMs, 10, 64)
								if err != nil {
									fmt.Println(err)
								}
								st := time.Unix(stimes64/1000, 0)

								etimes64, err := strconv.ParseInt(records.TimelineObjects[i].ActivitySegment.Duration.EndTimestampMs, 10, 64)
								if err != nil {
									fmt.Println(err)
								}
								et := time.Unix(etimes64/1000, 0)

								startlatstring := fmt.Sprintf("%f", float64(records.TimelineObjects[i].ActivitySegment.StartLocation.LatitudeE7)/1e7)
								startlongstring := fmt.Sprintf("%f", float64(records.TimelineObjects[i].ActivitySegment.StartLocation.LongitudeE7)/1e7)
								endlatstring := fmt.Sprintf("%f", float64(records.TimelineObjects[i].ActivitySegment.EndLocation.LatitudeE7)/1e7)
								endlongstring := fmt.Sprintf("%f", float64(records.TimelineObjects[i].ActivitySegment.EndLocation.LongitudeE7)/1e7)
								w.Write([]string{startlatstring, startlongstring, endlatstring, endlongstring, st.String(), et.String()})
							}
						}
					}
				}
			}
		}
	}
	fmt.Printf("Records Parsed: %v \n", totalrecords)
	w.Flush()
	fmt.Println("Complete!")
}


package main

import (
    "bufio"
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "os"
)


// Find way to print integers

type Top struct {
Data *Data `json:"data,omitempty"`
}

type Data struct {
    Enabled             string `json:"enabled"`
    Policy_id           string `json:"policy_id"`
    Name                string `json:"name"`
    Where_clause        string `json:"where_clause"`
    Type                string `json:"type"`
    Event_type          string `json:"event_type"`
    Select_value        string `json:"select_value"`
    Comparison          string `json:"comparison"`
    Critical_threshold *Critical `json:"critical_threshold,omitempty"`
}

type Critical struct {
    Value               string `json:"value"`
    Duration_minutes    string `json:"duration_minutes"`
    Time_function       string `json:"time_function"`
}

func main() {
// read from tmp csv
    csvFile, _ := os.Open("/tmp/gata.csv")
    reader := csv.NewReader(bufio.NewReader(csvFile))
// added struct to slice
    var condition []Top
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }
        condition = append(condition, Top{
           Data: &Data{
            Enabled: line[0],
            Policy_id:  line[1],
            Name:  line[2],
            Where_clause:  line[3],
            Type:  line[4],
            Event_type:  line[5],
            Select_value:  line[6],
            Comparison:  line[7],
            Critical_threshold: &Critical{
                Value:  line[8],
                Duration_minutes:  line[9],
                Time_function:  line[10],
            },



          },
        })
    }
    conditionJson, _ := json.Marshal(condition)
    fmt.Println(string(conditionJson))
}

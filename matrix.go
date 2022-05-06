package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Address struct {
	Date              string
	Location, Emotion string
	SubTotal          int
}

const DateInputLayout = "01/02/2006"

func timeParse(timeval string) string {
	tm, err := time.Parse(DateInputLayout, timeval)
	dateonly := tm.Format("20060102")
	if err != nil {
		fmt.Println(err)
	}
	return dateonly
}

func csvToMap(infilename string) map[string]int {
	infile, err := os.Open(infilename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Reading %s \n", infilename)
	}
	reader := csv.NewReader(infile)
	var dataMap = map[string]int{}

	count, ignoredline, startline := 0, 0, 1
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
		if count < startline {
			fmt.Printf("Ignoring line %d: %q\n", count+1, row)
			ignoredline++
		} else {
			datestr := timeParse(row[0])
			matrix := fmt.Sprintf("%s_%s_%s", datestr, row[1], row[2])
			_, ok := dataMap[matrix]
			if ok {
				dataMap[matrix]++
			} else {
				dataMap[matrix] = 1
			}
		}
		count++
	}
	return dataMap
}

func main() {
	start := time.Now()
	mapout := csvToMap("rando-input.csv")
	for k, v := range mapout {
		kstr := strings.Split(k, "_")
		fmt.Printf("%s => %d\n", kstr, v)
		//fmt.Printf("%s => %d\n", k, v)
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	//fmt.Println(mapout)

}

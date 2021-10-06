package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func main() {
	csvFile, err := os.Open("vaxthusgaserdata.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.Comma = ';'

	csvLines, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	res := industryMinMax(csvLines)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Branch", "Minsta utsläpp", "år minst utsläpp", "Max utsläpp", "år max utsläpp"})

	for _, r := range res {
		table.Append(r.Data())
	}
	//table.Render() // Send output

	fmt.Println("MY OWN TABLE")
	fmt.Printf("%s %s - %s       %s - %s \n", addEvenSpaces("Branch", 50), addEvenSpaces("Min", 5), addEvenSpaces("År", 10), addEvenSpaces("Max", 5), addEvenSpaces("År", 10))
	fmt.Println("----------------------------------------------------------------------------------------")
	for _, r := range res {
		fmt.Printf("%s %s -  %s     %s - %s \n", addEvenSpaces(r.Industry, 50), addEvenSpaces(strconv.Itoa(r.Min), 5), addEvenSpaces(r.MinYear, 10), addEvenSpaces(strconv.Itoa(r.Max), 5), addEvenSpaces(r.MaxYear, 10))
	}

}

func addEvenSpaces(str string, spaces int) string {
	addSpace := spaces - len(str)

	for i := 0; i < addSpace; i++ {
		str += " "
	}
	return str
}

func industryMinMax(csvLines [][]string) []IndustryMinMax {
	years := csvLines[0][1:]

	industriesMinMax := []IndustryMinMax{}

	for _, line := range csvLines[1:] {
		industry := IndustryMinMax{
			Industry: line[0],
		}
		for i, strVal := range line[1:] {
			val, _ := strconv.Atoi(strVal)
			if val > industry.Max {
				industry.Max = val
				industry.MaxYear = years[i]
			}
			if industry.Min == 0 || val < industry.Min {
				industry.Min = val
				industry.MinYear = years[i]
			}
		}
		industriesMinMax = append(industriesMinMax, industry)
	}
	return industriesMinMax
}

type IndustryMinMax struct {
	Industry string
	Min      int
	MinYear  string
	Max      int
	MaxYear  string
}

func (i IndustryMinMax) Data() []string {

	return []string{
		i.Industry,
		strconv.Itoa(i.Min),
		i.MinYear,
		strconv.Itoa(i.Max),
		i.MaxYear}
}

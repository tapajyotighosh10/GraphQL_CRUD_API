package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f, err := excelize.OpenFile("MyData.xlsx")
	if err != nil {
		fmt.Println(err)
	}
// Get value from cell by given worksheet name and axis.

	rows:= f.GetRows("Sheet1")

	
	fmt.Println("Excel file informations are")
	for _, row := range rows {
		for _, col := range row{
			fmt.Print(col, "\t")
		}
		fmt.Println()
	}
}

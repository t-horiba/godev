package main

// go get github.com/360EntSecGroup-Skylar/excelize

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	cell1, err := f.GetCellValue("Sheet1", "A1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell1)

	cell2, err := f.GetCellValue("Sheet1", "B1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell2)

	cell3, err := f.GetCellValue("Sheet1", "E1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell3)

	cell4, err := f.GetCellValue("Sheet1", "G1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell4)

	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

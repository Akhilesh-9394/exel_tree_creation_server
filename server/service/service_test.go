package service

import (
	"fmt"
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func BenchmarkGetDataFromStruct(b *testing.B) {
	f, _ := excelize.OpenFile("./assets/students.xlsx")
	rows, _ := f.Rows("Sheet1")
	fmt.Println(rows)
}

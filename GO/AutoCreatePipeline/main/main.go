package main

import (
	excelTool "AutoCreatePipeline/excel"
	"fmt"
	"log"
	"os"
)

var writeExcel excelTool.Excel

func main() {
	cells1 := map[string]any{"A1": "1", "A2": "2"}
	cells2 := map[string]any{"A3": ""}
	sheetNames := make(map[string]map[string]any)
	sheetNames["sheet1"] = cells1
	sheetNames["sheet2"] = cells2
	current_path, err := os.Getwd()
	if err != nil {
		log.Println("get current path failed")
	}
	log.Println(current_path)
	writeExcel.ExcelPath = fmt.Sprintf("%s/build", current_path)
	writeExcel.ExcelName = "test.xlsx"
	writeExcel.SheetNames = sheetNames
	log.Println(writeExcel)
	//excelTool.WriteExcel(writeExcel)
	excelTool.ReadToExcel("cell", writeExcel)
	//a := business.GetEnvKey("22868", business.CodingToken)
	//log.Println(reflect.TypeOf(a), a)
	//for i := 0; i < len(a); i++ {
	//	log.Println(a[i])
	//}

}

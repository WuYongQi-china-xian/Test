package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
)

//type Cells map[string]any

type Excel struct {
	ExcelName  string
	ExcelPath  string
	SheetNames map[string]map[string]any
}
type Excels []Excel

func WriteExcel(excel Excel) {
	if excel.ExcelPath != "" {
		err := os.MkdirAll(excel.ExcelPath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
			panic(fmt.Sprintf("create path %s failed", excel.ExcelPath))
		}
		log.Println(fmt.Sprintf("create path %s success", excel.ExcelPath))
	}

	f := excelize.NewFile()
	for sheetName, cells := range excel.SheetNames {
		// 创建一个工作表
		f.NewSheet(sheetName)
		for zuobiao, value := range cells {
			// 设置单元格的值
			f.SetCellValue(sheetName, zuobiao, value)
		}
	}

	// 根据指定路径保存文件
	if err := f.SaveAs(fmt.Sprintf("%s/%s", excel.ExcelPath, excel.ExcelName)); err != nil {
		log.Fatal(err)
		panic(fmt.Sprintf("save %s/%s excel failed", excel.ExcelPath, excel.ExcelName))
	}
	log.Println(fmt.Sprintf("save %s/%s success", excel.ExcelPath, excel.ExcelName))
}

func CellsContains(array []string, val string) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

func ReadToExcel(method string, excel Excel) Excel {
	f, err := excelize.OpenFile(fmt.Sprintf("%s/%s", excel.ExcelPath, excel.ExcelName))
	if err != nil {
		log.Println(err)
		panic(fmt.Sprintf("read %s/%s failed", excel.ExcelPath, excel.ExcelName))
	}
	log.Println(fmt.Sprintf("read %s/%s success", excel.ExcelPath, excel.ExcelName))
	// 获取工作表中指定单元格的值
	if method == "cell" {
		for sheetName, cells := range excel.SheetNames {
			for zuobiao, _ := range cells {
				// 获取工作表中指定单元格的值
				cells[zuobiao], err = f.GetCellValue(sheetName, zuobiao)
				if err != nil {
					log.Fatal(err)
					panic(fmt.Sprintf("read %s cell failed", zuobiao))
				}
			}
		}

	}

	// 获取 Sheet 上所有单元格
	if method == "sheet" {
		for sheetName, cells := range excel.SheetNames {
			// 获取 Sheet 上所有单元格
			rows, err := f.GetRows(sheetName)
			if err != nil {
				log.Println(fmt.Sprintf("read rows by sheetName %s failed", sheetName))
			}
			// i代表列号,j代表行号
			i, j := 1, 1
			for _, row := range rows {
				for _, colCell := range row {
					// 行号列号转化为坐标
					zuobiao, err := excelize.CoordinatesToCellName(i, j)
					if err != nil {
						log.Fatal(fmt.Sprintf("(%s,%s) to CellName failed", i, j))
						panic(fmt.Sprintf("(%s,%s) to CellName failed", i, j))
					}
					log.Println(fmt.Sprintf("colCell:%s zuobiao:%s (%s,%s)", colCell, zuobiao, i, j))
					cells[zuobiao] = colCell
					//列号自加，横行读取
					i++
				}
				// 列号回到第1行
				i = 1
				// 行号自加，竖向读取
				j++
			}
			log.Println(fmt.Sprintf("%s all cells: %v", sheetName, cells))
		}
	}
	log.Println(fmt.Sprintf("excel: %v", excel))
	return excel
}

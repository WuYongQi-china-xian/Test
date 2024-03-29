package business

import (
	"AutoCreatePipeline/excel"
	"fmt"
	"log"
	"os"
)

func GenerateExcel(templateIdList []string, token string) {
	var generateExcel excel.Excel
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatal("get current path failed")
	}
	generateExcel.ExcelPath = fmt.Sprintf("%s/build", rootPath)
	generateExcel.ExcelName = InOutExcelName
	// 记录生成流水线ID在sheet页中的坐标
	sheetNamesMap := make(map[string]map[string]any)
	//记录环境变量写完后的列值
	lieBiao := 0
	for _, templateId := range templateIdList {
		cellsMap := make(map[string]any)
		envKeyList := GetEnvKey(templateId, token)
		for i := 0; i < len(envKeyList); i++ {
			lieBiao = i + 1
			zuobiao := excel.GetZuoBiao(lieBiao, 1)
			cellsMap[zuobiao] = envKeyList[i]
		}
		zuobiao := excel.GetZuoBiao(lieBiao+1, 1)
		cellsMap[zuobiao] = GudingZiDuan1
		zuobiao = excel.GetZuoBiao(lieBiao+2, 1)
		cellsMap[zuobiao] = GudingZiDuan2
		zuobiao = excel.GetZuoBiao(lieBiao+3, 1)
		cellsMap[zuobiao] = GudingZiDuan3
		sheetNamesMap[templateId] = cellsMap
		// 列标归0
		lieBiao = 0
	}
	generateExcel.SheetNames = sheetNamesMap
	log.Println(fmt.Sprintf("generateExcel: %v", generateExcel))
	excel.WriteExcel(generateExcel)

}

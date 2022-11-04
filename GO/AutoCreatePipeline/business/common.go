package business

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

func ArrayContains(array []string, val string) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

func IsReg(reg, content string) bool {
	//解析正则表达式，如果成功返回解释器
	regCompile := regexp.MustCompile(reg)
	if regCompile == nil { //解释失败，返回nil
		log.Fatal("regexp err")
		return false
	}
	//根据规则提取关键信息
	result := regCompile.FindAllStringSubmatch(content, -1)
	log.Println(fmt.Sprintf("result = %s", result))
	if result != nil {
		return true
	} else {
		return false
	}
}

func JsonToMap(str string) map[string]string {
	//var strJson bytes.Buffer
	//_ = json.Indent(&strJson, []byte(str), "", "    ")
	//log.Println(fmt.Sprintf("strJson:%s", strJson.String()))
	var tempMap map[string]string

	err := json.Unmarshal([]byte(str), &tempMap)

	if err != nil {
		panic(err)
	}

	return tempMap
}

package business

import (
	"AutoCreatePipeline/excel"
	"AutoCreatePipeline/sdk"
	"AutoCreatePipeline/sdk/codingInterface"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
根excel填写的信息，按模板创建流水线
name: 流水线线名称 "自动建立流水线测试"
templateId: 模板id 48188
repoId: coding登记的代码仓库id，可通过仓库信息接口和工蜂url查询到 26294
projectId: coding项目id 2197
userEnv: 用户自定义环境变量
token: coding令牌
*/
func createPipeline(name, token string, templateId, repoId, projectId int, userEnv map[string]string) string {
	var pipelineReq codingInterface.CreatePipelineReq
	pipelineReq.Name = name
	pipelineReq.TemplateID = templateId
	pipelineReq.CodeSourceType = "git"
	pipelineReq.RepoID = repoId
	pipelineReq.ConfigType = 0 // 请求要求写死为0
	pipelineReq.Customize.TriggerTemplate = []string{"trigger_git"}
	pipelineReq.PipelineTags = []any{}
	pipelineReq.Tags = []any{}
	pipelineReq.PipelineJSON.Version = "2.0"
	pipelineReq.PipelineJSON.Stages = []any{}
	pipelineReq.ProjectId = projectId
	pipelineReq.UserEnv = userEnv
	pipelineInfo := sdk.CreatePipeline(pipelineReq, token)
	pipelineId := pipelineInfo.ID
	pipelineUrl := fmt.Sprintf("%s/p/PipeLineDemo/ci/job/%s/build/current", sdk.CodingRootApi, strconv.Itoa(pipelineId))
	log.Println(fmt.Sprintf("pipelineUrl: %s", pipelineUrl))
	return pipelineUrl
}

func CreatePipeByExcel(templateIdList []string, projectId, token string) {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatal("get current path failed")
	}
	InOutPutPath := fmt.Sprintf("%s/build", rootPath)
	shouhangMap := make(map[int]string) //记录首行和列的关系
	feishouhhangMap := make(map[string]any)
	for _, templateId := range templateIdList {
		rows, f := excel.ReadExcel(InOutPutPath, InOutExcelName, templateId)
		// i代表列号,j代表行号
		i, j := 1, 1
		for _, row := range rows {
			for _, colCell := range row {
				fmt.Print(colCell, "\t")
				if j == 1 {
					shouhangMap[i] = colCell
				} else {
					feishouhhangMap[shouhangMap[i]] = colCell
				}
				//列号自加，横行读取
				i++
			}
			// 首行不能处理业务
			if j != 1 {
				// 例用数据处理业务,将每行中的数据拆成固定映射关系和环境变量映射关系
				envMap := make(map[string]string)
				gudingMap := make(map[string]string)

				for k, v := range feishouhhangMap {
					if k == GudingZiDuan2 || k == GudingZiDuan1 || k == GudingZiDuan3 {
						// 非环境变量参数
						gudingMap[k] = fmt.Sprint(v)
					} else {
						// 环境变量参数
						envMap[k] = fmt.Sprint(v)
					}
				}
				// 创建流水线
				repoId := GetRepoId(projectId, fmt.Sprint(gudingMap[GudingZiDuan1]), token)
				templateIdInt, err := strconv.Atoi(templateId)
				if err != nil {
					panic(fmt.Sprintf("str %s to int failed", templateId))
				}
				projectIdInt, err := strconv.Atoi(projectId)
				if err != nil {
					panic(fmt.Sprintf("str %s to int failed", projectId))
				}
				pipelineUrl := createPipeline(gudingMap[GudingZiDuan2], token, templateIdInt, repoId, projectIdInt, envMap)
				writeZuobiao := excel.GetZuoBiao(i, j)
				// 回写结果
				err = f.SetCellHyperLink(templateId, writeZuobiao, pipelineUrl, "External")
				// 为单元格设置字体和下划线样式
				style, err := f.NewStyle(`{"font":{"color":"#1265BE","underline":"single"}}`)
				err = f.SetCellStyle(templateId, writeZuobiao, writeZuobiao, style)
				err = f.SetCellValue(templateId, writeZuobiao, pipelineUrl)
				if err != nil {
					log.Fatal(fmt.Sprintf("write sheetname:%s zuobiao:%s value:%s faied", templateId, writeZuobiao, pipelineUrl))
				}
				if err := f.SaveAs(fmt.Sprintf("%s/%s", InOutPutPath, InOutExcelName)); err != nil {
					log.Fatal(fmt.Sprintf("save as %s/%s failed: %v", InOutPutPath, InOutExcelName, err))
				}
				// 清理缓存
				envMap = nil
				gudingMap = nil

			}

			// 列号回到第1行
			i = 1
			// 行号自加，竖向读取
			j++
		}
		// 下个sheet页清空缓存
		shouhangMap = nil
		feishouhhangMap = nil
	}
	log.Println("所有流水线创建完毕")
}

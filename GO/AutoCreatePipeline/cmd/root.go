package cmd

import (
	"AutoCreatePipeline/business"
	"AutoCreatePipeline/yamlTool"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"reflect"
	"strings"
)

var templateIds string
var token string
var compPipelineIDJson string
var projectId string

// rootCmd 代表没有调用子命令时的基础命令
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "流水线搭建小工具",
	Long: `download:1.通过流水线模板ID下载环境变量名称到excel表格中，为自动创建流水线做准备，多个模板id请用英文逗号分隔
create:2.填写excel的内容，根据excel中的sheet页搭建流水线`,
	Args: cobra.ExactArgs(1),
	// 如果有相关的 action 要执行，请取消下面这行代码的注释
	// Run: func(cmd *cobra.Command, args []string) { },
}

//.\go_build_main_go.exe download --templateids "模板id1;模板idn" --token "coding令牌"
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "通过流水线模板ID下载环境变量名称到excel表格中，为自动创建流水线做准备，多个模板id请用英文逗号分隔",
	Long:  `通过流水线模板ID下载环境变量名称到excel表格中，为自动创建流水线做准备，多个模板id请用英文逗号分隔`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("start download" + strings.Join(args, " "))
		if templateIds == "" || token == "" {
			panic("templateIds和token不能为空")
		}
		// 匹配到 数字和英文逗号以外的字符
		if business.IsReg("[^\\d,]", templateIds) {
			panic(fmt.Sprintf("templateIds:%s 格式不正确，只能用英文逗号和数字", templateIds))
		}
		templateIdsList := strings.Split(templateIds, ",")
		business.GenerateExcel(templateIdsList, token)
	},
}

//.\go_build_main_go.exe create --templateids "模板id1;模板idn" --token "coding令牌" --projectid "数字id"
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "填写excel的内容，根据excel中的sheet页搭建流水线",
	Long:  `填写excel的内容，根据excel中的sheet页搭建流水线`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("start create: " + strings.Join(args, " "))
		if templateIds == "" || token == "" {
			panic("templateIds和token不能为空")
		}
		// 匹配到 数字和英文逗号以外的字符
		if business.IsReg("[^\\d,]", templateIds) {
			panic(fmt.Sprintf("templateIds:%s 格式不正确，只能用英文逗号和数字", templateIds))
		}
		templateIdsList := strings.Split(templateIds, ",")
		business.CreatePipeByExcel(templateIdsList, projectId, token)

	},
}

//.\go_build_main_go.exe generateScan --compidjson "{'a':'1','b':'2'}"
var generateScanCmd = &cobra.Command{
	Use:   "generateScan",
	Short: "填写组件名和流水线id生成批量扫描流水线的yaml",
	Long:  `填写组件名和流水线id生成批量扫描流水线的yaml，如参格式要为"{'组件名称':'流水线id','a':'1','b':'2'}"`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("start generateScan: " + strings.Join(args, " "))
		log.Println(fmt.Sprintf("compPipelineIDJson:%s", compPipelineIDJson))
		log.Println(fmt.Sprintf("type:%s", reflect.TypeOf(compPipelineIDJson)))
		guanxiJsonStr := strings.Replace(compPipelineIDJson, "'", "\"", -1)
		log.Println(fmt.Sprintf("guanxiJsonStr:%s", guanxiJsonStr))
		guanxi := business.JsonToMap(guanxiJsonStr)
		log.Println(fmt.Sprintf("guanxi:%v", guanxi))
		yamlTool.GetBatchScanTemplateYml(guanxi)
	},
}

//.\go_build_main_go.exe generatePack --compidjson "{'a':'1','b':'2'}"
var generatePackCmd = &cobra.Command{
	Use:   "generatePack",
	Short: "填写组件名和流水线id生成批量编译流水线的yaml",
	Long:  `填写组件名和流水线id生成批量编译流水线的yaml，如参格式要为"{'组件名称':'流水线id','a':'1','b':'2'}"`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("start generatePack: " + strings.Join(args, " "))
		log.Println(fmt.Sprintf("compPipelineIDJson:%s", compPipelineIDJson))
		log.Println(fmt.Sprintf("type:%s", reflect.TypeOf(compPipelineIDJson)))
		guanxiJsonStr := strings.Replace(compPipelineIDJson, "'", "\"", -1)
		log.Println(fmt.Sprintf("guanxiJsonStr:%s", guanxiJsonStr))
		guanxi := business.JsonToMap(guanxiJsonStr)
		log.Println(fmt.Sprintf("guanxi:%v", guanxi))
		yamlTool.GetBatchPackTemplateYml(guanxi)
	},
}

func init() {
	downloadCmd.PersistentFlags().StringVar(&templateIds, "templateids", "", "流水线模板id的值多个用英文逗号分隔")
	downloadCmd.PersistentFlags().StringVar(&token, "token", "", "流水线项目令牌")
	createCmd.PersistentFlags().StringVar(&templateIds, "templateids", "", "流水线模板id的值多个用英文逗号分隔")
	createCmd.PersistentFlags().StringVar(&token, "token", "", "流水线项目令牌")
	createCmd.PersistentFlags().StringVar(&projectId, "projectid", "", "coding项目id")
	generateScanCmd.PersistentFlags().StringVar(&compPipelineIDJson, "compidjson", "", "要json格式字符串，如{\"组件名\": \"流水线id\",\"a\": 1}")
	generatePackCmd.PersistentFlags().StringVar(&compPipelineIDJson, "compidjson", "", "要json格式字符串，如{\"组件名\": \"流水线id\",\"a\": 1}")
}

// Execute 将所有子命令添加到root命令并适当设置标志。
// 这由 main.main() 调用。它只需要对 rootCmd 调用一次。
func Execute() {
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(generateScanCmd)
	rootCmd.AddCommand(generatePackCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

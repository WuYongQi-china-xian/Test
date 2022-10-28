package cmd

import (
	"AutoCreatePipeline/business"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

var templateIds string
var token string

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

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "填写excel的内容，根据excel中的sheet页搭建流水线",
	Long:  `填写excel的内容，根据excel中的sheet页搭建流水线`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, " "))
	},
}

func init() {
	downloadCmd.PersistentFlags().StringVar(&templateIds, "templateids", "", "流水线模板id的值多个用英文逗号分隔")
	downloadCmd.PersistentFlags().StringVar(&token, "token", "", "流水线项目令牌")
}

// Execute 将所有子命令添加到root命令并适当设置标志。
// 这由 main.main() 调用。它只需要对 rootCmd 调用一次。
func Execute() {
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(createCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

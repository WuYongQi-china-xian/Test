package yaml

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type batchScanTemplate struct {
	Stages []struct {
		Stage string `yaml:"stage"`
		Tasks []Task
		If    string `yaml:"if"`
	} `yaml:"stages"`
	Version string `yaml:"version"`
	Worker  struct {
		Strategy string        `yaml:"strategy"`
		Label    string        `yaml:"label"`
		Tools    []interface{} `yaml:"tools"`
		Docker   struct {
			Image string `yaml:"image"`
		} `yaml:"docker"`
	} `yaml:"worker"`
	Env struct {
		SCANBRANCH struct {
			Value   string `yaml:"value"`
			Type    string `yaml:"type"`
			Require bool   `yaml:"require"`
		} `yaml:"SCAN_BRANCH"`
		MRCOMPONENTS struct {
			Value   string `yaml:"value"`
			Type    string `yaml:"type"`
			Require bool   `yaml:"require"`
			Option  string `yaml:"option"`
		} `yaml:"MR_COMPONENTS"`
	} `yaml:"env"`
	Finally struct {
		Success struct {
			Cmds []struct {
				Plugin string `yaml:"plugin"`
				Params struct {
					TextMode string `yaml:"text_mode"`
					Text     string `yaml:"text"`
				} `yaml:"params"`
				Label string `yaml:"label"`
			} `yaml:"cmds"`
		} `yaml:"success"`
		Failure struct {
			Cmds []struct {
				Plugin string      `yaml:"plugin"`
				Params interface{} `yaml:"params,omitempty"`
			} `yaml:"cmds"`
		} `yaml:"failure"`
	} `yaml:"finally"`
	Setup            []string `yaml:"setup"`
	CodeCheckoutType string   `yaml:"code_checkout_type"`
	Manual           struct {
	} `yaml:"manual"`
	MaximumBuilds int `yaml:"maximum_builds"`
}

type Task struct {
	Task  string   `yaml:"task"`
	Cmds  []Cmd    `yaml:"cmds"`
	Temps []string `yaml:"temps,omitempty"`
	If    string   `yaml:"if"`
}

type Cmd struct {
	Plugin string `yaml:"plugin"`
	Params struct {
		Pipeline    string `yaml:"pipeline"`
		Sync        bool   `yaml:"sync"`
		Branch      string `yaml:"branch"`
		TriggerBase string `yaml:"trigger_base"`
	} `yaml:"params"`
	Label string `yaml:"label"`
}

func mapToSlice(m map[string]string) (k, v []string) {
	kArray := make([]string, 0, len(m))
	vArray := make([]string, 0, len(m))
	for k, v := range m {
		kArray = append(kArray, k)
		vArray = append(vArray, v)
	}
	return kArray, vArray
}

/*
基于batchScanTemplate.yaml和组件与流水线id的map重复设置代码扫描task生产新的yaml
    compTaskIdMaps := make(map[string]string)
	compTaskIdMaps["server_manager"] = "937182"
	compTaskIdMaps["stat_gateway"] = "937179"
	compTaskIdMaps["stream_manager"] = "937159"
	compTaskIdMaps["multi_bitrate_adaptive"] = "937145"
	compTaskIdMaps["customer_task_center"] = "937142"
	compTaskIdMaps["media_master"] = "937138"
	compTaskIdMaps["source_master"] = "937132"
	compTaskIdMaps["media_conf"] = "937131"
	compTaskIdMaps["source_conf"] = "937129"
	compTaskIdMaps["media_interface"] = "937126"
	compTaskIdMaps["source_interface"] = "937124"
	compTaskIdMaps["source_dispatch_agent"] = "937122"
	compTaskIdMaps["source_strategyer"] = "937094"
	compTaskIdMaps["live_tlss"] = "936875"
	yamlTool.GetBatchScanTemplateYml(compTaskIdMaps)
*/
func GetBatchScanTemplateYml(compTaskIdMaps map[string]string) {
	compArray, _ := mapToSlice(compTaskIdMaps)
	mrComp := strings.Join(compArray, ";")
	log.Println(fmt.Sprintf("MR_COMPONENTS:%s", mrComp))
	// 获取工作空间根路径
	root_path, err := os.Getwd()
	if err != nil {
		log.Fatal("get root path failed")
	}
	// 读取基准yaml转为字节
	file, err := ioutil.ReadFile(fmt.Sprintf("%s/templates/batchScanTemplate.yaml", root_path))
	if err != nil {
		log.Fatal(err)
	}
	// 字节写入struct
	var bst batchScanTemplate
	err2 := yaml.Unmarshal(file, &bst)

	if err2 != nil {
		log.Fatal(err2)
	}
	// 细化struct，方便修改部门struct的值
	var task Task
	var tasks []Task
	var cmd Cmd
	var cmds []Cmd

	// 将词典中的循环写入细化后的strut，方便后面用来重新赋值给读出来的的struct
	for comp, pipelineid := range compTaskIdMaps {
		task.Task = fmt.Sprintf("%s代码扫描", comp)
		task.If = fmt.Sprintf("${MR_COMPONENTS} ~= %s", comp)
		task.Temps = []string{}
		cmd.Plugin = "pipeline_start"
		cmd.Params.Pipeline = pipelineid
		cmd.Params.Sync = true
		cmd.Params.Branch = "$SCAN_BRANCH"
		cmd.Params.TriggerBase = "branch"
		cmd.Label = "启动流水线"
		cmds = append(cmds, cmd)

		task.Cmds = cmds
		tasks = append(tasks, task)
		cmds = nil
	}
	log.Println(fmt.Sprintf("newTasks:%v", tasks))
	// 将读出的原始struct改值
	bst.Stages[0].Tasks = tasks
	bst.Env.MRCOMPONENTS.Option = mrComp
	log.Println(fmt.Sprintf("bst: %v", bst))
	// 将struct变为yaml数据
	bstYaml, err := yaml.Marshal(&bst)
	if err != nil {
		log.Fatal("Error while Marshaling. %v", err)
	}
	outputFile := fmt.Sprintf("%s/build/newBatchScanTemplate.yaml", root_path)
	// 将yaml数据写入文件
	err = ioutil.WriteFile(outputFile, bstYaml, 0644)
	if err != nil {
		panic("Unable to write data into the file")
	}
}

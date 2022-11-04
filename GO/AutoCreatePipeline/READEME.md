# 流水线小工具
- .\go_build_main_go.exe generatePack --compidjson "{'组件名':'流水线id'','a':'1','b':'2'}"
    * 根据组件名称和流水线id的关系生成一个CLE能用的批量编译流水线yaml
- .\go_build_main_go.exe generateScan --compidjson "{'组件名':'流水线id'','a':'1','b':'2'}"
    * 根据组件名称和流水线id的关系生成一个能用的批量扫描流水线yaml
- .\go_build_main_go.exe download --templateids "模板id1,模板id2,模板idn" --token "coding令牌"
    * 根据coding的模板id，生成一个带有环境变量信息的excel用于收集信息
- .\go_build_main_go.exe create --templateids "模板id1,模板id2,模板idn" --token "coding令牌" --projectid "coding项目id"
    * 根据上个命令收集信息的excel批量创建流水线
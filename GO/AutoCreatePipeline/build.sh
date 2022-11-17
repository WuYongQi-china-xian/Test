#!/bin/bash

# 约定根目录执行

#循环数组执行shell命令，如果shell命令执行失败，停止继续执行，返回执行失败的错误码
#cmds=("make spp_redis_ping CODE_VERSION=QCLOUD"
#	"cd release/spp_redis_ping/client/"
#	"tar -czf ${ZHIYUN_PACKAGE_NAME}.tgz bin/"
#	"mv -f ${ZHIYUN_PACKAGE_NAME}.tgz $QCI_WORKSPACE"
#	"cd $QCI_WORKSPACE")
#exc_cmd_check_ret "${cmds[@]}"	
function exc_cmd_check_ret(){
    for cmd in "$@"
    do
		${cmd};ret=$?
		if [[ $ret -eq 0 ]];then
		    echo "${cmd}执行成功；返回值为${ret}"
		else
		    echo "${cmd}执行失败；返回值为${ret}"
			return $ret
		fi
    done
}

function main(){
    # mac or linux
    cmds=("go build -o ./go_build_main_go ./main/main.go"
        "chmod u+x ./go_build_main_go"
        "zip -r SmallTool.zip ./templates/* ./go_build_main_go")
    exc_cmd_check_ret "${cmds[@]}"
}

main
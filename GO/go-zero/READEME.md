# goctl常用命令
- goctl api new xxx
- goctl api go -api xxx.api -dir xxx

# 启动命令
- cd 项目根目录
- go run robot/robot.go -f robot/etc/robot-api.yaml

# 访问命令
- curl -H "Content-Type: application/json" -X POST -d '{"msgType":"str","eventType":"str","chatType":"str","msgContent":"a","chatId":"str","botKey":"str","hookUrl":"str","botName":"str","userName":"str","msgId":"str","chatInfoUrl":"str","responseUrl":"str"}'  http://localhost:8888/robot/api/callback

# 参考资料：
- [go-zero](https://github.com/zeromicro/go-zero)
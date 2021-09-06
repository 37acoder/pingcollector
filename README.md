# Ping Collector
ping收集器是一个使用Golang With InfluxDB的网络延迟检测工具
## /service
该目录下是主要功能逻辑的实现

## main.go 
是程序的入口，初始化了Service以及处理Service之间的依赖关系

## notice
记得配置一下config.json，指定influxdb的ip port和token
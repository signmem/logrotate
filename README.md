# logrotate

## 作用
> 支持自定义日志格式
> 支持对日志进行 logrotate 截断并备份
> 支持自定义备份数量


# usage
## 备份相关

> cfg.json 为配置文件
> LogMaxAge 单位秒 - 可保留的日志最大时间长度
> LogRotateAge 单位秒 - 多少秒执行一次 logrotate 操作


## 日志相关

> 日志文件格式 - 修改 g/logger.go 文件 fmt.Sprintf("%s.%s", logfile, "修改你的格式")
> 日志内容格式 - 修改 g/logger.go WriterSink() 函数

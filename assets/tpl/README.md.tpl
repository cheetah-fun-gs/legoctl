# {{.ProjectName}}

这是一个[goso](https://github.com/cheetah-fun-gs/goso)项目，可以使用[goso-cli](https://github.com/cheetah-fun-gs/goso-cli)工具来管理。
要求go1.13。  

# 添加接口

## 请求和响应符合规范的接口

1. 参考```internal/biz/handlers/common/ping.go```创建自己的handler文件，路径格式```internal/biz/handlers/xxx/xxx.go```  
2. 使用[goso-cli](https://github.com/cheetah-fun-gs/goso-cli)的```gen```命令生成框架代码

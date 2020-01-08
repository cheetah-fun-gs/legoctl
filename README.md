# legoctl
lego管理工具，要求go1.13

# 下载
```bash
go get github.com/cheetah-fun-gs/legoctl
```

# 创建项目
```
legoctl project new <path> [-t template-name] [-p package-name]
```

# 模板说明
- [基础项目模板 lego](assets/templates/lego/README.md)

# 命令说明
```
Usage:
  legoctl project [flags]
  legoctl project [command]

Available Commands:
  gen         生成代码
  new         创建项目

Flags:
  -h, --help   help for project

Global Flags:
  -p, --package-name string    包名称：默认目录名
  -t, --template-name string   模板名称 (default "lego")
  -v, --verbose                是否输出详细日志

Use "legoctl project [command] --help" for more information about a command.
```

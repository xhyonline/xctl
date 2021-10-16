# xctl
xctl 是一款 `Golang` 轻量级命令行代码生成器

它主要用来创建项目使用的,创建后的项目更像一款框架。

并且项目中自带 github action 配置文件

非常适合个人开发者使用

## 安装命令

`go install github.com/xhyonline/xctl@v0.1.20211017`

## Golang 版本要求

Go 1.16 以上

## 运行示例如下
```
$ xctl create
必填参数不能为空:  --with-name  --with-mod
你可以根据自己的需求创建一个项目,示例如下

Usage:
   create [flags]

Examples:
xctl create --with-name myapp --with-mod github.com/xhyonline/myapp

Flags:
  -h, --help                help for create
      --with-githubAction   是否初始化 github action 集成
      --with-mod string     必填项 初始化 go mod 例如: github.com/myapp
      --with-mysql          是否使用 mysql 数据库
      --with-name string    必填项 应用名称,例如:myapp,
      --with-redis          是否使用 redis 缓存
```

## 使用示例
执行一下命令

```
xctl create --with-name myapp --with-mod github.com/xhyonline/myapp --with-mysql --with-redis --with-githubAction
```

**你就能获得如下结构的 HTTP 项目,并且它自身就支持优雅停止**

```
myapp/
|-- component
|   |-- init.go
|   |-- logger.go
|   |-- mysql.go
|   `-- redis.go
|-- configs
|   |-- common
|   |   |-- mysql.toml
|   |   `-- redis.toml
|   `-- init.go
|-- .github
|   `-- workflows
|       `-- github-action.yml # CI-构建配置
|-- .golangci.yml # 代码质量检测
|-- go.mod
|-- go.sum
|-- internal
|   `-- http.go
|-- main.go # 入口
|-- middleware # 中间件
|   `-- cors.go # 跨域配置
`-- router # 路由
    `-- router.go

```

**其中 github-action.yml** 开箱即用,配置如下,你可以根据自行需求进行更改
```
name: CI构建
on:
  push:
    branches: [ main,master ]
  pull_request:
    branches: [ main ,master ] # merge到main分支时触发部署

env:
  APP_NAME: x-http # 给 APP 起一个名字

jobs:
  build-and-deploy:
    runs-on: ubuntu-latest
    steps:
      - name: 检出代码
        uses: actions/checkout@master

      - name: 设置环境 Golang 环境
        uses: actions/setup-go@v2
        with:
          go-version: 1.17

      - name: 代码质量检测
        uses: golangci/golangci-lint-action@v2
        with:
          # Optional: version of golangci-lint to use in form of v1.2 or v1.2.3 or `latest` to use the latest version
          version: v1.29

      - name: 构建 BuiLd
        run: |
          export GOPROXY=https://goproxy.io,direct
          go build -o app

      - name: upx 压缩二进制文件
        uses: crazy-max/ghaction-upx@v1
        with:
          version: latest
          files: |
            app
          args: -fq

      - name: 同步文件
        uses: burnett01/rsync-deployments@5.1
        with:
          switches: -avzr --delete
          path: ./app
          remote_path: /micro-server/$APP_NAME # 发布到远程主机,当然你需要自己创建 /micro-server 目录 $APP_NAME 是全局的变量
          remote_host: ${{ secrets.Host }}
          remote_port: 22
          remote_user: root
          remote_key: ${{ secrets.DeploySecret }} # 请使用 ssh-keygen -t rsa 生成秘钥对,然后将公钥拷贝到要操纵的目标器的/root/.ssh/authorized_keys里,再把私钥黏贴到 github 后台的secret里

      - name: 执行重启命令
        uses: appleboy/ssh-action@master
        with:
          host: ${{ secrets.Host }}
          username: root
          key: ${{ secrets.DeploySecret }}
          port: 22
          script: | # 请自行在这里执行应用的重启命令
            pwd
            ls /micro-server

      - name: 构建结果通知
        uses: zzzze/webhook-trigger@master
        if: always()
        with:
          data: "{'event_type':'build-result','status':'${{ job.status }}',
          'repository':'${{ github.repository }}','job':'${{ github.job }}',
          'workflow':'${{ github.workflow }}'}"
          webhook_url: ${{ secrets.WebHookURL }}
          options: "-H \"Accept: application/vnd.github.everest-preview+json\" -H \"Authorization: token ${{ secrets.TOKEN }}\""

```
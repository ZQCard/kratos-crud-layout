# Kratos Project CRUD Template

基于[kratos-layout](https://github.com/go-kratos/kratos-layout) 生成的crud项目模板

### 中间件

服务发现 consul

缓存 redis

数据库 mysql

链路追踪 jaeger

### 使用方式

PS：前提已经安装好kratos


# 开始使用
### 创建项目
通过 kratos 命令创建项目指定模板：

```bash
kratos new administrator -r https://github.com/ZQCard/kratos-crud-layout.git
# 国内拉取失败可使用gitee源
kratos new administrator -r https://gitee.com/zq11/kratos-crud-layout.git
```

### 使用命令
```bash
# 初始化项目代码
make newServiceInit
```

### 更改配置文件
```bash
# 如新增配置项, 需更改 ./internal/conf/conf.proto 且执行make config
vim ./configs/config.yaml
```

### 更改数据表名称
```bash
vim ./internal/data/entity/xxx.go

TableName()
```

### 执行程序
```bash
kratos run
```
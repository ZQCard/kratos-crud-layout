# Kratos Project CRUD Template
在使用go-kratos开发grpc服务时候，使用官方 kratos-layout 需要更改很多东西，甚至会不生效。因为开发一套自定义的[crud模板](https://github.com/ZQCard/kratos-crud-layout )。生成以项目文件夹首字母为大写的grpc服务

基于[kratos-layout](https://github.com/go-kratos/kratos-layout) 生成的crud项目模板,根据文件夹名称生成相应的服务

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

目前数据表名称只有id字段

```bash
vim ./internal/data/entity/administrator.go TableName()
```

### 执行程序
```bash
kratos run
```

### 最后

基于以上步骤，修改相关文件(proto、service、biz、data)，基于自己的业务逻辑填充数据库字段，进行crud
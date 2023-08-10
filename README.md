# EtcdVision

## 背景

这是一个ETCD数据库可视化操作项目。

因为工作中需要经常使用`etcdctl`查看或修改etcd数据，但是命令行很不方便。也尝试过其他很多etcd可视化操作工作，但要不就是不好用/不好部署，要不就是需要收费限制。想了一下，基于众多的etcd sdk实现etcd的可视化管理操作也并不复杂，索性决定自己造一个轮子，所以有了这个项目。

## 当前状态

当前还处于初始开发阶段。

当前已有功能：

- etcd连接管理（增/删）
- 可视化查询键列表并支持模糊查询
- 可视化查询值并支持json语法高亮
- 可视化编辑并更细保存值

未来计划功能（TODO）：

- 支持etcd https连接以及添加认证
- 支持etcd连接属性编辑
- 支持键列表树形显示
- 支持添加/删除键值对
- 可视化编辑值添加其他更多格式和语法高亮（yaml/raw）
- UI界面优化

## 架构

此项目为前后端分离架构。

后端是`go`语言，使用`gin`框架提供rest接口给前端调用。

[前端](https://github.com/PangQingcheng/EtcdVision-UI)是`vue3`框架加`element plus`组件库。

后端接口列表： docs/api_TODO.md(当前还没时间写)

由于作者是后端开发，缺少设计审美和前端开发经验，所以欢迎有兴趣的朋友加入帮忙开发前端页面。亦或是可以基于此后端接口自己独立开发出前端操作页面。

## 部署运行

### 手动编译

```shell
git clone https://github.com/PangQingcheng/EtcdVision.git
go build -o etcdvision main.go
```

### make

```shell
make build
```

### docker

```shell
docker run -d -p 8080:8080 rengar/tcdvision:v1.0.0
```



## 开发




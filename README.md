# CRDProject

> make manifests: 这个命令会根据你的 Go 代码中的 Kubernetes API 定义生成对应的 CRD 清单（YAML 文件）。这些清单描述了你的 CRD 的结构，包括其 Spec 和 Status 结构，以及任何 Kubernetes RBAC 规则等.


> make generate: 这个命令会运行 Go 的代码生成工具，根据你的 API 定义生成对应的客户端代码。这包括 DeepCopy 方法（这些方法用于 Kubernetes 的 API 机制），以及可能的客户端集（clientsets）、列表器（listers）和通知器（informers）等。


```
make install
```
> make install该命令通常用于将 CRD 定义安装到你的 Kubernetes 集群中。具体来说，它会运行 Kubernetes 的自定义资源生成器，将你的 Go 类型定义（在上面的例子中是 MySQLInstanceSpec、MySQLInstanceStatus 和 MySQLInstance）转换为 Kubernetes 的自定义资源定义（CRD）的 YAML 文件，并将这个 YAML 文件应用到你的 Kubernetes 集群中，从而在你的集群中创建一个新的自定义资源类型。

```
make run
```

> make run：这个命令会启动 Operator 控制器，让它开始监听你的自定义资源的事件，并对事件进行响应。在你的日志中，你可以看到你的控制器成功启动，并开始监听 MySQLInstance 资源的事件。一旦有关于 MySQLInstance 的事件（如新建、修改或删除等），你的控制器就会被调用，然后它会执行相应的操作以使系统状态和 MySQLInstance 对象的状态保持一致。


通过kubectl get crd查看CRD的定义

```shell
kubectl get crd mysqlinstances.mysql.my.test.domain -o yaml
```

根据上面的CRD定义，spec应该包含以下字段：database，password，username和version。假设database是你想要连接的数据库名称，username和password分别是用户名和密码，而version可能表示你想要使用的MySQL版本。一个可能的mysql-instance.yaml文件可能如下：

```yaml
apiVersion: mysql.my.test.domain/v1
kind: MySQLInstance
metadata:
  name: mysql-instance
spec:
  database: mydb
  username: user
  password: pass
  version: "5.7"
```

创建资源

```shell
kubectl apply -f mysql-instance.yaml 
```

查看资源

```shell
kubectl get mysqlinstances.mysql.my.test.domain
```

更新资源(修改字段)

```shell
kubectl apply -f mysql-instance.yaml
```

删除资源

```shell
kubectl delete -f mysql-instance.yaml
```


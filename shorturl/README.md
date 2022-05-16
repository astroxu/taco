# shorturl

######
- etcd: docker 安装 [https://www.cnblogs.com/zjdxr-up/p/15409176.html](https://www.cnblogs.com/zjdxr-up/p/15409176.html)
  (brew install etcd)

```sh
#使用镜像 https://hub.docker.com/r/bitnami/etcd
docker pull bitnami/etcd:latest

#创建docker 网络
docker network create app-tier --driver bridge

#启动etcd服务器实例 –network app-tier 指定启动实例的网络配置
docker run -d --name etcd-server \
--network app-tier \
--publish 2379:2379 \
--publish 2380:2380 \
--env ALLOW_NONE_AUTHENTICATION=yes \
--env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \
bitnami/etcd:latest

# 启动etcd客户端实例，并连接上步骤服务
docker run -it --rm \
--network app-tier \
--env ALLOW_NONE_AUTHENTICATION=yes \
bitnami/etcd:latest etcdctl --endpoints http://etcd-server:2379 set /message Hello
```

- mysql: docker 安装 [https://www.jianshu.com/p/eb3d9129d880](https://www.jianshu.com/p/eb3d9129d880)

```sh
docker pull mysql/mysql-server:latest

#创建并启动MySQL服务容器
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql/mysql-server

#这时使用宿主机连接没有授权访问，需要进入mysql修改mysql访问权限。
docker exec -it mysql bash
bash-4.2# mysql -u root -p123456

#授权
mysql>CREATE USER 'root'@'%' IDENTIFIED BY 'root';
mysql>GRANT ALL ON *.* TO 'root'@'%';

#刷新权限
mysql> flush privileges;

#修改root用户密码
mysql> ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY '123456';

#刷新权限
mysql> flush privileges;
```

- redis: docker 安装 [https://blog.csdn.net/github_39581355/article/details/118938850](https://blog.csdn.net/github_39581355/article/details/118938850)

```sh
docker pull redis:latest

#创建并启动MySQL服务容器
docker run -itd --name redis-test -p 6379:6379 redis

#这时使用宿主机连接没有授权访问，需要进入mysql修改mysql访问权限。
ocker exec -it redis-test bash

#输入redis-cli
```

- 安装环境
```sh
#安装 protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.2

#安装 protoc
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/protoc-3.14.0-linux-x86_64.zip
unzip protoc-3.14.0-linux-x86_64.zip
mv bin/protoc /usr/local/bin/

#安装 goctl 工具
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/tal-tech/go-zero/tools/goctl@latest
```

- 在 `shorturl` 目录下执行 `go mod init shorturl` 初始化 `go.mod`

```sh
module shorturl

go 1.15

require (
  github.com/golang/mock v1.4.3
  github.com/golang/protobuf v1.4.2
  github.com/tal-tech/go-zero v1.1.4
  golang.org/x/net v0.0.0-20200707034311-ab3426394381
  google.golang.org/grpc v1.29.1
)
```

- 在 `shorturl/api` 目录下通过 `goctl` 生成 `api/shorturl.api`：
```sh
goctl api -o shorturl.api
```

- 使用 `goctl` 生成 `API Gateway` 代码
```sh
goctl api go -api shorturl.api -dir .
```

- 在 `rpc/transform` 目录下编写 `transform.proto` 文件 可以通过命令生成 `proto` 文件模板
```sh
goctl rpc template -o transform.proto
```

- 用 goctl 生成 rpc 代码，在 rpc/transform 目录下执行命令
```sh
goctl rpc proto -src transform.proto -dir .
```

- 查看服务是否注册
```sh
$ETCDCTL_API=3 etcdctl get transform.rpc --prefix
transform.rpc/7587851893787585061
127.0.0.1:8080
```

- 在 rpc/transform/model 目录下执行如下命令生成 CRUD+cache 代码，-c 表示使用 redis cache
```sh
goctl model mysql ddl -c -src shorturl.sql -dir .
```



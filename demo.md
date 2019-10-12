### Rabbit MQ

```
windows
安装erlang
wget http://www.rabbitmq.com/releases/erlang/erlang-19.0.4-1.el7.centos.x86_64.rpm

wget http://www.rabbitmq.com/releases/rabbitmq-server/v3.6.10/rabbitmq-server-3.6.10-1.el7.noarch.rpm

rpm -ivh erlang-19.0.4-1.el7.centos.x86_64.rpm
rpm -ivh rabbitmq-server-3.6.10-1.el7.noarch.rpm

systemctl start rabbitmq-server
systemctl stop rabbitmq-server
rabbitmq-plugins list 插件查看
rabbitmq-plugins enable rabbitmq_management 插件安装卸载
rabbitmq-plugins disable rabbitmq_management 插件安装卸载

本机端口 15672

mac 
brew install rabbitmq-server

export PATH=$PATH:/usr/local/sbin

启动RabbitMQ

前台运行rabbitmq-server
后台运行brew service start rabbitmq

Mac doc
https://www.rabbitmq.com/install-homebrew.html

默认账号密码 guest
```

### Iris 框架级速入门
```
官网： https://iris-go.com/
框架文档： https://github.com/iris-contrib/examples

```
### goland 快捷键
```
go tools gofmt: 格式优化 option+cmd+shift+f


```
















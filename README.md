# secondsKill
go-web

拷贝文件夹至目的路径
```
cp -r * path1 path2
```

### 工作模式
一个消息只能被一个消费者消费
```
先开启recieve文件 消费者
在开启publish文件 生产者
```

### Routing, 路由模式
一个消息被多个消费者获取。并且消息的目标队列可被生产者指定

### Topic, 话题模式
一个消息被多个消费者获取。消息的目标queue可用BindingKey以通配符，（#：一个或多个，*：一个词）的方式指定



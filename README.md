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

GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="/Users/pk/go:/Users/pk/pkvendor"
GOPROXY=""
GORACE=""
GOROOT="/usr/local/go"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/j4/lqrhsvfs16g4yzckl1_52sgw0000gn/T/go-build277301785=/tmp/go-build -gno-record-gcc-switches -fno-common"

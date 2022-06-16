# [Modules和Packages的区别](https://levelup.gitconnected.com/using-modules-and-packages-in-go-36a418960556)

在go语言中，package是包含一系列.go文件的目录。Module是一系列packages的集合，内置了依赖管理和版本管理。Modules在Go1.11中引入，但是知道Go1.16，go命令才默认以版本感知模式构建packages。

## Go环境变量

安装好Go环境之后，第一步需要配置Go环境变量，其中最终要的三个变量是**GOMODULE**、**GOPATH**、**GOROOT**，可以使用`go env`命令查看当前Go环境变量设置：

```shell
wangheng@wangheng-VirtualBox:~/workspace/go/studygo$ go env
GO111MODULE="on"
GOARCH="amd64"
GOBIN=""
GOCACHE="/home/wangheng/.cache/go-build"
GOENV="/home/wangheng/.config/go/env"
GOEXE=""
GOEXPERIMENT=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOINSECURE=""
GOMODCACHE="/home/wangheng/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/home/wangheng/go"
GOPRIVATE=""
GOPROXY="https://goproxy.cn,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GOVCS=""
GOVERSION="go1.17"
GCCGO="gccgo"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD="/dev/null"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build146176763=/tmp/go-build -gno-record-gcc-switches"
```

* **GOROOT**—— specifies where you go SDK is located.
* **GOPATH**—— specifies the root of your workspace(where your packages and dependencies are located)
* **GO111MODULE**—— specifies how Go import your packages. It can assume the following three values: "on", "off", "auto"

## Go1.11的特别之处

在Go1.11之前所有的Go项目必须放在环境变量**GOPATH**确定的目录下。从Go1.11开始，项目可以放在**GOPATH**确定的目录外面，并且Go1.11引入了**modules**的概念。一个Module可能包含若干个packages，并且在它的根目录下一定有一个**go.mod**文件。**go.mod**文件定义了module的导入路径，同时也定义了它的依赖路径。引入modules后：

* Go项目不必要在GOPATH目录下
* packages管理更方便

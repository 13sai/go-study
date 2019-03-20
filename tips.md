#### go get安装x/sys失败的解决

```
git clone https://github.com/golang/net.git $GOPATH/src/github.com/golang/net

git clone https://github.com/golang/sys.git $GOPATH/src/github.com/golang/sys

git clone https://github.com/golang/tools.git $GOPATH/src/github.com/golang/tools

ln -s $GOPATH/src/github.com/golang $GOPATH/src/golang.org/x
```


#### go的环境变量设置（Linux）

在/etc/profile最下面加入以下代码：

```
# 可自行改成你想的目录
export GO_INSTALL_DIR=$HOME
export GOROOT=$GO_INSTALL_DIR/go
export GOPATH=$HOME/13sai
export PATH=$GOPATH/bin:$PATH:$GO_INSTALL_DIR/go/bin
```


#### import

> import _  会执行package里的init函数，但不需要包所有函数

> import .  解决循环依赖问题


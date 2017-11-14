# cloudgo:开发web服务程序
## 概述
开发简单 web 服务程序 cloudgo，了解 web 服务器工作原理。

## 任务要求
1、熟悉go服务器工作原理，编程web服务程序，要求要有详细的注释，说明是否使用了哪些架构。

2、基于现有web库，编写一个简单web应用类似 cloudgo。

3、使用curl工具访问web程序。

4、使用ab工具对 web执行压力测试。

## 实验说明

实验使用了Martini框架。

在src/github.com路径下新建文件夹go-martini 。

cd到go-martini路径下，下载martini包。

命令：git clone https://github.com/go-martini/martini.git 。

然后进入到go-martini下执行命令：go install 。

在安装martini架构的过程中还会提示需要你安装另一个包，用上面同样方法git clone 然后go install 就可以了。

Martini是一个非常新的 Go 语言的 Web 框架，使用 Go 的 net/http 接口开发，类似 Sinatra 或者 Flask 之类的框架，你可使用自己的 DB 层、会话管理和模板。

**Martini**　的特性：

1、使用非常简单。

2、无侵入设计。

3、可与其他 Go 的包配合工作。

4、超棒的路径匹配和路由。

5、模块化设计，可轻松添加工具。

6、大量很好的处理器和中间件。

7、很棒的开箱即用特性。

8、完全兼容 http.HandlerFunc 接口。


## 实验结果测试

### 1、开启服务

使用cmd命令窗口一直进入到cloudgo目录下，执行go run main.go -p8080,使用8080端口来启动web服务，这时候打开浏览器进入到 http://localhost:8080/ 就可以显示出在代码里自定义的参数：hello world!

![1](https://github.com/imhejiamin/cloudgo/blob/master/pic/1.png)

### 2、curl工具测试

在windows下并没有curl命令工具，所以需要到官网（ https://curl.haxx.se/download.html ）下载，下载的是一个压缩包，根据自己的系统和系统位数选择下载的版本，解压之后，打开文件夹I386，
能够看到curl.exe。将这个可执行文件移动到 windows/system32 下即可。弹出对话框询问需要管理员权限，选择继续。
这样我们就可以在任何路径下使用curl工具进行测试。

直接重新打开一个cmd窗口，输入curl -v http://localhost:8080/ ，会出现以下信息，并能够成功访问到目标url。

![2](https://github.com/imhejiamin/cloudgo/blob/master/pic/2.png)

### 3、ab工具压力测试

由于windows系统下没有ab工具，所以去官网（ http://httpd.apache.org/download.cgi ）下载ab工具包，注意根据自己的系统和位数选择自己需要的工具包即可。
下载完之后解压文件到任意位置，解压后得到的文件夹名字为Apache24（不同版本可能是22 23之类的）。

然后cd到路径Apache24/bin，在这个路径下执行命令ab -n 1000 -c 100 http://localhost:8080/ 。

注意在进行ab测试前要先开启此程序的web服务（也就是保留步骤1的cmd窗口/浏览器访问http://localhost:8080/ 能够正常显示hello world!）

然后测试效果如下图，表示1000个请求全部成功。这时候开启服务的cmd窗口也会给出相应的响应。

![3](https://github.com/imhejiamin/cloudgo/blob/master/pic/3.png)

![4](https://github.com/imhejiamin/cloudgo/blob/master/pic/4.png)

![5](https://github.com/imhejiamin/cloudgo/blob/master/pic/5.png)

![6](https://github.com/imhejiamin/cloudgo/blob/master/pic/6.png)

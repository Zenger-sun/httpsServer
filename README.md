# httpsServer

一个轻量的https服务器以及client实现，包含证书生成以及本地静态文件，用于测试https的各类需求。

## Quick start

### Requirments
Go version>=1.15

### Installation
```go

git clone https://github.com/Zenger-sun/httpsServer.git

// 根据需要修改配置 
vim cert/ca.conf
vim cert/server.conf

sh cert/gen.sh // 配置了默认值，执行时只需要一路回车
```

### Build & Run

```go

go build -o server.exe
```

双击启动server.exe，如出现秒退，请检查证书是否生成  
打开浏览器访问: https://127.0.0.1:8000/, 就能看到放在web中的index.html  
或者通过接\*.html的方式: https://127.0.0.1:8000/index.html, 访问到web中的其他静态资源

附上benchmark

goos: windows  
goarch: amd64  
pkg: github/Zenger-sun/httpsServer/client  
cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz  
Benchmark_httpsServer  
Benchmark_httpsServer-8   	     350	   4942069 ns/op  
PASS  
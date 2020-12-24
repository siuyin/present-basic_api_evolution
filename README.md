# Evolving APIs with Go
While Go is the language used in demonstration code,
I have avoided using features unique to Go.

Thus these concepts should be applicable across different
computer languages as well as remote procedure call techniques
such as REST (http/json) or gRPC (http2/protocol buffers).

These presentations are written in Go [present](https://godoc.org/golang.org/x/tools/present).

Main, Basic, Level 1 content is in present-basic_api_evolution.slide.

Level 2 content: present-api_evolution-level2.slide.

```
docker run -it --name present -v godata:/home/siuyin/go -p 3999:3999 siuyin/go:dev
go get golang.org/x/tools/cmd/present
setup .bashrc to have ~/go/bin in PATH or export PATH=~/go/bin:$PATH

present -http 0.0.0.0:3999 -orighost {docker_host_IP}
eg1. present -http 0.0.0.0:3999 -orighost 192.168.99.100
eg2. present -http 0.0.0.0:3999 -orighost 127.0.0.1
```

# TCP-CDN

给你的 TCP 加上 CDN, 本项目用于 MC 服务器的代理。

## 原理

TCP 转 WebSocket 出服务端后过 CDN 到客户端后 WebSocket 转 CDN。


## 使用到的东西

- 百度云加速
- WebSockify
- Proxy

## 开始

- 服务器所在的主机安装 WebSockify。安装可以使用包管理器或者 pip。
- 运行 `websockify 0.0.0.0:80 服务器地址:端口` 会在 80 端口上开启 WebSocket。
- 百度云加速配置加速地址为你的服务器地址。
- 获取本项目 Proxy 目录下的 Go 文件直接编译，或更改IP后编译。
- 运行 `Proxy 127.0.0.1:25565 ws://你的域名` 即可把 WebScocket 转发成 TCP 并开放在本地 25565 端口下。
- Mc 写入地址 localhost 即可。

## 特别声明

- 百度云加速需要备案域名

- 如果 80 端口被占用，请配置端口转发，比如这样。

```
server {
        listen 80;
        server_name ws.jihuayu.site;

        location / {
                proxy_pass http://127.0.0.1:12345;
                proxy_http_version 1.1;
                proxy_set_header Upgrade $http_upgrade;
                proxy_set_header Connection "upgrade";
        }
}
```
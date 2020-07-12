## report-current-ip

### Why?

个人目前使用树莓派 3 刷了 openwrt 作为路由器，openwrt 启动后会有 sshd，但是每次连接都需要有 IP 地址，在公司网络中，IP 地址每次重启后就会发生变化，设置静态 IP 地址又怕重复
，从而想的是，至少每次重启机器时能将当前 IP 地址上报到自己域名，这样每次连接时直接使用域名即可，这样就会很方便。

### Usage

1. build for raspberrypi

```
make GOARCH=arm
```

2. TODO run by sysvinit

## report-current-ip

### Why?

个人目前使用树莓派 3 刷了 openwrt 作为路由器，openwrt 启动后会有 sshd，但是每次连接都需要有 IP 地址，在公司网络中，IP 地址每次重启后就会发生变化，设置静态 IP 地址又怕重复
，从而想的是，至少每次重启机器时能将当前 IP 地址上报到自己域名，这样每次连接时直接使用域名即可，这样就会很方便。

### feature
1. 支持多后端，目前只实现了aliyun dns
2. 任务保证，虽然任务每次执行一次，但会保证在上报成功后才结束任务

### Usage

1. build for raspberrypi

```
make GOARCH=arm
```

2. run by sysvinit
```shell
#!/bin/sh
#
# report-current-ip       Starts report-current-ip.
#

umask 077

start() {
	printf "Starting report-current-ip: "
	cd /opt/report-current-ip
	./report-current-ip
	echo "OK"
}
stop() {
	printf "Stopping report-current-ip: "
	killall report-current-ip
	echo "OK"
}
restart() {
	stop
	start
}

case "$1" in
  start)
	start
	;;
  stop)
	stop
	;;
  restart|reload)
	restart
	;;
  *)
	echo "Usage: $0 {start|stop|restart}"
	exit 1
esac

exit $?
```

3. run by openwrt
```
cat > /etc/init.d/report-current-ip << EOF
#!/bin/sh /etc/rc.common
#
# report-current-ip       Starts report-current-ip.
#

START=99

start() {
	printf "Starting report-current-ip: "
	cd /opt/report-current-ip
	./report-current-ip
	echo "OK"
}
stop() {
	printf "Stopping report-current-ip: "
	killall report-current-ip
	echo "OK"
}
restart() {
	stop
	start
}
EOF
/etc/init.d/report-current-ip enable
```

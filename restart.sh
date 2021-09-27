#!/bin/bash

# github: https://github.com/Alvin9999/new-pac/wiki/Goflyway%E5%85%8D%E8%B4%B9%E8%B4%A6%E5%8F%B7


proxy_ip=104.238.141.147
proxy_port=12345
proxy_password=dongtaiwang.com
local_port=8100
goflyway_dir=`pwd`

# kill 8100 goflyway
pid=`netstat -tnlp | grep ':'${local_port} | awk '{print $7}' | awk -F '/' '{print $1}'`
if [ $pid ]
then
	echo "goflyway exist pid:${pid} !!!  running  kill"
	kill ${pid}
else
	echo 'goflyway not exist!'
fi

# start goflyway
echo 'go to start goflyway!'
nohup ${goflyway_dir}'/goflyway' -up="${proxy_ip}:${proxy_port}" -k="${proxy_password}" -l=":"${local_port} > firefly.log 2>&1 &



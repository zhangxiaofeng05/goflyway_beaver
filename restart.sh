#!/bin/bash

# reference
# github: https://github.com/Alvin9999/new-pac/wiki/Goflyway%E5%85%8D%E8%B4%B9%E8%B4%A6%E5%8F%B7
# shell config: https://www.cnblogs.com/kakaisgood/p/8330576.html

local_port=8100
goflyway_dir=/root/zhangxiaofeng/hack/goflyway/
log_file=${goflyway_dir}firefly.log

source ${goflyway_dir}account.env

# kill 8100 goflyway
# lsof -i:8100 | awk '{print $2}' | grep -v PID | xargs kill 
pid=`netstat -tnlp | grep ':'${local_port} | awk '{print $7}' | awk -F '/' '{print $1}'`
if [[ $pid ]]
then
	echo `date "+%Y-%m-%d %H:%M:%S  "` "goflyway exist pid:${pid} !!!  running  kill" >> ${log_file}
	kill ${pid}
else
	echo `date "+%Y-%m-%d %H:%M:%S  "` 'goflyway not exist!' >> ${log_file}
fi

# start goflyway
echo `date "+%Y-%m-%d %H:%M:%S"` 'go to start goflyway!' >> ${log_file} 
nohup ${goflyway_dir}'/goflyway' -up="${PROXY_IP}:${PROXY_PORT}" -k="${PROXY_PASSWORD}" -l=":"${local_port} >> ${log_file} 2>&1 &



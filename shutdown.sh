#/bin/bash

local_port=8100

# kill 8100 goflyway
pid=`netstat -tnlp | grep ':'${local_port} | awk '{print $7}' | awk -F '/' '{print $1}'`
if [ $pid ]
then
	echo "goflyway exist pid:${pid} !!!  running  kill"
	kill ${pid}
else
	echo 'goflyway not exist!'
fi


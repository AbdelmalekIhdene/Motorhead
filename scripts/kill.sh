pid=$(<~/run/motorhead.pid)
(set -x; kill -15 $pid)
while ps -p $pid &> /dev/null;
do sleep 1; done

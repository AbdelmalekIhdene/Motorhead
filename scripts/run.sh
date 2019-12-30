mkdir -p ~/run
nohup ./motorhead &> /dev/null & echo $! > ~/run/motorhead.pid

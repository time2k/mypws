export GOROOT=/usr/local/go
export GOPATH=/data/gopath
export GOPROXY=https://goproxy.io
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH

cd /data/www/mypws
if [ ! -d  ./bin ]
then
    mkdir ./bin
else
	printf "\033[032m bin directory already exists \033[0m\n"
fi
if [ ! -d  ./bin_bak ]
then
    mkdir ./bin_bak
else
	printf "\033[032m bin_bak directory already exists \033[0m\n"
fi

printf "\033[36m Backup old binary program and Clean bin folder \033[0m\n"
/bin/cp -f ./bin/* ./bin_bak/
rm -rf ./bin/*

printf "\033[36m Building Binary... \033[0m\n"

DATES=`date +%Y%m%d%H%M`
NAME="mypws-http"
go build -tags "release" -o ./bin/$NAME-server ./server/http_server.go

printf "\033[36m Build Complete! \033[0m\n"
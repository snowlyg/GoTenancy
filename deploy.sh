# 使用方法:
#   ./deploy.sh production
#   ./deploy.sh dev

if [ -n "$*" ]; then
  env=$*
else
  env=dev
fi

echo "部署中 \033[1;31m$env\033[0m 从封装 \033[1;33m$(git branch | sed -n '/\* /s///p')\033[0m..."

# build seeds.go
echo "构建主体数据填充..."
GOOS=linux GOARCH=amd64 go build -o db/seeds/main db/seeds/main.go db/seeds/seeds.go

go run  main.go -compile-templates=true

echo "部署中..."
harp -s $env deploy

# 请确定你会运行 `ssh deployer@influxdb.theplant-dev.com`，或者联系 sa@theplant.jp
influxdb_table=$(git config --local remote.origin.url|sed -n 's#.*/\([^.]*\)\.git#\1#p')
user=$(git config user.name || whoami)
checksum=$(git rev-parse --short HEAD | tr -d '\n')
ssh deployer@influxdb.theplant-dev.com -- /home/deployer/deployment_record "$influxdb_table" "$user" "$env" "$checksum" || echo "发送数据到 influxdb 失败"

harp -s $env log

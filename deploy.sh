#!/bin/bash

# 设置变量
IMAGE_NAME="aidenxiezh/quotation-server"
# SERVER_IP="$1"
# SERVER_PASSWORD="$2"
VERSION="$1"

# 如果没有指定版本，则使用git最新commit id值的前6位
if [ -z "$VERSION" ]; then
  VERSION=$(git rev-parse HEAD | cut -c1-6)
fi

go build

# 构建镜像
docker build -t $IMAGE_NAME:$VERSION .
# docker push $IMAGE_NAME:$VERSION

# # 登录到远程服务器
# sshpass -p "$SERVER_PASSWORD" ssh root@$SERVER_IP <<EOF
#   # 拉取最新镜像
#   docker pull $IMAGE_NAME:$VERSION

#   # 停止旧的容器
#   docker stop quotation-ui || true

#   # 启动新的容器
#   docker run -d --name quotation-server -p 8000:8000 --rm $IMAGE_NAME:$VERSION
#   # 检查docker run是否成功
#   if [ $? -eq 0 ]
#   then
#    echo "Docker run was successful."
#   else
#    echo "Docker run failed."
#   fi

# EOF
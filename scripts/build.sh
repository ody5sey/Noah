#!/bin/bash

set -eux

# 二进制文件名
PROJECT="noah"
# 启动文件PATH
MAIN_PATH="cmd/main.go"

# build flag
LDFLAGS="-s -w"
#			   -X "main.BuildVersion=${VERSION}"
#			   -X "main.BuildDate=$(shell /bin/date "+%F %T")"

function create() {
  #判断文件夹是否存在，不存在则创建
  if [ ! -d "$1" ]; then
    mkdir -p "$1"
  fi
}

#build
function build() {
  # 创建部署文件夹
  create deployments

  #build for windows
  create deployments/windows
  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags "${LDFLAGS}" -o deployments/windows/"${PROJECT}".exe "${MAIN_PATH}"
  cd deployments
  zip -r windows.zip ./windows
  rm -r windows
  cd ../

  #build for macOS
  create "deployments/darwin"
  CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags "${LDFLAGS}" -o deployments/darwin/"${PROJECT}" "${MAIN_PATH}"
  cd deployments
  zip -r darwin.zip ./darwin
  rm -r darwin
  cd ../

  #build for linux
  create "deployments/linux"
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags "${LDFLAGS}" -o deployments/linux/"${PROJECT}" "${MAIN_PATH}"
  cd deployments
  zip -r linux.zip ./linux
  rm -r linux
  cd ../

}

build


#构建可执行文件，在deployments目录下面
.PHONY : build
build:
	@echo "build开始"
	@chmod +x ./scripts/build.sh && ./scripts/build.sh
	@echo "build结束"


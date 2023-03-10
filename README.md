# search-engine

第三届字节跳动青训营项目：搜索引擎

## 运行方法

1.设置基本依赖服务
> sudo docker-compose up

2.运行search RPC服务器服务
> cd cmd/search
> sh build.sh
> sudo sh output/bootstrap.sh

3.运行User RPC 服务器服务
> cd cmd/user
> sh build.sh
> sudo sh output/bootstrap.sh

4.运行Collection 服务器服务
> cd cmd/collection
> sh build.sh
> sudo sh output/bootstrap.sh

5.运行 API 服务器服务
> cd cmd/api
> chmod +x run.sh
> ./run.sh

# AcFun-Helper

<p align="center">
<img src="https://i.loli.net/2020/05/28/2k8dPLiGEZNHjny.png" width="120">
</p>
<h1 align="center">AcFun-DougaInfo</h1>

## 介绍
🍰AcFun-DougaInfo 用于获取AcFun站点上相关的稿件信息。

## 备注
 ```bash
 protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative .\server\protoLib\full.proto
 ```

 ```bash
 goverter gen ./server/protoLib/
 ```
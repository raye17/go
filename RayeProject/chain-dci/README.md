# chain-dci

蚂蚁链版权保护接入指南:https://antdigital.com/docs/11/305591#900136f865aiw
账号:18051299227

测试环境调用入口:EndPoint：https://openapi-sit.antchain.antgroup.com/gateway.do

线上环境调用入口:EndPoint：https://openapi.antchain.antgroup.com/gateway.do

```text
#DCI申领流程
https://cdns.fontree.cn/fonchain-main/test/image/0/chain-dci/file/24bfa6de-cf39-4bb5-83bf-6a64d2b0d534.png
```
```text
#数字版权登记流程
https://cdns.fontree.cn/fonchain-main/test/image/0/chain-dci/file/b2b9ebf8-5e8a-4697-af11-c07f788a6f77.png
```

```text
1,不需要的参数就不设置
2,文件像素要求是 下限400*400，上限5000*5000
```


### protot文件编译指令

```shell
protoc --proto_path=. --proto_path=D:/go_workspace/src  --go_out=./pb --govalidators_out=./pb --go-triple_out=./pb ./pb/dci.proto

protoc --proto_path=. --go-grpc_out=./pb --proto_path=D:/go_workspace/src  --go_out=./pb --govalidators_out=./pb --go-triple_out=./pb ./pb/dci.proto
```
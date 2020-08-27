# 客户订单管理 服务端

使用 golang gin框架

## 介绍
客户订单管理系统服务端 本质是个WEBAPI，使用注册服务的形式运行在程序之中。与mysql数据库建立沟通

## 文档
|API|用途|参数| 例子| 返回内容|
|  :----:  | ----  | ----  |----  |----  |
|getcustomerlist |获取客户表内容| needrow: 显示的总行数 <br>page:需要显示的当前页数</br> |http://127.0.0.1:8080/getcustomerlist?needrow=20&page=1 |"{\"message\":[{\"ID\":2,\"NAME\":\"test2\",\"CREATE_DATE\":\"0001-01-01T00:00:00Z\",\"CLOSE_DATE\":\"0001-01-01T00:00:00Z\",\"CLOSED\":0}]}" |
|getcustomerdetaillist|获取对应客户明细表|id: 对应客户的id号<br>needrow: 显示的总行数 </br><br>page:需要显示的当前页数</br>| http://127.0.0.1:8080/getcustomerdetaillist?id=1&needrow=20&page=1 | "{\"message\":[{\"ID\":2,\"NAME\":\"test2\",\"CREATE_DATE\":\"0001-01-01T00:00:00Z\",\"CLOSE_DATE\":\"0001-01-01T00:00:00Z\",\"CLOSED\":0}]}" |
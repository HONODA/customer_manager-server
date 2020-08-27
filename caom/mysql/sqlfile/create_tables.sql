create database CAOM;
use  CAOM;--使用CAOM数据库
--新建客户表
create table customer(id int(11),name varchar(255),created_date datetime,closed_date datetime,closed int(1));
--新建客户明细表
create table customerdetail(id int(11),name varchar(255),created_date datetime,closed_date datetime,closed int(1));

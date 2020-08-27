package mainapi

import (
	"./common"
	"./mysql"
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jander/golog/logger"
)

//分页显示客户列表
func Getcustomerlist(c *gin.Context) {
	page := c.DefaultQuery("page","1")
	needrow := c.DefaultQuery("needrow","20")
	con := new(sqls.Conninf).Conn()
	//var totalcount,pagecount
	mylist := list.New()
	rows, err := con.Query(fmt.Sprintf("CALL sp_viewPage('id,name','customer','1=1','Id desc',%s,%s,@page ,@totalcount)",page,needrow))
	if err != nil {
		errstr ,_:= fmt.Println("查询【客户列表】访问【数据库失败】：%s", err)
		c.JSON(404,gin.H{"message":errstr})
		logger.Error(fmt.Println(errstr))
		return
	}
	for rows.Next() {
		cust := new(common.Customer)
		if scanerr := rows.Scan(&cust.ID, &cust.NAME); scanerr != nil {
			logger.Error(fmt.Println("【客户列表】Scan错误：%s", err))
			return
		}
		if DATASHOWJSON {
			jsn, err := json.Marshal(cust)
			if err != nil {
				logger.Error(fmt.Println("客户列表转换JSON错误：%s", err))
				return
			}
			mylist.PushBack(string(jsn))
		} else {
			c.JSON(404, gin.H{"message": "暂时没有结构化STRUCT"})
			//mylist.PushBack(cust)
		}

	}
	str := ""
	str = convertList2Json(mylist)
	c.JSON(200, str)

}

//分页显示客户明细列表
func Getcustomerdetaillist(c *gin.Context){
	id := c.DefaultQuery("id","-1")
	if id == "-1"{
		c.JSON(404,gin.H{"message":"无数据"})
		return
	}
	mylist := list.New()
	needrow := c.DefaultQuery("needrow","1")
	page := c.DefaultQuery("page","1")
	con := new(sqls.Conninf).Conn()
	rows ,err :=con.Query(fmt.Sprintf("CALL sp_viewPage('id,name','customerdetail','id =%s','Id desc',%s,%s,@page ,@totalcount)",id,page,needrow))
	if err != nil{
		c.JSON(404,gin.H{"message":"【客户明细表】查询出错"})
		logger.Error(fmt.Println("【客户明细表】访问【数据库查询失败】：%s", err))
		return
	}
	for rows.Next(){
		custdetial := new (common.CustomerDetail)
		if scanerr := rows.Scan(&custdetial.ID,&custdetial.NAME);scanerr != nil{
			logger.Error(fmt.Println("将【客户明细表】转换成STRUCT失败:%s",scanerr))
			return
		}
		if DATASHOWJSON {
			jsn, err := json.Marshal(custdetial)
			if err != nil {
				logger.Error(fmt.Println("客户列表转换JSON错误：%s", err))
				return
			}
			mylist.PushBack(string(jsn))
		} else {
			c.JSON(404, gin.H{"message": "暂时没有结构化STRUCT"})
			//mylist.PushBack(cust)
		}
	}
	str := ""
	str = convertList2Json(mylist)
	c.JSON(200, str)
}

func InsertcustomerInf(c *gin.Context){

}
package mainapi

import (
	"./common"
	"./mysql"
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/jander/golog/logger"
)

//分页显示客户列表
func Getcustomerlist(idex int) *list.List {
	con := new(sqls.Conninf).Conn()
	//var totalcount,pagecount
	mylist := list.New()
	rows, err := con.Query("CALL sp_viewPage('id,name','customer','1=1','Id desc',1 ,20 ,@page ,@totalcount)")
	if err != nil {
		logger.Error(fmt.Println("客户列表读取错误：%s", err))
		return nil
	}
	for rows.Next() {
		cust := new(common.Customer)
		if scanerr := rows.Scan(&cust.ID, &cust.NAME); scanerr != nil {
			logger.Error(fmt.Println("客户列表Scan错误：%s", err))
			return nil
		}
		if DATASHOWJSON {
			jsn, err := json.Marshal(cust)
			if err != nil {
				logger.Error(fmt.Println("客户列表转换JSON错误：%s", err))
				return nil
			}
			mylist.PushBack(string(jsn))
		} else {
			mylist.PushBack(cust)
		}

	}
	return mylist
}

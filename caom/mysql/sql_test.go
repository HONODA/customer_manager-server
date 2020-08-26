package sqls

import "fmt"
import "testing"

func TestConn(t *testing.T) {
	fmt.Print("hello")
	con := new(Conninf)
	database := con.Conn()
	if database != nil {
		var totalcount, pagecount = 0, 0

		rows, querr := database.Queryx("CALL sp_viewPage('id,name','customer','1=1','Id desc',1 ,20 ,@page ,@totalcount)")
		var idx int
		var name = ""
		if querr != nil {
			fmt.Printf("query database failed, err:%v", querr)
			return
		} //查询数据库出错
		for rows.Next() {
			if scanerr := rows.Scan(&idx, &name); scanerr != nil {
				fmt.Printf("scan failed, err:%v", scanerr)
				return
			} //获取单条数据
			fmt.Println(idx)
			fmt.Println(name)
		}
		if rows.NextResultSet() {
			for rows.Next() {
				if scanerr := rows.Scan(&pagecount, &totalcount); scanerr != nil {
					fmt.Printf("scan failed, err:%v", scanerr)
					return
				}
			}
			fmt.Println(pagecount)
			fmt.Println(totalcount)
		}

	}

}

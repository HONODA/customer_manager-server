package sqls

import "fmt"
import "testing"

func TestConn(t *testing.T) {
	fmt.Print("hello")
	con := new(Conninf)
	database := con.Conn()
	if database != nil {
		var totalcount, pagecount = 0, 0

		row := database.QueryRow("CALL sp_viewPage('id,name','customer','1=1','Id desc',1 ,20 ,? ,?)", totalcount, pagecount)
		var idx int
		var name = ""
		if err := row.Scan(&idx, &name); err != nil {
			fmt.Printf("scan failed, err:%v", err)
			return
		}
		fmt.Println(idx)
		fmt.Println(name)
	}

}

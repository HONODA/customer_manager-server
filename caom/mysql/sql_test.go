package sqls

import "fmt"
import "testing"

func TestConn(t *testing.T) {
	fmt.Print("hello")
	con := new(conninf)
	database := con.Conn()
	if database != nil {
		row := database.QueryRow("select id, name from customer")
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

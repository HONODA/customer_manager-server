package sqls

import _ "github.com/go-sql-driver/mysql"
import "github.com/jmoiron/sqlx"
import "github.com/jander/golog/logger"
import "fmt"

type Conninf struct {
	database string
	host     string
	port     string
	user     string
	pasw     string
}

func (inf *Conninf) Inf() {
	inf.database = "CAOM"
	inf.host = "localhost"
	inf.pasw = ""
	inf.user = "root"
	inf.port = "3306"

}

func (i *Conninf) Conn() *sqlx.DB {
	i.Inf()
	database, err := sqlx.Open("mysql", i.user+":"+i.pasw+"@tcp("+i.host+":"+i.port+")/"+i.database)
	if err != nil {
		logger.Fatal(err)
		fmt.Println("HN:ERROR LOADING DATABASE")
		return nil
	}
	return database
}

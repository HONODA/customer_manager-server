package mainapi

import (
	"container/list"
	"github.com/gin-gonic/gin"
)

const (
	DATASHOWJSON = true
)

func convertList2Json(l *list.List) string {
	var str = "{\"message\":["
	var count = l.Len() - 1
	for i := l.Front(); i != nil; i = i.Next() {
		if count > 0 {
			str += i.Value.(string) + "," //强制转string
		} else if count == 0 {
			str += i.Value.(string) //强制转string
		}
		count--
	}
	str += "]}"
	return str
}
func RunServerApp() {
	r := gin.Default()
	r.POST("/getcustomerlist", Getcustomerlist)
	r.POST("/getcustomerdetaillist",Getcustomerdetaillist)
	r.Run() // listen and serve on 0.0.0.0:8080
}

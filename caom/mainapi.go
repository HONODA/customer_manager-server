package mainapi

import (
	"container/list"
	"github.com/gin-gonic/gin"
)

const (
	DATASHOWJSON = true
)

func convertList2Json(l *list.List) string {
	var str = "["
	var count = l.Len() - 1
	for i := l.Front(); i != nil; i = i.Next() {
		if count > 0 {
			str += i.Value.(string) + "," //强制转string
		} else if count == 0 {
			str += i.Value.(string) //强制转string
		}
	}
	str += "]"
	return str
}
func RunServerApp() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		list := Getcustomerlist(1)
		str := ""
		if DATASHOWJSON {
			str = convertList2Json(list)
		} else {
			c.JSON(500, gin.H{"message": "暂时没有结构化STRUCT"})
			return
		}
		c.JSON(200, str)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

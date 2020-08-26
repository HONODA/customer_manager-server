package common

import(
	"github.com/jander/golog/logger"
	)

func Loggerinit(){
	file := "error.log"
	rotatinghander := logger.NewRotatingHandler("./",file,4,4*2048*2048)
	logger.SetHandlers(rotatinghander)
}

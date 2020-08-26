package main

import (
	"./caom"
	"./caom/common"
	"fmt"
	"github.com/jander/golog/logger"
	"github.com/kardianos/service"
	"os"
)

type program struct{}

func (p *program) Start(s service.Service) error {

	go p.run()

	return nil

}

func (p *program) run() {
	fmt.Println("运行服务")

	mainapi.RunServerApp()

}

func (p *program) Stop(s service.Service) error {

	return nil

}

func main() {
	common.Loggerinit()
	svcConfig := &service.Config{

		Name: "CAOMSERVER", //服务显示名称

		DisplayName: "CAOM", //服务名称

		Description: "客户订单系统CAOM后台服务", //服务描述

	}

	prg := &program{}

	s, err := service.New(prg, svcConfig)

	if err != nil {
		logger.Fatal(err)
	}

	if err != nil {

		logger.Fatal(err)

	}

	if len(os.Args) > 1 {

		if os.Args[1] == "install" {

			err = s.Install()
			if err != nil {
				logger.Println(err)
			} else {
				logger.Println("服务安装成功")
			}
			return

		}

		if os.Args[1] == "remove" {
			err = s.Uninstall()
			if err != nil {
				logger.Println(err)
			} else {
				logger.Println("服务卸载成功")
			}
			return

		}

	}

	err = s.Run()

	if err != nil {

		logger.Error(err)

	}

}

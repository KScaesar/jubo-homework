package main

import (
	"log"

	"github.com/KScaesar/jubo-homework/backend/configs"
	"github.com/KScaesar/jubo-homework/backend/ioc"
)

func main() {
	log.Default().Println("server load config !")

	cfg, err := configs.NewProjectConfig()
	if err != nil {
		log.Default().Panic(err)
	}

	server, err := ioc.NewHttpServerV1(cfg)
	if err != nil {
		log.Default().Panic(err)
	}

	log.Default().Println("server run !")
	err = server.Run(":" + cfg.ServerPort)
	if err != nil {
		log.Default().Panic(err)
	}
}

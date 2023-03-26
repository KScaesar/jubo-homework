package main

import (
	"fmt"

	"github.com/KScaesar/jubo-homework/backend/configs"
)

func main() {
	cfg, err := configs.NewProjectConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(*cfg, cfg.Pgsql)
}

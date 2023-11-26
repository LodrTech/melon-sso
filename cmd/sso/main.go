package main

import (
	"fmt"

	"github.com/LodrTech/melon-sso/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
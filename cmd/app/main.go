package main

import (
	"GolangQuizlet/internal/config"
	"fmt"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
	//init logger: slog
	//init storage: postgres
	//init router : chi, chi render
	//run server :
}

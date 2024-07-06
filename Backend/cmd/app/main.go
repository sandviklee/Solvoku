package main

import (
	"fmt"
	"github.com/sandviklee/solvoku/internal/app"
	"time"
)

func main() {
	start := time.Now()
	app.Run()
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("%fs was the time it took to solve this Sudoku.", elapsed.Seconds())
}

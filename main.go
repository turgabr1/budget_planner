package main

import (
	"budget_planner/cmd"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//repository.OpenDatabase()
	cmd.Execute()
}

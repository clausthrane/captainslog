package main

import (
	"log"
	"os"
	"github.com/clausthrane/captainslog/view/cli"
	"github.com/clausthrane/captainslog/repository/tasks"
	"github.com/clausthrane/captainslog/services"
)

var logger = log.New(os.Stdout, " ", log.Ldate | log.Ltime | log.Lshortfile)

func main() {
	dao := task_repository.NewFileBackedTaskRepository("/tmp/tasks")
	service := services.NewTaskService(dao, "nisse")
	cli.NewCli(service)
}

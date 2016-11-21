package captainslog

import (
	"github.com/urfave/cli"
	"github.com/clausthrane/captainslog/captainslog/repository/tasks"
	"github.com/clausthrane/captainslog/captainslog/view/api"
	"github.com/clausthrane/captainslog/captainslog/services"
	"context"
	"github.com/clausthrane/captainslog/captainslog/entities"
	"os"
)

type CaptainsLog struct {
	Commands *api.Commands
	Queries  *api.Queries
}

func NewCaptainsLog(c *cli.Context, repositoryRoot string, projectID entities.ProjectID) *CaptainsLog {

	dao := task_repository.NewFileBackedTaskRepository(repositoryRoot)
	service := services.NewTaskService(dao, projectID)
	ctx := context.WithValue(context.Background(), "user", entities.UserID(os.Getenv("USER")))

	return &CaptainsLog{
		api.NewCommands(ctx, service),
		api.NewQueries(ctx, service),
	}
}

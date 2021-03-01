package remembrancer

import (
	"context"
	"fmt"

	"google.golang.org/api/tasks/v1"
)

type Remembrancer struct {
	tasksSrv tasks.Service
}

func New() (*Remembrancer, error) {
	ctx := context.Background()
	tasksSrv, err := tasks.NewService(ctx)
	if err != nil {
		return nil, err // TODO: Wrap error
	}

	return &Remembrancer{}, nil
}

func (*Remembrancer) GetTasks() {

	tasksList, err := tasksSrv.Tasklists.List().MaxResults(10).Do()
	if err != nil {
		panic(err)
	}

	for _, task := range tasksList.Items {
		fmt.Println(task.Title)
	}
}

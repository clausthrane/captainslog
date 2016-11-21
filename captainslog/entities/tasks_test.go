package entities

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"context"
)


func TestTasksCanBeInstatiated(t *testing.T) {
	assert := assert.New(t)

	ctx := context.WithValue(context.Background(), "user", UserID("foobar"))
	task := NewTask(ctx, "hello")
	assert.NotNil(task, "expecting err")
	assert.Equal("hello", task.Description)
}

func TestTasksCanBeMarshalled(t *testing.T) {
	assert := assert.New(t)

	ctx := context.WithValue(context.Background(), "user", UserID("foobar"))
	task := NewTask(ctx, "hello")
	assert.NotNil(task, "expecting err")
	assert.Equal("{\"task_id\":\"foobar\",\"user\":\"foobar\",\"des", string(task.marshal()[:40]))
}

func TestTasksGroupsBeMarshalled(t *testing.T) {
	assert := assert.New(t)

	ctx := context.WithValue(context.Background(), "user", UserID("foobar"))

	task1 := NewTask(ctx, "foo")
	task2 := NewTask(ctx, "bar")

	g := NewTaskGroup()
	g.Add(TaskGroupID("nisse"), task1)
	g.Add(TaskGroupID("nisse2"), task2)
	assert.NotNil(g.Marshal())
}
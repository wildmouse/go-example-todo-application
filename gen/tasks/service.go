// Code generated by goa v3.0.6, DO NOT EDIT.
//
// tasks service
//
// Command:
// $ goa gen github.com/wild-mouse/go-example-todo-application/design

package tasks

import (
	"context"
)

// The task service.
type Service interface {
	// GetTask implements get_task.
	GetTask(context.Context, *GetTaskPayload) (res *Task, err error)
	// GetTasks implements get_tasks.
	GetTasks(context.Context) (res []*Task, err error)
	// AddTask implements add_task.
	AddTask(context.Context, *Task) (res *Task, err error)
	// UpdateTask implements update_task.
	UpdateTask(context.Context, *Task) (res *Task, err error)
	// DeleteTask implements delete_task.
	DeleteTask(context.Context, *DeleteTaskPayload) (res *Task, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "tasks"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"get_task", "get_tasks", "add_task", "update_task", "delete_task"}

// GetTaskPayload is the payload type of the tasks service get_task method.
type GetTaskPayload struct {
	// ID of task
	ID uint32
}

// Task is the result type of the tasks service get_task method.
type Task struct {
	// ID is the unique id of the task.
	ID *uint32
	// Name of task
	Name string
}

// DeleteTaskPayload is the payload type of the tasks service delete_task
// method.
type DeleteTaskPayload struct {
	ID string
}

// Code generated by goa v3.0.6, DO NOT EDIT.
//
// tasks HTTP client types
//
// Command:
// $ goa gen github.com/wild-mouse/go-example-todo-application/design

package client

import (
	"unicode/utf8"

	tasks "github.com/wild-mouse/go-example-todo-application/gen/tasks"
	goa "goa.design/goa/v3/pkg"
)

// AddTaskRequestBody is the type of the "tasks" service "add_task" endpoint
// HTTP request body.
type AddTaskRequestBody struct {
	// ID is the unique id of the task.
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of task
	Name string `form:"name" json:"name" xml:"name"`
}

// UpdateTaskRequestBody is the type of the "tasks" service "update_task"
// endpoint HTTP request body.
type UpdateTaskRequestBody struct {
	// Name of task
	Name string `form:"name" json:"name" xml:"name"`
}

// GetTaskResponseBody is the type of the "tasks" service "get_task" endpoint
// HTTP response body.
type GetTaskResponseBody struct {
	// ID is the unique id of the task.
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of task
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// GetTasksResponseBody is the type of the "tasks" service "get_tasks" endpoint
// HTTP response body.
type GetTasksResponseBody []*TaskResponse

// AddTaskResponseBody is the type of the "tasks" service "add_task" endpoint
// HTTP response body.
type AddTaskResponseBody struct {
	// ID is the unique id of the task.
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of task
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// UpdateTaskResponseBody is the type of the "tasks" service "update_task"
// endpoint HTTP response body.
type UpdateTaskResponseBody struct {
	// ID is the unique id of the task.
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of task
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// DeleteTaskResponseBody is the type of the "tasks" service "delete_task"
// endpoint HTTP response body.
type DeleteTaskResponseBody struct {
	// ID is the unique id of the task.
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of task
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// TaskResponse is used to define fields on response body types.
type TaskResponse struct {
	// ID is the unique id of the task.
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of task
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// NewAddTaskRequestBody builds the HTTP request body from the payload of the
// "add_task" endpoint of the "tasks" service.
func NewAddTaskRequestBody(p *tasks.Task) *AddTaskRequestBody {
	body := &AddTaskRequestBody{
		ID:   p.ID,
		Name: p.Name,
	}
	return body
}

// NewUpdateTaskRequestBody builds the HTTP request body from the payload of
// the "update_task" endpoint of the "tasks" service.
func NewUpdateTaskRequestBody(p *tasks.Task) *UpdateTaskRequestBody {
	body := &UpdateTaskRequestBody{
		Name: p.Name,
	}
	return body
}

// NewGetTaskTaskOK builds a "tasks" service "get_task" endpoint result from a
// HTTP "OK" response.
func NewGetTaskTaskOK(body *GetTaskResponseBody) *tasks.Task {
	v := &tasks.Task{
		ID:   body.ID,
		Name: *body.Name,
	}
	return v
}

// NewGetTasksTaskOK builds a "tasks" service "get_tasks" endpoint result from
// a HTTP "OK" response.
func NewGetTasksTaskOK(body []*TaskResponse) []*tasks.Task {
	v := make([]*tasks.Task, len(body))
	for i, val := range body {
		v[i] = &tasks.Task{
			ID:   val.ID,
			Name: *val.Name,
		}
	}
	return v
}

// NewAddTaskTaskCreated builds a "tasks" service "add_task" endpoint result
// from a HTTP "Created" response.
func NewAddTaskTaskCreated(body *AddTaskResponseBody) *tasks.Task {
	v := &tasks.Task{
		ID:   body.ID,
		Name: *body.Name,
	}
	return v
}

// NewUpdateTaskTaskOK builds a "tasks" service "update_task" endpoint result
// from a HTTP "OK" response.
func NewUpdateTaskTaskOK(body *UpdateTaskResponseBody) *tasks.Task {
	v := &tasks.Task{
		ID:   body.ID,
		Name: *body.Name,
	}
	return v
}

// NewDeleteTaskTaskOK builds a "tasks" service "delete_task" endpoint result
// from a HTTP "OK" response.
func NewDeleteTaskTaskOK(body *DeleteTaskResponseBody) *tasks.Task {
	v := &tasks.Task{
		ID:   body.ID,
		Name: *body.Name,
	}
	return v
}

// ValidateGetTaskResponseBody runs the validations defined on
// get_task_response_body
func ValidateGetTaskResponseBody(body *GetTaskResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	return
}

// ValidateAddTaskResponseBody runs the validations defined on
// add_task_response_body
func ValidateAddTaskResponseBody(body *AddTaskResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	return
}

// ValidateUpdateTaskResponseBody runs the validations defined on
// update_task_response_body
func ValidateUpdateTaskResponseBody(body *UpdateTaskResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	return
}

// ValidateDeleteTaskResponseBody runs the validations defined on
// delete_task_response_body
func ValidateDeleteTaskResponseBody(body *DeleteTaskResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	return
}

// ValidateTaskResponse runs the validations defined on taskResponse
func ValidateTaskResponse(body *TaskResponse) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	return
}

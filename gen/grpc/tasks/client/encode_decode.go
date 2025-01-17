// Code generated by goa v3.0.6, DO NOT EDIT.
//
// tasks gRPC client encoders and decoders
//
// Command:
// $ goa gen github.com/wild-mouse/go-example-todo-application/design

package client

import (
	"context"

	taskspb "github.com/wild-mouse/go-example-todo-application/gen/grpc/tasks/pb"
	tasks "github.com/wild-mouse/go-example-todo-application/gen/tasks"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildGetTaskFunc builds the remote method to invoke for "tasks" service
// "get_task" endpoint.
func BuildGetTaskFunc(grpccli taskspb.TasksClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.GetTask(ctx, reqpb.(*taskspb.GetTaskRequest), opts...)
		}
		return grpccli.GetTask(ctx, &taskspb.GetTaskRequest{}, opts...)
	}
}

// EncodeGetTaskRequest encodes requests sent to tasks get_task endpoint.
func EncodeGetTaskRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*tasks.GetTaskPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "get_task", "*tasks.GetTaskPayload", v)
	}
	return NewGetTaskRequest(payload), nil
}

// DecodeGetTaskResponse decodes responses from the tasks get_task endpoint.
func DecodeGetTaskResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*taskspb.GetTaskResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "get_task", "*taskspb.GetTaskResponse", v)
	}
	if err := ValidateGetTaskResponse(message); err != nil {
		return nil, err
	}
	res := NewGetTaskResult(message)
	return res, nil
}

// BuildGetTasksFunc builds the remote method to invoke for "tasks" service
// "get_tasks" endpoint.
func BuildGetTasksFunc(grpccli taskspb.TasksClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.GetTasks(ctx, reqpb.(*taskspb.GetTasksRequest), opts...)
		}
		return grpccli.GetTasks(ctx, &taskspb.GetTasksRequest{}, opts...)
	}
}

// DecodeGetTasksResponse decodes responses from the tasks get_tasks endpoint.
func DecodeGetTasksResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*taskspb.GetTasksResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "get_tasks", "*taskspb.GetTasksResponse", v)
	}
	if err := ValidateGetTasksResponse(message); err != nil {
		return nil, err
	}
	res := NewGetTasksResult(message)
	return res, nil
}

// BuildAddTaskFunc builds the remote method to invoke for "tasks" service
// "add_task" endpoint.
func BuildAddTaskFunc(grpccli taskspb.TasksClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.AddTask(ctx, reqpb.(*taskspb.AddTaskRequest), opts...)
		}
		return grpccli.AddTask(ctx, &taskspb.AddTaskRequest{}, opts...)
	}
}

// EncodeAddTaskRequest encodes requests sent to tasks add_task endpoint.
func EncodeAddTaskRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*tasks.Task)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "add_task", "*tasks.Task", v)
	}
	return NewAddTaskRequest(payload), nil
}

// DecodeAddTaskResponse decodes responses from the tasks add_task endpoint.
func DecodeAddTaskResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*taskspb.AddTaskResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "add_task", "*taskspb.AddTaskResponse", v)
	}
	if err := ValidateAddTaskResponse(message); err != nil {
		return nil, err
	}
	res := NewAddTaskResult(message)
	return res, nil
}

// BuildUpdateTaskFunc builds the remote method to invoke for "tasks" service
// "update_task" endpoint.
func BuildUpdateTaskFunc(grpccli taskspb.TasksClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.UpdateTask(ctx, reqpb.(*taskspb.UpdateTaskRequest), opts...)
		}
		return grpccli.UpdateTask(ctx, &taskspb.UpdateTaskRequest{}, opts...)
	}
}

// EncodeUpdateTaskRequest encodes requests sent to tasks update_task endpoint.
func EncodeUpdateTaskRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*tasks.Task)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "update_task", "*tasks.Task", v)
	}
	return NewUpdateTaskRequest(payload), nil
}

// DecodeUpdateTaskResponse decodes responses from the tasks update_task
// endpoint.
func DecodeUpdateTaskResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*taskspb.UpdateTaskResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "update_task", "*taskspb.UpdateTaskResponse", v)
	}
	if err := ValidateUpdateTaskResponse(message); err != nil {
		return nil, err
	}
	res := NewUpdateTaskResult(message)
	return res, nil
}

// BuildDeleteTaskFunc builds the remote method to invoke for "tasks" service
// "delete_task" endpoint.
func BuildDeleteTaskFunc(grpccli taskspb.TasksClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.DeleteTask(ctx, reqpb.(*taskspb.DeleteTaskRequest), opts...)
		}
		return grpccli.DeleteTask(ctx, &taskspb.DeleteTaskRequest{}, opts...)
	}
}

// EncodeDeleteTaskRequest encodes requests sent to tasks delete_task endpoint.
func EncodeDeleteTaskRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*tasks.DeleteTaskPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "delete_task", "*tasks.DeleteTaskPayload", v)
	}
	return NewDeleteTaskRequest(payload), nil
}

// DecodeDeleteTaskResponse decodes responses from the tasks delete_task
// endpoint.
func DecodeDeleteTaskResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*taskspb.DeleteTaskResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "delete_task", "*taskspb.DeleteTaskResponse", v)
	}
	if err := ValidateDeleteTaskResponse(message); err != nil {
		return nil, err
	}
	res := NewDeleteTaskResult(message)
	return res, nil
}

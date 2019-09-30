// Code generated by goa v3.0.6, DO NOT EDIT.
//
// tasks gRPC client
//
// Command:
// $ goa gen github.com/wild-mouse/go-example-todo-application/design

package client

import (
	"context"

	taskspb "github.com/wild-mouse/go-example-todo-application/gen/grpc/tasks/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli taskspb.TasksClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the tasks service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: taskspb.NewTasksClient(cc),
		opts:    opts,
	}
}

// CountTasks calls the "CountTasks" function in taskspb.TasksClient interface.
func (c *Client) CountTasks() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildCountTasksFunc(c.grpccli, c.opts...),
			nil,
			DecodeCountTasksResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

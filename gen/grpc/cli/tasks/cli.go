// Code generated by goa v3.0.6, DO NOT EDIT.
//
// tasks gRPC client CLI support package
//
// Command:
// $ goa gen github.com/wild-mouse/go-example-todo-application/design

package cli

import (
	"flag"
	"fmt"
	"os"

	tasksc "github.com/wild-mouse/go-example-todo-application/gen/grpc/tasks/client"
	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `tasks count-tasks
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` tasks count-tasks` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		tasksFlags = flag.NewFlagSet("tasks", flag.ContinueOnError)

		tasksCountTasksFlags = flag.NewFlagSet("count-tasks", flag.ExitOnError)
	)
	tasksFlags.Usage = tasksUsage
	tasksCountTasksFlags.Usage = tasksCountTasksUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "tasks":
			svcf = tasksFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "tasks":
			switch epn {
			case "count-tasks":
				epf = tasksCountTasksFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "tasks":
			c := tasksc.NewClient(cc, opts...)
			switch epn {
			case "count-tasks":
				endpoint = c.CountTasks()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// tasksUsage displays the usage of the tasks command and its subcommands.
func tasksUsage() {
	fmt.Fprintf(os.Stderr, `The task service.
Usage:
    %s [globalflags] tasks COMMAND [flags]

COMMAND:
    count-tasks: CountTasks implements count_tasks.

Additional help:
    %s tasks COMMAND --help
`, os.Args[0], os.Args[0])
}
func tasksCountTasksUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tasks count-tasks

CountTasks implements count_tasks.

Example:
    `+os.Args[0]+` tasks count-tasks
`, os.Args[0])
}

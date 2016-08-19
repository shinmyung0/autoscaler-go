package cli

import (
	"fmt"

	"github.com/shinmyung0/clite"
)

var program *clite.Program

func placeholder(args []string) int {
	fmt.Println("noop function called with", args)
	return 0
}

// This function is executed on file load
func init() {

	program = clite.NewProgram("autoscaler").
		HasCommand("server", "Runs the autoscaler server", RunServer).
		HasCommand("services", "Displays the currently detectable services", placeholder).
		HasCommand("instances", "Displays detectable instances for a service", placeholder).
		HasCommand("policies", "Displays all currently set policies", placeholder).
		HasCommand("status", "Show evaluation status for a service", placeholder).
		HasCommand("get-policy", "Get details on the policies for a service", placeholder).
		HasCommand("set-policy", "Set policies from file", placeholder).
		HasCommand("remove-policy", "Remove policy for a service", placeholder).
		HasCommand("remove-all-policies", "Remove all policies", placeholder)

	// program.Command("server").
	// 	HasOptionalArg("-v", "Verbose mode. The server will output debug logs.")
	program.Command("instances").
		HasRequiredArg("service name", "Name of the service to find instances for")

	program.Command("status").
		HasRequiredArg("service name", "Name of service to get evaluation status for")

	program.Command("get-policy").
		HasRequiredArg("service name", "Name of service to get details for")

	program.Command("set-policy").
		HasRequiredArg("file name", "Path of JSON file containing policy definitions")

	program.Command("remove-policy").
		HasRequiredArg("service name", "Name of service to remove policy")
}

func Run(args []string) int {
	return program.Run(args)
}

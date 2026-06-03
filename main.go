package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: tswitch <task-id>\n")
		os.Exit(1)
	}

	newTaskID, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid task ID: %v\n", err)
		os.Exit(1)
	}

	if err := switchTask(newTaskID); err != nil {
		fmt.Fprintf(os.Stderr, "switch task: %v\n", err)
		os.Exit(1)
	}
}

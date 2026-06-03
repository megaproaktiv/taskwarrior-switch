package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// switchTask stops the currently active task and starts a new one.
func switchTask(newTaskID int) error {
	activeTaskID, err := findActiveTask()
	if err != nil {
		return fmt.Errorf("find active task: %w", err)
	}

	if activeTaskID > 0 {
		if err := stopTask(activeTaskID); err != nil {
			return fmt.Errorf("stop task %d: %w", activeTaskID, err)
		}
		fmt.Printf("Stopped task %d\n", activeTaskID)
	}

	if err := startTask(newTaskID); err != nil {
		return fmt.Errorf("start task %d: %w", newTaskID, err)
	}
	fmt.Printf("Started task %d\n", newTaskID)

	return nil
}

// findActiveTask parses todo output to find the currently active task.
// Returns 0 if no active task is found.
func findActiveTask() (int, error) {
	cmd := exec.Command("todo")
	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("execute todo: %w", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	
	for scanner.Scan() {
		line := scanner.Text()
		
		// Skip header and empty lines
		if strings.HasPrefix(line, "ID ") || strings.TrimSpace(line) == "" {
			continue
		}

		// The format is: ID [Active] Age ...
		// Active column is position-based and may be empty
		// We need to check if there are two consecutive time values after the ID
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}
		
		// fields[0] = ID
		// fields[1] = could be Active time OR Age (if Active is empty)
		// fields[2] = Age or next field
		
		// Check if field[1] and field[2] are both time values
		// This means field[1] is Active time and field[2] is Age
		field1IsTime := strings.HasSuffix(fields[1], "min") || 
		                strings.HasSuffix(fields[1], "h") || 
		                strings.HasSuffix(fields[1], "d")
		field2IsTime := strings.HasSuffix(fields[2], "min") || 
		                strings.HasSuffix(fields[2], "h") || 
		                strings.HasSuffix(fields[2], "d")
		
		if field1IsTime && field2IsTime {
			// Both are times, so field[1] is Active time
			taskID, err := strconv.Atoi(fields[0])
			if err != nil {
				continue
			}
			return taskID, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scan output: %w", err)
	}

	return 0, nil
}

// stopTask stops the specified task.
func stopTask(taskID int) error {
	cmd := exec.Command("todo", strconv.Itoa(taskID), "stop")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("execute: %w", err)
	}
	return nil
}

// startTask starts the specified task.
func startTask(taskID int) error {
	cmd := exec.Command("todo", "start", strconv.Itoa(taskID))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("execute: %w", err)
	}
	return nil
}

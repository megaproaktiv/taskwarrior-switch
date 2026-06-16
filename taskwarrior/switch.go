package taskwarrior

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

const activeMarker = "*"

// switchTask stops the currently active task and starts a new one.
func SwitchTask(newTaskID int) error {
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

// findActiveTask runs "todo ls" and finds the task marked with "*" in the A column.
// Returns 0 if no active task is found.
func findActiveTask() (int, error) {
	cmd := exec.Command("todo", "ls")
	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("execute todo ls: %w", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()

		// Skip header and empty lines
		if strings.HasPrefix(line, "ID ") || strings.TrimSpace(line) == "" {
			continue
		}

		// Format: ID A Project Description
		// fields[0] = ID, fields[1] = "*" if active, else Project
		fields := strings.Fields(line)
		if len(fields) >= 2 && fields[1] == activeMarker {
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

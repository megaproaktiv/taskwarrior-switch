package taskwarrior_test

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseActiveTask(t *testing.T) {
	tests := []struct {
		name       string
		todoOutput string
		wantID     int
	}{
		{
			name: "single digit active task",
			todoOutput: `ID A Project   Description
10   otto      step by step
 5 * education Dashboard
 7   presales  Davero`,
			wantID: 5,
		},
		{
			name: "two digit active task",
			todoOutput: `ID A Project   Description
10 * otto      step by step
 2   otto      Aufgabe Gernot
 5   education Dashboard`,
			wantID: 10,
		},
		{
			name: "no active task",
			todoOutput: `ID A Project   Description
 1   otto      otto Elavate discussion
 2   otto      Aufgabe Gernot`,
			wantID: 0,
		},
		{
			name: "active task at end of list",
			todoOutput: `ID A Project   Description
 1   otto      otto Elavate discussion
 2   otto      Aufgabe Gernot
 9 * conrad    Email Architecture`,
			wantID: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID := parseActiveTaskFromOutput(tt.todoOutput)
			if gotID != tt.wantID {
				t.Errorf("got %d, want %d", gotID, tt.wantID)
			}
		})
	}
}

// parseActiveTaskFromOutput mirrors the parsing logic in findActiveTask for testing.
func parseActiveTaskFromOutput(output string) int {
	for line := range strings.SplitSeq(output, "\n") {
		if strings.HasPrefix(line, "ID ") || strings.TrimSpace(line) == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) >= 2 && fields[1] == "*" {
			var taskID int
			if n, _ := fmt.Sscanf(fields[0], "%d", &taskID); n == 1 {
				return taskID
			}
		}
	}
	return 0
}

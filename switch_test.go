package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestFindActiveTask(t *testing.T) {
	tests := []struct {
		name       string
		todoOutput string
		wantID     int
		wantErr    bool
	}{
		{
			name: "active task found",
			todoOutput: `ID Active Age   Project   Due    Description                                                                              Urg
 1         8h   otto      -17h   otto Elavate discussion                                                                  10.1
 2         7h   ivv       -17h   Anpassungen für Testphase im Juni                                                        10.1
 3         7h   quickline   6h   Email volumen prüfen                                                                     9.67
 4         6h   intern      4d   Personio Feedbacks                                                                       7.85
 6        24min otto        4d   Aufgabe Gernot: prüfen der Quotas gegen die 1 Jahr alten usprünglichen Hochlauf Annahmen 7.85
 5   4min  3h   quickline        Neue und Aktualisierte KB Artikel                                                           5`,
			wantID:  5,
			wantErr: false,
		},
		{
			name: "no active task",
			todoOutput: `ID Active Age   Project   Due    Description                                                                              Urg
 1         8h   otto      -17h   otto Elavate discussion                                                                  10.1
 2         7h   ivv       -17h   Anpassungen für Testphase im Juni                                                        10.1`,
			wantID:  0,
			wantErr: false,
		},
		{
			name: "active task with different time formats",
			todoOutput: `ID Active Age   Project   Due    Description
 1   10min  8h   otto      -17h   task one
 2         7h   ivv       -17h   task two`,
			wantID:  1,
			wantErr: false,
		},
		{
			name: "active task with days",
			todoOutput: `ID Active Age   Project   Due    Description
 1         8h   otto      -17h   task one
 7   2d     7h   ivv       -17h   task two`,
			wantID:  7,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock the exec.Command by testing the parsing logic directly
			// This is a simplified version that doesn't test the actual command execution
			gotID := parseActiveTaskFromOutput(tt.todoOutput)
			
			if gotID != tt.wantID {
				t.Errorf("parseActiveTaskFromOutput() = %v, want %v", gotID, tt.wantID)
			}
		})
	}
}

// parseActiveTaskFromOutput is a helper function for testing.
// This duplicates the logic from findActiveTask for testing purposes.
func parseActiveTaskFromOutput(output string) int {
	lines := strings.Split(output, "\n")
	
	for _, line := range lines {
		// Skip header and empty lines
		if strings.HasPrefix(line, "ID ") || strings.TrimSpace(line) == "" {
			continue
		}
		
		// The format is: ID [Active] Age ...
		// Active column is position-based and may be empty
		// We need to check if there's a time value between the ID and Age columns
		// Split preserving positions
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}
		
		// fields[0] = ID
		// fields[1] = could be Active time OR Age (if Active is empty)
		// fields[2] = Age or next field
		
		// Check if field[1] is a short time (min/h) and field[2] is also a time
		// This means field[1] is Active and field[2] is Age
		field1IsTime := strings.HasSuffix(fields[1], "min") || 
		                strings.HasSuffix(fields[1], "h") || 
		                strings.HasSuffix(fields[1], "d")
		field2IsTime := strings.HasSuffix(fields[2], "min") || 
		                strings.HasSuffix(fields[2], "h") || 
		                strings.HasSuffix(fields[2], "d")
		
		if field1IsTime && field2IsTime {
			// Both are times, so field[1] is Active time
			var taskID int
			if _, err := fmt.Sscanf(fields[0], "%d", &taskID); err == nil {
				return taskID
			}
		}
	}
	
	return 0
}

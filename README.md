# taskwarrior-switch

A simple Go tool to switch between active tasks in taskwarrior/todo.

## Installation

### Using Task (Recommended)

```bash
task install
```

This will build and install `tswitch` to `/usr/local/bin/`.

### Manual Build

```bash
go build -o tswitch
# Optionally move to your PATH
sudo mv tswitch /usr/local/bin/
```

### Available Tasks

- `task build` - Build the binary with optimizations
- `task install` - Build and install to system PATH
- `task test` - Run tests
- `task test:coverage` - Run tests with coverage report
- `task clean` - Remove build artifacts
- `task fmt` - Format Go code
- `task lint` - Run linter (requires golangci-lint)

## Usage

```bash
tswitch <task-id>
```

Example:
```bash
tswitch 2
```

This will:
1. Find the currently active task
2. Stop it with `todo <active-id> stop`
3. Start the new task with `todo start 2`

## How it works

The tool parses the output of the `todo` command to identify which task is currently active (the one with a time value in the "Active" column), then executes the appropriate `todo` commands to switch tasks.

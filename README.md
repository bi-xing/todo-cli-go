# todo-cli-go

# MIRROR
For issues and pull requests, please visit the main repository.

https://codeberg.org/bixing/todo-cli-go

___

A minimal to-do list CLI written in Go. Manage tasks from your terminal with plain JSON persistence.

## Usage

| Command | Description |
| --- | --- |
| `todo list` | Show all current tasks with numbered indices. |
| `todo add <task>` | Add a new task. With no argument, opens an interactive prompt. |

Tasks are stored as a plain JSON array in `todo.json`.

## Commands

**`todo list`** — prints numbered tasks. If there are none, it says so quietly.

```
$ todo list
No tasks listed.

$ todo add "write the readme"
Task added: write the readme

$ todo list
1. write the readme
```

**`todo add [task]`** — accepts a task string directly as an argument. If called with no argument, a `huh` form prompt appears for interactive input.

```
$ todo add "review PR"
Task added: review PR
```

## Stack

- [Go](https://go.dev) 1.26 — language and toolchain
- [Cobra](https://github.com/spf13/cobra) — CLI framework
- [Huh](https://github.com/charmbracelet/huh) — terminal form (interactive add prompt)
- JSON file storage — zero dependencies beyond Go stdlib for persistence

## Project structure

```
todo-cli-go/
  main.go        — CLI entry point, command definitions, JSON I/O
  todo.json      — task storage (auto-created at runtime)
  datafile.json  — (placeholder, currently unused)
  go.mod / go.sum
  LICENSE        — MIT
```

## License

MIT — see `LICENSE` for details.

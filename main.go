package main

import (
	"fmt"
	"os"
  "log"
  "encoding/json"
  "github.com/spf13/cobra"
)

const datafile = "todo.json"

func loadTask() []string {
  tasks := []string
  data, err := os.ReadFile(datafile)
  if err == nil {
    json.Unmarshal(data, &tasks)
  }
  if err != nil {
    fmt.Println("No tasks listed.")
  }
  return tasks
}

func saveTask(tasks []string) {
  data, err := json.MarshalIndent(tasks, "", "\t")
  if err != nil {
    log.Fatalf("Error: %v", err)
  }
  os.WriteFile(datafile, data, 0644)
}

func main() {
  fmt.Printf(`
             `)
}

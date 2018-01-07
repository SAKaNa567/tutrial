
package main

import (
  "fmt"
  )

type Task struct {
  ID int
  detail string
  done bool
}

func main(){
  var task *Task = new(Task)
  task.ID = 1
  task.detail = "buy the milk"
  fmt.Println(task.done)
  }





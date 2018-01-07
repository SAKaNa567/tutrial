package main

import (
  "fmt"
)

type Task struct {
  ID int 
  Detail string
  done bool
}

func NewTask(id int, detail string) *Task {
  task := &Task {
    ID:id,
    Detail:detail,
    done:false,
    }
    return task
}


//func (task Task) String() string {
 // str := fmt.Sprintf("%d)%s",task.ID,task.Detail)
  //return str
 // }

func (task *Task) finish() {
  task.done = true
}


func main(){
  task :=  NewTask(1,"buy the milk")
  task.finish()
  fmt.Printf("%+v",task)
}
  

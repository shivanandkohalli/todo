package main

import (
	"fmt"

	"github.com/rs/xid"
)

type TodoListStruct struct {
	uniqueID   string
	task       string
	taskStatus bool
}

var todoList []TodoListStruct

// GetCompleteList Fetch the complete active todo list present
func GetCompleteList() []TodoListStruct {
	return todoList
}

func AddToList(task string) {
	todoTask := createNewTodo(task)
	todoList = append(todoList, todoTask)
}

func DeleteFromList(uniqueID string) {
	pos := findPosByUniqueID(uniqueID)
	if pos == -1 {
		fmt.Println("Error: Not present in TODO List")
		return
	}
	todoList = append(todoList[:pos], todoList[pos+1:]...)
}

func CompleteTask(uniqueID string) {
	pos := findPosByUniqueID(uniqueID)
	if pos == -1 {
		fmt.Println("Error: CompleteTask uuID not present in TODO List")
		return
	}
	todoList[pos].taskStatus = true
}
func findPosByUniqueID(uniqueID string) int {
	for i, todo := range todoList {
		if todo.uniqueID == uniqueID {
			return i
		}
	}
	return -1
}

func createNewTodo(task string) TodoListStruct {
	return TodoListStruct{
		uniqueID:   xid.New().String(),
		task:       task,
		taskStatus: false,
	}
}
func main() {
	fmt.Println("Todo App")
	AddToList("Exercise")
	AddToList("Meditate")
	list := GetCompleteList()
	fmt.Println(list)
	CompleteTask(list[0].uniqueID)
	fmt.Println(list)
}

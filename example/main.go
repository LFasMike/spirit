package main

import (
	"fmt"
	"github.com/gogap/spirit"

	"os"
)

func main() {
	todoSpirit := spirit.NewClassicSpirit("example", "a example of todo component", "1.0.0")

	todoComponent := spirit.NewBaseComponent("todo")

	todo := new(Todo)

	todoComponent.RegisterHandler("new_task", todo.NewTask)
	todoComponent.RegisterHandler("delete_task", todo.DeleteTask)
	todoComponent.RegisterHandler("done_task", todo.DoneTask)

	todoSpirit.RegisterHeartbeaters(new(ConsoleHeartbeater))

	todoSpirit.Hosting(todoComponent).Build().Run(initial)
}

func initial(configFile string) (err error) {
	//todo something initial before run
	env := os.Getenv("DEBUG_LEVEL")

	fmt.Println("DEBUG_LEVEL:", env)

	env2 := os.Getenv("DEBUG_LEVEL2")

	fmt.Println("DEBUG_LEVEL2:", env2)

	return
}

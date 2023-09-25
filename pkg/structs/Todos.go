package structs

type Todo struct {
	ID        int
	Task      string
	Completed bool
}

type AddTodo struct {
	Task      string
	Completed bool
}

package services

import (
	"database/sql"
	"fmt"

	"github.com/sep1ol/new-stack/pkg/structs"
)

type TodosService struct {
	db *sql.DB
}

func Todos(db *sql.DB) *TodosService {
	return &TodosService{
		db: db,
	}
}

func (s *TodosService) AddTodo(todo structs.AddTodo) (*structs.Todo, error) {
	query := `
		INSERT INTO todos (task, completed)
		VALUES ($1, $2)
		RETURNING id;
	`

	var ID int
	err := s.db.QueryRow(query, todo.Task, todo.Completed).Scan(&ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &structs.Todo{
		ID:        ID,
		Task:      todo.Task,
		Completed: todo.Completed,
	}, nil
}

func (s *TodosService) GetTodos() ([]structs.Todo, error) {
	var todos []structs.Todo
	query := `
		SELECT * FROM todos;
	`

	rows, err := s.db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var todo structs.Todo
		err := rows.Scan(&todo.ID, &todo.Task, &todo.Completed)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *TodosService) DeleteTodo(id int) error {
	query := `
		DELETE FROM todos WHERE id = $1
	`

	_, err := s.db.Exec(query, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

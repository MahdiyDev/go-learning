package todo

import "server/lib"

func Insert(conn lib.DatabaseConnection, todo *TodoMessage) error {
	err := conn.FetchOne(
		&TodoEntity{},
		"insert into todos(title, context) values ($1, $2) returning *",
		todo.GetTitle(),
		todo.GetContext(),
	)

	return err
}

func ReadAll(conn lib.DatabaseConnection, todo *TodoMessage) (*TodoMessageGetResponse, error) {
	todos := make([]TodoEntity, 0)
	todosRes := TodoMessageGetResponse{}
	err := conn.FetchAll(
		&todos,
		"select * from todos",
	)

	for i := 0; i < len(todos); i++ {
		todoMsg := TodoMessage{}
		todoMsg.Title = todos[i].Title
		todoMsg.Context = todos[i].Context
		todosRes.Todos = append(todosRes.Todos, &todoMsg)
	}

	return &todosRes, err
}

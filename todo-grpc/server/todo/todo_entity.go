package todo

type TodoEntity struct {
	ID      uint   `db:"id"`
	Title   string `db:"title"`
	Context string `db:"context"`
}

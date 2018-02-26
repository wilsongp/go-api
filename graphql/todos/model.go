package todos

// COLLECTION : The name of the collection in the db
const COLLECTION = "todos"

//Todo document model
type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

//Todos is an array of Todo
type Todos []Todo

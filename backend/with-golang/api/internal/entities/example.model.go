package entities

type Body struct {
	Message int `json:"message" binding:"required"`
}

type Example struct {
	Model
	Body
}

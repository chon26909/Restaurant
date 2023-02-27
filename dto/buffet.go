package dto

type CreateBuffet struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Price       int16  `json:"price"`
}

type UpdateBuffet struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Price       int16  `json:"price"`
}

type BuffetResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Price       int16  `json:"price"`
}

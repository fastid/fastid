package handlers

// Errors - structure for error response
type Errors struct {
	Message   string `json:"message"`
	Field     string `json:"field"`
	Tag       string `json:"tag"`
	ActualTag string `json:"actual_tag"`
}

type Error struct {
	Errors  []Errors `json:"errors,omitempty"`
	Message string   `json:"message,omitempty"`
}

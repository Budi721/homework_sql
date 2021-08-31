package web

type MovieRequest struct {
	Title       string `json:"title"`
	Slug        string `json:"slug" validate:"required"`
	Description string `json:"description"`
	Duration    int    `json:"duration" validate:"gt=0"`
	Image       string `json:"image"`
}

package web

type MovieRequest struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Image       string `json:"image"`
}

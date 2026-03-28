package models

type Error struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	// Other fields ignored because they aren't needed
}

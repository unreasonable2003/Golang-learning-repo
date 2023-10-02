package recipes

type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingrediant `json:"ingredients"`
}

type Ingrediant struct {
	Name string `json:"name"`
}

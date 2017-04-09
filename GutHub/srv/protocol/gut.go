package protocol

// Gut a sample cooking guide
type Gut struct {
	ID           int          `json:"id"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Ingredients  []Ingredient `json:"ingredients"`
	Instructions string       `json:"instructions"`
}

// Ingredient the composition of a meal
type Ingredient struct {
	Amount         string `json:"amount"`
	AmountUnits    string `json:"amountUnits"`
	IngredientName string `json:"ingredientName"`
}

package model


type  DietPlanSearchResult struct{
	Results      []*FoodData `json:"results"`
	Suggestions []*FoodData `json:"suggestions"`
}

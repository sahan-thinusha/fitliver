package model



type  FoodRequest struct{
	Keyword     string `json:"keyword"`
	AlcoholFree bool `json:"alcoholFree"`
	CeleryFree bool `json:"celeryFree"`
	CrustaceanFree bool `json:"crustaceanFree"`
	DairyFree bool `json:"dairyFree"`
	EggFree bool `json:"eggFree"`
	FishFree bool `json:"fishFree"`
	GlutenFree bool `json:"glutenFree"`
	Keto bool `json:"keto"`
	KidneyFriendly bool `json:"kidneyFriendly"`
	Kosher bool `json:"kosher"`
	LowPotassium bool `json:"lowPotassium"`
	LupineFree bool `json:"lupineFree"`
	MustardFree bool `json:"mustardFree"`
	NoOilAdded bool `json:"noOilAdded"`
	NoSugar bool `json:"noSugar"`
	Paleo bool `json:"paleo"`
	PeanutFree bool `json:"peanutFree"`
	Pescatarian bool `json:"pescatarian"`
	PorkFree bool `json:"porkFree"`
	RedMeatFree bool `json:"redMeatFree"`
	SesameFree bool `json:"sesameFree"`
	ShellfishFree bool `json:"shellfishFree"`
	SoyFree bool `json:"soyFree"`
	SugarConscious bool `json:"sugarConscious"`
	TreeNutFree bool `json:"treeNutFree"`
	Vegan bool `json:"vegan"`
	Vegetarian bool `json:"vegetarian"`
	WheatFree bool `json:"wheatFree"`
	Balanced bool `json:"balanced"`
	HighFiber bool `json:"highFiber"`
	HighProtein bool `json:"highProtein"`
	LowCarb bool `json:"lowCarb"`
	LowFat bool `json:"lowFat"`
	LowSodium bool `json:"lowSodium"`
}

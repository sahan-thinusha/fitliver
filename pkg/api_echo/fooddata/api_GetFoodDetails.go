package fooddata



import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	op "fitliver/pkg/operation/diet_plan"
	"github.com/labstack/echo/v4"
)

func GetFoodData(c echo.Context) (*model.DietPlanSearchResult, error) {
	foodData := model.FoodRequest{}
	if error := c.Bind(&foodData); error != nil {
		return nil, error
	}

	data := []string{}

	if foodData.AlcoholFree {
		data = append(data,env.AlcoholFree)
	}

	if foodData.CeleryFree {
		data = append(data,env.CeleryFree)
	}

	if foodData.CrustaceanFree {
		data = append(data,env.CrustaceanFree)
	}

	if foodData.DairyFree {
		data = append(data,env.DairyFree)
	}
	if foodData.EggFree{
		data = append(data,env.EggFree)
	}

	if foodData.FishFree {
		data = append(data,env.FishFree)
	}

	if foodData.GlutenFree {
		data = append(data,env.GlutenFree)
	}


	if foodData.WheatFree {
		data = append(data,env.WheatFree)
	}
	if foodData.Vegetarian{
		data = append(data,env.Vegetarian)
	}

	if foodData.Vegan {
		data = append(data,env.Vegan)
	}

	if foodData.TreeNutFree {
		data = append(data,env.TreeNutFree)
	}

	if foodData.SugarConscious {
		data = append(data,env.SugarConscious)
	}
	if foodData.SoyFree{
		data = append(data,env.SoyFree)
	}

	if foodData.ShellfishFree {
		data = append(data,env.ShellfishFree)
	}

	if foodData.SesameFree {
		data = append(data,env.SesameFree)
	}

	if foodData.RedMeatFree {
		data = append(data,env.RedMeatFree)
	}
	if foodData.PorkFree{
		data = append(data,env.PorkFree)
	}

	if foodData.Pescatarian {
		data = append(data,env.Pescatarian)
	}

	if foodData.PeanutFree {
		data = append(data,env.PeanutFree)
	}

	if foodData.Paleo {
		data = append(data,env.Paleo)
	}
	if foodData.NoSugar{
		data = append(data,env.NoSugar)
	}

	if foodData.MustardFree {
		data = append(data,env.MustardFree)
	}

	if foodData.LupineFree {
		data = append(data,env.LupineFree)
	}

	if foodData.NoOilAdded {
		data = append(data,env.NoOilAdded)
	}


	if foodData.Kosher {
		data = append(data,env.Kosher)
	}


	result := op.GetFoodData(foodData.Keyword,data)
	return result, nil
}

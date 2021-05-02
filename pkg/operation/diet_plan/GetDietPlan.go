package diet_plan

import (
	"encoding/json"
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"github.com/go-resty/resty/v2"
)


type FoodResult struct {
	Text   string `json:"text"`
	Parsed []struct {
		Food struct {
			Foodid    string `json:"foodId"`
			Label     string `json:"label"`
			Nutrients struct {
				EnercKcal float64     `json:"ENERC_KCAL"`
				Procnt    float64 `json:"PROCNT"`
				Fat       float64 `json:"FAT"`
				Chocdf    float64     `json:"CHOCDF"`
				Fibtg     float64     `json:"FIBTG"`
			} `json:"nutrients"`
			Category      string `json:"category"`
			Categorylabel string `json:"categoryLabel"`
			Image         string `json:"image"`
		} `json:"food"`
	} `json:"parsed"`
	Hints []struct {
		Food struct {
			Foodid    string `json:"foodId"`
			Label     string `json:"label"`
			Nutrients struct {
				EnercKcal float64     `json:"ENERC_KCAL"`
				Procnt    float64 `json:"PROCNT"`
				Fat       float64 `json:"FAT"`
				Chocdf    float64     `json:"CHOCDF"`
				Fibtg     float64     `json:"FIBTG"`
			} `json:"nutrients"`
			Category      string `json:"category"`
			Categorylabel string `json:"categoryLabel"`
			Image         string `json:"image"`
		} `json:"food,omitempty"`
		Measures []struct {
			URI       string `json:"uri"`
			Label     string `json:"label"`
			Weight    int    `json:"weight"`
			Qualified []struct {
				Qualifiers []struct {
					URI   string `json:"uri"`
					Label string `json:"label"`
				} `json:"qualifiers"`
				Weight int `json:"weight"`
			} `json:"qualified,omitempty"`
		} `json:"measures"`
	} `json:"hints"`
	Links struct {
		Next struct {
			Title string `json:"title"`
			Href  string `json:"href"`
		} `json:"next"`
	} `json:"_links"`
}

func GetFoodData(keyword string,params []string) *model.DietPlanSearchResult {
	client := resty.New()
	queryParams := make(map[string]string)
	queryParams["app_id"] = env.EAppId
	queryParams["app_key"] =env.EApiKey
	queryParams["ingr"] = keyword
	for _,data := range params{
		queryParams["health"] = data
	}
		resp, _ := client.R().
		SetQueryParams(queryParams).
		Get(env.FoodAPI)
	data := FoodResult{}
	if resp.StatusCode() >= 200 && resp.StatusCode() <= 204 {
		_ = json.Unmarshal([]byte(resp.String()), &data)
	}
	foodData := []*model.FoodData{}
	for _,food := range data.Parsed{
		fd := model.FoodData{}
		fd.Nutrients = struct {
			EnercKcal float64     `json:"enerc_kcal"`
			Procnt    float64 `json:"procnt"`
			Fat       float64 `json:"fat"`
			Chocdf    float64     `json:"chocdf"`
			Fibtg     float64     `json:"fibtg"`
		}(food.Food.Nutrients)
		fd.Category = food.Food.Category
		fd.Categorylabel = food.Food.Categorylabel
		fd.Label = food.Food.Label
		foodData = append(foodData,&fd)
	}
	foodDataSuggestions := []*model.FoodData{}
	for _,food := range data.Hints{
		fd := model.FoodData{}
		fd.Nutrients = struct {
			EnercKcal float64     `json:"enerc_kcal"`
			Procnt    float64 `json:"procnt"`
			Fat       float64 `json:"fat"`
			Chocdf    float64     `json:"chocdf"`
			Fibtg     float64     `json:"fibtg"`
		}(food.Food.Nutrients)
		fd.Category = food.Food.Category
		fd.Categorylabel = food.Food.Categorylabel
		fd.Label = food.Food.Label
		foodDataSuggestions = append(foodDataSuggestions,&fd)
	}
	dietresult := model.DietPlanSearchResult{}
	dietresult.Results = foodData
	dietresult.Suggestions = foodDataSuggestions
	return &dietresult
}
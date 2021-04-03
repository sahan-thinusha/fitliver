package model

type BMI struct {
	Bmi float64 `json:"bmi"`
	RiskCategory string `json:"riskcategory"`
	NormalWeight float64 `json:"normalweight"`
	BMR float64 `json:"bmr"`
	Remark string `json:"remark"`
}

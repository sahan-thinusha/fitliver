package model

type BMI struct {
	Bmi float64 `json:"bmi"`
	RiskCategory string `json:"riskcategory"`
	NormalWeight float64 `json:"normalweight"`
	Remark string `json:"remark"`
}

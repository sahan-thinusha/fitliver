package bmi

import (
	"fitliver/pkg/model"
	"fmt"
	"math"
)

func CalculateBMI(weight float64,height float64)  (*model.BMI,error){
	bmi := weight / math.Pow(height, 2)


	riskCategory := ""
	if bmi < float64(18.5) {
		riskCategory = "Underweight"
	} else if bmi < 25 {
		riskCategory = "Normal weight"
	} else if bmi < 30 {
		riskCategory = "Overweight"
	} else {
		riskCategory = "Obese"
	}

	// calculate normal weight based on height and bmi = 25
	normalWeight := 25 * math.Pow(height, 2)
	delta := weight - normalWeight

	remark := ""
	if (delta > 0) && (bmi > 30) {
		remark = fmt.Sprintf("You need to reduce %0.2v kilograms.\n", math.Abs(delta))
	}

	if (delta < 0) && (bmi < float64(18.5)) {
		remark = fmt.Sprintf("You need to increase %0.2v kilograms.\n", math.Abs(delta))
	}

	bmiResult := model.BMI{}
	bmiResult.Bmi = bmi
	bmiResult.NormalWeight = normalWeight
	bmiResult.RiskCategory = riskCategory
	bmiResult.Remark = remark

	return &bmiResult,nil
}
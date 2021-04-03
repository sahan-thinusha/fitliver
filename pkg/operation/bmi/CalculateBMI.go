package bmi

import (
	"fitliver/pkg/model"
	"fmt"
	"math"
)

func CalculateBMI(weight float64,height float64)  (*model.BMI,error){
	bmi := weight / math.Pow(height, 2)


	riskCategory := ""
	if bmi < 18.5 {
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

func calcOHarrisBenedict(wKG float64, hM float64, gender string, age uint8) (oHBBMR float64) {
	if gender == "male" {
		oHBBMR = 66.4730 + (13.7515 * wKG) + (5.0033 * (hM * 100)) - (6.755 * float64(age))
		oHBBMR = math.Round((oHBBMR * 10) / 10)
	}
	if gender == "female" {
		oHBBMR = 655.0955 + (9.5634 * wKG) + (1.8496 * (hM * 100)) - (4.6756 * float64(age))
		oHBBMR = math.Round((oHBBMR * 10) / 10)
	}
	return
}

/*
	THIS FUNCTION CALCULATES THE REVISED HARRIS-BENEDICT BMR EQUATION
		MEN - BMR = 88.362 + (13.397 x weight in kg) + (4.799 x height in cm) - (5.677 x age in years)
		WOMEN - BMR = 447.593 + (9.247 x weight in kg) + (3.098 x height in cm) - (4.330 x age in years)
*/
func  calcRHarrisBenedict(wKG float64, hM float64, gender string, age uint8) (rHBBMR float64) {
	if gender == "male" {
		rHBBMR = math.Round(((88.362 + (13.397 * wKG) + (4.799 * (hM * 100)) - (5.677 * float64(age))) * 10) / 10)
	}
	if gender == "female" {
		rHBBMR = math.Round(((447.593 + (9.247 * wKG) + (3.098 * (hM * 100)) - (4.330 * float64(age))) * 10) / 10)
	}
	return
}

/*
	THIS FUNCTION CALCULATES THE MIFFLIN-ST JEOR BMR EQUATION
		MEN - BMR (metric) = (10 × weight in kg) + (6.25 × height in cm) - (5 × age in years) + 5
		WOMEN - BMR (metric) = (10 × weight in kg) + (6.25 × height in cm) - (5 × age in years) - 161
*/
func  calcMifflinStJeor(wKG float64, hM float64, gender string, age uint8) (mjBMR float64) {
	if gender == "male" {
		mjBMR = math.Round((((10 * wKG) + (6.25 * (hM * 100)) - (5 * float64(age)) + 5) * 10) / 10)
	}
	if gender == "female" {
		mjBMR = math.Round((((10 * wKG) + (6.25 * (hM * 100)) - (5 * float64(age)) - 161) * 10) / 10)
	}
	return
}

func  calcTDEE(a float64) (tdee float64,mjBMR float64) {
	tdee = math.Round(((mjBMR * a) * 10) / 10)
	return
}

func  subPFromTDEE(tdee float64, percent float64) (calTar float64) {
	pCals := ((percent * tdee) / 100)
	calTar = math.Round(((tdee - pCals) * 10) / 10)
	return
}

func addPFromTDEE(tdee float64, percent float64) (calTar float64) {
	pCals := ((percent * tdee) / 100)
	calTar = math.Round(((tdee + pCals) * 10) / 10)
	return
}
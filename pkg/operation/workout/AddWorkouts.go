package workout

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func AddWorkout(){
	db :=env.RDB
	workouts := []*model.Workout{}
	db.Find(&workouts)
	if len(workouts) == 0 {
		w1 := model.Workout{}
		w1.Name = "2Full Body Strength Workout"
		w1.Reps = "20 min"
		w1.VideoLink = "https://www.youtube.com/watch?v=N4HbeyDChFw"
		db.Create(&w1)

		w2 := model.Workout{}
		w2.Name = "Full Body Dumbbell Workout"
		w2.Reps = "15 min"
		w2.VideoLink = "https://www.youtube.com/watch?v=xqVBoyKXbsA"
		db.Create(&w2)

		w3 := model.Workout{}
		w3.Name = " Home Strength Workout"
		w3.Reps = "30 min"
		w3.VideoLink = "https://www.youtube.com/watch?v=YdB1HMCldJY"
		db.Create(&w3)

		w4 := model.Workout{}
		w4.Name = "Beginner's At-Home Cardio Workout"
		w4.Reps = "15 min"
		w4.VideoLink = "https://www.youtube.com/watch?v=VHyGqsPOUHs"
		db.Create(&w4)

		w5 := model.Workout{}
		w5.Name = "Cardio Dance Fitness Workout"
		w5.Reps = "15 min"
		w5.VideoLink = "https://www.youtube.com/watch?v=yN3GgCUmmXw&ab_channel=TheStudiobyJamieKinkeade"
		db.Create(&w5)

		w6 := model.Workout{}
		w6.Name = "Aerobic Reduction of Belly Fat Quickly | Aerobic Everyday for Best Body Shape"
		w6.Reps = "24 min"
		w6.VideoLink = "https://www.youtube.com/watch?v=HSO-PjkqBY4"
		db.Create(&w6)

		w7 := model.Workout{}
		w7.Name = "Total Body Strength Training"
		w7.Reps = "30 min"
		w7.VideoLink = "https://www.youtube.com/watch?v=mUns8O4YL5M"
		db.Create(&w7)

		w8 := model.Workout{}
		w8.Name = "HIIT Cardio Workout"
		w8.Reps = "30 min"
		w8.VideoLink = "https://www.youtube.com/watch?v=ml6cT4AZdqI"
		db.Create(&w8)
	}

}


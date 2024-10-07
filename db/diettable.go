package db

// type UserWorkoutDiet struct {
// 	BodyTypeID int    `gorm:"column:body_type_id"`
// 	AgeGroupID int    `gorm:"column:age_group_id"`
// 	WorkOut    int    `gorm:"column:workout"`
// 	Diet       string `gorm:"column:diet"`
// }

// // TableName overrides the table name used by GORM
// func (UserWorkoutDiet) TableName() string {
// 	return "user_workout_diet"
// }

// // Group 1 is Age less than 20
// // Group 2 is Age between 21 and  40
// // Group 3 is Age above 41 a

// // Body type 1 Underweight = <18.5
// // Body type 2 Normal weight = 18.5–24.9
// // Body type 3 Overweight = 25–29.9

// var userWorkoutDiet = []UserWorkoutDiet{
// 	{
// 		BodyTypeID: 1,
// 		AgeGroupID: 1,
// 		WorkOut:    1,
// 		Diet:       "cat,dog,duck",
// 	},
// 	{
// 		BodyTypeID: 2,
// 		AgeGroupID: 1,
// 		WorkOut:    2,
// 		Diet:       "cat,dog,duck",
// 	},
// 	{
// 		BodyTypeID: 3,
// 		AgeGroupID: 1,
// 		WorkOut:    3,
// 		Diet:       "cat,dog,duck",
// 	},
// 	{
// 		BodyTypeID: 1,
// 		AgeGroupID: 2,
// 		WorkOut:    4,
// 		Diet:       "cat,dog,duck",
// 	},
// 	{
// 		BodyTypeID: 2,
// 		AgeGroupID: 2,
// 		WorkOut:    5,
// 		Diet:       "cat,dog,duck",
// 	},
// 	{
// 		BodyTypeID: 3,
// 		AgeGroupID: 2,
// 		WorkOut:    6,
// 		Diet:       "cat,dog,duck",
// 	},
// 	{
// 		BodyTypeID: 1,
// 		AgeGroupID: 3,
// 		WorkOut:    7,
// 		Diet:       "cat,dog,duck",
// 	},
// 	{
// 		BodyTypeID: 2,
// 		AgeGroupID: 3,
// 		WorkOut:    8,
// 		Diet:       "cat,dog,duck",
// 	},
// 	{
// 		BodyTypeID: 3,
// 		AgeGroupID: 3,
// 		WorkOut:    9,
// 		Diet:       "cat,dog,duck",
// 	},
// }

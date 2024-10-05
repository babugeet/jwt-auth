package models

type User struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	Age       int64  `json:"age"`
	Height    int64  `json:"height"`
	Weight    int64  `json:"weight"`
	BMI       int64  `json:"bmi"`
}

type WorkOutPlan struct {
}

type FoodChart struct {
}

type Database interface {
	GetUser(userId string) string
	PutUser()
	// UpdateUser()
}

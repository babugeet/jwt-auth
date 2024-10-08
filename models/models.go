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
type Weekday struct {
	// ID        uint   `gorm:"primaryKey"`
	Monday    string `gorm:"not null"`
	Tuesday   string `gorm:"not null"`
	Wednesday string `gorm:"not null"`
	Thursday  string `gorm:"not null"`
	Friday    string `gorm:"not null"`
	Saturday  string `gorm:"not null"`
	Sunday    string `gorm:"not null"`
}

type CardioList struct {
	Squats       string `json:"squats,omitempty"`
	Deadlift     string `json:"deadlift,omitempty"`
	PushUps      string `json:"pushups,omitempty"`
	PullUps      string `json:"pullups,omitempty"`
	JumpingJacks string `json:"jumpingjacks,omitempty"`
	WeightLift   string `json:"weightlift,omitempty"`
	BenchPress   string `json:"benchpress,omitempty"`
	Lunges       string `json:"lunges,omitempty"`
	LegPress     string `json:"legpress,omitempty"`
	Running      string `json:"running,omitempty"`
	Cycle        string `json:"cycle,omitempty"`
	Swimming     string `json:"swimming,omitempty"`
	Walking      string `json:"walking,omitempty"`
}
type DietPlan struct {
	// gorm.Model
	// Day       string `json:"day,omitempty"`       // Day of the week (Monday, Tuesday, etc.)
	Breakfast string `json:"breakfast,omitempty"` // Breakfast meal
	Lunch     string `json:"lunch,omitempty"`     // Lunch meal
	Dinner    string `json:"dinner,omitempty"`    // Dinner meal
}

type Database interface {
	GetUser(username string) (bool, User)
	AddUser(user User) error
	GetUserWorkOutCardioPlanfromDB(bodyTypeID int, ageGroupID int) ([]Weekday, []Weekday)
	GetUserDietPlanfromDB(bodyTypeID int, ageGroupID int) DietPlan

	GetReps(ageGroupID int, excerciseType string) (string, string)
	// Close() error
	// UpdateUser()
}

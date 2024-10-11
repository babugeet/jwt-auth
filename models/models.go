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
	Pushups      string `json:"pushups,omitempty"`
	Pullups      string `json:"pullups,omitempty"`
	Jumpingjacks string `json:"jumpingjacks,omitempty"`
	Weightlift   string `json:"weightlift,omitempty"`
	Benchpress   string `json:"benchpress,omitempty"`
	Lunges       string `json:"lunges,omitempty"`
	Legpress     string `json:"legpress,omitempty"`
	Running      string `json:"running,omitempty"`
	Cycling      string `json:"cycling,omitempty"`
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

type Workoutplan struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	Date             string `gorm:"uniqueIndex:idx_user_date" json:"date"` // Use time.Time for date operations
	Username         string `gorm:"uniqueIndex:idx_user_date" json:"username"`
	Squatsdone       int    `json:"squatsdone"`
	Deadliftdone     int    `json:"deadliftdone"`
	Pushupsdone      int    `json:"pushupsdone"`
	Pullupsdone      int    `json:"pullupsdone"`
	Jumpingjacksdone int    `json:"jumpingjacksdone"`
	Weightliftdone   int    `json:"weightliftdone,omitempty"`
	Benchpressdone   int    `json:"benchpressdone,omitempty"`
	Lungesdone       int    `json:"lungesdone,omitempty"`
	Legpressdone     int    `json:"legpressdone,omitempty"`
	Runningdone      int    `json:"runningdone,omitempty"`
	Cyclingdone      int    `json:"cyclingdone,omitempty"`
	Swimmingdone     int    `json:"swimmingdone,omitempty"`
	Walkingdone      int    `json:"walkingdone,omitempty"`
	Squats           string `json:"squats,omitempty"`
	Deadlift         string `json:"deadlift,omitempty"`
	Pushups          string `json:"pushups,omitempty"`
	Pullups          string `json:"pullups,omitempty"`
	Jumpingjacks     string `json:"jumpingjacks,omitempty"`
	Weightlift       string `json:"weightlift,omitempty"`
	Benchpress       string `json:"benchpress,omitempty"`
	Lunges           string `json:"lunges,omitempty"`
	Legpress         string `json:"legpress,omitempty"`
	Running          string `json:"running,omitempty"`
	Cycling          string `json:"cycling,omitempty"`
	Swimming         string `json:"swimming,omitempty"`
	Walking          string `json:"walking,omitempty"`
	Water            int    `json:"water,omitempty"`
	Steps            int    `json:"steps,omitempty"`
}

type Workouttodaylist struct {
	Workout []Workouttoday
}
type Workouttoday struct {
	Name   string
	Target string
	Done   int
}

type Database interface {
	GetUser(username string) (bool, User)
	AddUser(user User) error
	GetUserWorkOutCardioPlanfromDB(bodyTypeID int, ageGroupID int) ([]Weekday, []Weekday)
	GetUserDietPlanfromDB(bodyTypeID int, ageGroupID int) DietPlan

	GetReps(ageGroupID int, excerciseType string) (string, string)
	WriteTarget2DB(user string, workoutplan *Workoutplan) error
	GetUserWorkoutDetails4mDB(user string) Workoutplan
	GetUserWorkoutDetails42day4mDB(user string, cardio []byte, workout []byte) Workouttodaylist
	// Close() error
	// UpdateUser()
}

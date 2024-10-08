package db

type Workout struct {
	ID         uint `gorm:"primaryKey"`
	BodyTypeID int  `gorm:"not null"`
	AgeGroupID int  `gorm:"not null"`
	Cardio     int  `gorm:"not null"`
	Workout    int  `gorm:"not null"`
}

// TableName overrides the table name used by GORM
func (Workout) TableName() string {
	return "usertype_workout_map"
}

var workouts = []Workout{
	{ID: 1, BodyTypeID: 1, AgeGroupID: 1, Cardio: 1, Workout: 1},
	{ID: 2, BodyTypeID: 2, AgeGroupID: 1, Cardio: 1, Workout: 1},
	{ID: 3, BodyTypeID: 3, AgeGroupID: 1, Cardio: 1, Workout: 1},
	{ID: 4, BodyTypeID: 1, AgeGroupID: 2, Cardio: 2, Workout: 2},
	{ID: 5, BodyTypeID: 2, AgeGroupID: 2, Cardio: 2, Workout: 2},
	{ID: 6, BodyTypeID: 3, AgeGroupID: 2, Cardio: 2, Workout: 2},
	{ID: 7, BodyTypeID: 1, AgeGroupID: 3, Cardio: 3, Workout: 3},
	{ID: 8, BodyTypeID: 2, AgeGroupID: 3, Cardio: 3, Workout: 3},
	{ID: 9, BodyTypeID: 3, AgeGroupID: 3, Cardio: 3, Workout: 3},
}

type WorkoutSchedule1 struct {
	ID uint `gorm:"primaryKey"`
	// ExerciseOrder int    `gorm:"not null"` // Order of the exercise
	Monday    string `gorm:"not null"`
	Tuesday   string `gorm:"not null"`
	Wednesday string `gorm:"not null"`
	Thursday  string `gorm:"not null"`
	Friday    string `gorm:"not null"`
	Saturday  string `gorm:"not null"`
	Sunday    string `gorm:"not null"`
}

// TableName overrides the table name used by GORM
func (WorkoutSchedule1) TableName() string {
	return "workout_schedule1"
}

type WorkoutSchedule2 struct {
	ID uint `gorm:"primaryKey"`
	// ExerciseOrder int    `gorm:"not null"` // Order of the exercise
	Monday    string `gorm:"not null"`
	Tuesday   string `gorm:"not null"`
	Wednesday string `gorm:"not null"`
	Thursday  string `gorm:"not null"`
	Friday    string `gorm:"not null"`
	Saturday  string `gorm:"not null"`
	Sunday    string `gorm:"not null"`
}

// TableName overrides the table name used by GORM
func (WorkoutSchedule2) TableName() string {
	return "workout_schedule2"
}

type WorkoutSchedule3 struct {
	ID uint `gorm:"primaryKey"`
	// ExerciseOrder int    `gorm:"not null"` // Order of the exercise
	Monday    string `gorm:"not null"`
	Tuesday   string `gorm:"not null"`
	Wednesday string `gorm:"not null"`
	Thursday  string `gorm:"not null"`
	Friday    string `gorm:"not null"`
	Saturday  string `gorm:"not null"`
	Sunday    string `gorm:"not null"`
}

// TableName overrides the table name used by GORM
func (WorkoutSchedule3) TableName() string {
	return "workout_schedule3"
}

var schedules1 = []WorkoutSchedule1{
	{ID: 1, Monday: "squats", Tuesday: "jacks", Wednesday: "legpress", Thursday: "squats", Friday: "jacks", Saturday: "legpress", Sunday: "squats"},
	{ID: 2, Monday: "pushups", Tuesday: "weightlift", Wednesday: "weightlift", Thursday: "pushups", Friday: "weightlift", Saturday: "weightlift", Sunday: "pushups"},
	{ID: 3, Monday: "pullups", Tuesday: "benchpress", Wednesday: "squats", Thursday: "Pullups", Friday: "benchpress", Saturday: "squats", Sunday: "pullups"},
}

var schedules2 = []WorkoutSchedule2{
	{ID: 1, Monday: "squats", Tuesday: "jacks", Wednesday: "legpress", Thursday: "squats", Friday: "jacks", Saturday: "legpress", Sunday: "squats"},
	{ID: 2, Monday: "pushups", Tuesday: "weightlift", Wednesday: "weightlift", Thursday: "pushups", Friday: "weightlift", Saturday: "weightlift", Sunday: "pushups"},
	{ID: 3, Monday: "pullups", Tuesday: "benchpress", Wednesday: "squats", Thursday: "Pullups", Friday: "benchpress", Saturday: "squats", Sunday: "pullups"},
}

var schedules3 = []WorkoutSchedule3{
	{ID: 1, Monday: "squats", Tuesday: "jacks", Wednesday: "legpress", Thursday: "squats", Friday: "jacks", Saturday: "legpress", Sunday: "squats"},
	{ID: 2, Monday: "pushups", Tuesday: "weightlift", Wednesday: "weightlift", Thursday: "pushups", Friday: "weightlift", Saturday: "weightlift", Sunday: "pushups"},
	{ID: 3, Monday: "pullups", Tuesday: "benchpress", Wednesday: "squats", Thursday: "pullups", Friday: "benchpress", Saturday: "squats", Sunday: "pullups"},
}

type StrengthSchedule1 struct {
	ID uint `gorm:"primaryKey"`
	// ExerciseOrder int    `gorm:"not null"` // Order of the exercise
	Monday    string `gorm:"not null"`
	Tuesday   string `gorm:"not null"`
	Wednesday string `gorm:"not null"`
	Thursday  string `gorm:"not null"`
	Friday    string `gorm:"not null"`
	Saturday  string `gorm:"not null"`
	Sunday    string `gorm:"not null"`
}

func (StrengthSchedule1) TableName() string {
	return "strength_schedule1"
}

type StrengthSchedule2 struct {
	ID uint `gorm:"primaryKey"`
	// ExerciseOrder int    `gorm:"not null"` // Order of the exercise
	Monday    string `gorm:"not null"`
	Tuesday   string `gorm:"not null"`
	Wednesday string `gorm:"not null"`
	Thursday  string `gorm:"not null"`
	Friday    string `gorm:"not null"`
	Saturday  string `gorm:"not null"`
	Sunday    string `gorm:"not null"`
}

func (StrengthSchedule2) TableName() string {
	return "strength_schedule2"
}

type StrengthSchedule3 struct {
	ID uint `gorm:"primaryKey"`
	// ExerciseOrder int    `gorm:"not null"` // Order of the exercise
	Monday    string `gorm:"not null"`
	Tuesday   string `gorm:"not null"`
	Wednesday string `gorm:"not null"`
	Thursday  string `gorm:"not null"`
	Friday    string `gorm:"not null"`
	Saturday  string `gorm:"not null"`
	Sunday    string `gorm:"not null"`
}

func (StrengthSchedule3) TableName() string {
	return "strength_schedule3"
}

var strengthschedules1 = []StrengthSchedule1{
	{ID: 1, Monday: "running", Tuesday: "running", Wednesday: "running", Thursday: "running", Friday: "running", Saturday: "running", Sunday: "running"},
	{ID: 2, Monday: "cycling", Tuesday: "cycling", Wednesday: "cycling", Thursday: "cycling", Friday: "cycling", Saturday: "cycling", Sunday: "cycling"},
	{ID: 3, Monday: "swimming", Tuesday: "swimming", Wednesday: "swimming", Thursday: "swimming", Friday: "swimming", Saturday: "swimming", Sunday: "swimming"},
}

var strengthschedules2 = []StrengthSchedule2{
	{ID: 1, Monday: "running", Tuesday: "running", Wednesday: "running", Thursday: "running", Friday: "running", Saturday: "running", Sunday: "running"},
	{ID: 2, Monday: "cycling", Tuesday: "cycling", Wednesday: "cycling", Thursday: "cycling", Friday: "cycling", Saturday: "cycling", Sunday: "cycling"},
	{ID: 3, Monday: "swimming", Tuesday: "swimming", Wednesday: "swimming", Thursday: "swimming", Friday: "swimming", Saturday: "swimming", Sunday: "swimming"},
}

var strengthschedules3 = []StrengthSchedule3{
	{ID: 1, Monday: "walking", Tuesday: "walking", Wednesday: "walking", Thursday: "walking", Friday: "walking", Saturday: "walking", Sunday: "walking"},
	{ID: 2, Monday: "cycling", Tuesday: "cycling", Wednesday: "cycling", Thursday: "cycling", Friday: "cycling", Saturday: "cycling", Sunday: "cycling"},
	{ID: 3, Monday: "swimming", Tuesday: "swimming", Wednesday: "swimming", Thursday: "swimming", Friday: "swimming", Saturday: "swimming", Sunday: "swimming"},
}

type ExerciseData struct {
	ID           uint   `gorm:"primaryKey"`
	Age          int    `gorm:"not null"`
	Running      string `gorm:"not null"`
	Cycling      string `gorm:"not null"`
	Swimming     string `gorm:"not null"`
	Walking      string `gorm:"not null"`
	Squats       string `gorm:"not null"`
	Deadlift     string `gorm:"not null"`
	Pushups      string `gorm:"not null"`
	Pullups      string `gorm:"not null"`
	Jumpingjacks string `gorm:"not null"`
	Weightlift   string `gorm:"not null"`
	Benchpress   string `gorm:"not null"`
	Lunges       string `gorm:"not null"`
	Legpress     string `gorm:"not null"`
}

func (ExerciseData) TableName() string {
	return "exercise_data"
}

var exercises = []ExerciseData{
	{ID: 1, Age: 1, Running: "3000 m", Cycling: "5000 m", Swimming: "300 m", Walking: "5000 m", Squats: "40 reps", Deadlift: "50 reps", Pushups: "40 reps", Pullups: "15 reps", Jumpingjacks: "40 reps", Weightlift: "50 kg", Benchpress: "50 kg", Lunges: "40 reps", Legpress: "30 kg"},
	{ID: 2, Age: 2, Running: "2000 m ", Cycling: "3000 m", Swimming: "200 m", Walking: "3000 m", Squats: "20 reps", Deadlift: "25 reps", Pushups: "10 reps", Pullups: "10 reps", Jumpingjacks: "20 reps", Weightlift: "30 kg", Benchpress: "30 kg", Lunges: "25 reps", Legpress: "20 kg"},
	{ID: 3, Age: 3, Running: "1000 m", Cycling: "1000 m", Swimming: "100 m", Walking: "1500 m", Squats: "0 reps", Deadlift: "0 reps", Pushups: "0 reps", Pullups: "0 reps", Jumpingjacks: "0 reps", Weightlift: "20 kg", Benchpress: "20 kg", Lunges: "15 reps", Legpress: "20 kg"},
}

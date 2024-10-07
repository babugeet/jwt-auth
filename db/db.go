package db

import (
	"fmt"
	"jwt-auth/models"
	"jwt-auth/variables"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Mydb struct {
	db *gorm.DB
}

const Exercise_table = "exercise_data"

type Table struct {
	Username  string `gorm:"primaryKey"`
	Firstname string
	Lastname  string
	Password  string
	Gender    string
	Age       int64
	Height    int64
	Weight    int64
	BMI       int64
}

// Custom table name method
func (Table) TableName() string {
	return "user_database" // Custom table name
}

func NewMydb() models.Database {
	dsn := "host=localhost user=postgres password=Abhi@1234 dbname=postgres port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Table{})
	if err != nil {
		fmt.Printf("Error migrating the User table in database: %v", err)
	}
	err = db.AutoMigrate(&Workout{})
	if err != nil {
		fmt.Printf("Error migrating the UserWorkoutDiet table in database: %v", err)
	}
	for _, j := range workouts {
		result := db.Create(&j)
		if result.Error != nil {
			fmt.Printf("failed to insert data: %v", result.Error)
		}
	}
	// Workout plan 1
	// Workout plan 2
	// Workout plan 3
	err = db.AutoMigrate(&WorkoutSchedule1{})
	if err != nil {
		fmt.Printf("Error migrating the UserWorkoutDiet table in database: %v", err)
	}
	err = db.AutoMigrate(&WorkoutSchedule2{})
	if err != nil {
		fmt.Printf("Error migrating the UserWorkoutDiet table in database: %v", err)
	}
	err = db.AutoMigrate(&WorkoutSchedule3{})
	if err != nil {
		fmt.Printf("Error migrating the UserWorkoutDiet table in database: %v", err)
	}

	for _, j := range schedules1 {
		result := db.Create(&j)
		if result.Error != nil {
			fmt.Printf("failed to insert data: %v", result.Error)
		}
	}
	for _, j := range schedules2 {
		result := db.Create(&j)
		if result.Error != nil {
			fmt.Printf("failed to insert data: %v", result.Error)
		}
	}
	for _, j := range schedules3 {
		result := db.Create(&j)
		if result.Error != nil {
			fmt.Printf("failed to insert data: %v", result.Error)
		}
	}
	// StrngthPlan
	err = db.AutoMigrate(&StrengthSchedule1{})
	if err != nil {
		fmt.Printf("Error migrating the UserWorkoutDiet table in database: %v", err)
	}
	err = db.AutoMigrate(&StrengthSchedule2{})
	if err != nil {
		fmt.Printf("Error migrating the UserWorkoutDiet table in database: %v", err)
	}
	err = db.AutoMigrate(&StrengthSchedule3{})
	if err != nil {
		fmt.Printf("Error migrating the UserWorkoutDiet table in database: %v", err)
	}

	for _, j := range strengthschedules1 {
		fmt.Println(j)
		result := db.Create(&j)
		if result.Error != nil {
			fmt.Printf("failed to insert data: %v", result.Error)
		}
	}
	for _, j := range strengthschedules2 {
		result := db.Create(&j)
		if result.Error != nil {
			fmt.Printf("failed to insert data: %v", result.Error)
		}
	}
	for _, j := range strengthschedules3 {
		result := db.Create(&j)
		if result.Error != nil {
			fmt.Printf("failed to insert data: %v", result.Error)
		}
	}

	err = db.AutoMigrate(&ExerciseData{})
	if err != nil {
		fmt.Printf("Error migrating the ExerciseData table in database: %v", err)
	}
	for _, exercise := range exercises {
		db.Create(&exercise)
	}
	log.Println("Data inserted successfully!")
	return &Mydb{db}
}

func ConvertToUser(t Table) models.User {
	return models.User{
		Username:  t.Username,
		Firstname: t.Firstname,
		Lastname:  t.Lastname,
		Password:  t.Password,
		Gender:    t.Gender,
		Age:       t.Age,
		Height:    t.Height,
		Weight:    t.Weight,
		BMI:       t.BMI,
	}
}

func (m Mydb) GetUser(username string) (bool, models.User) {
	var user Table
	fmt.Println("Reached GetUser")

	// Use First to get the first matching row (if exists)
	result := m.db.Find(&user, "username = ?", username)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("User not found")
			return false, models.User{}
		}
		fmt.Println("Error fetching user: " + result.Error.Error())
		return false, models.User{}
	}

	return true, ConvertToUser(user)
}

func (m Mydb) AddUser(user models.User) error {
	// Map models.User to Table struct
	newUser := Table{
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Password:  user.Password,
		Gender:    user.Gender,
		Age:       user.Age,
		Height:    user.Height,
		Weight:    user.Weight,
		BMI:       user.BMI,
	}

	result := m.db.Create(&newUser)
	if result.Error != nil {
		fmt.Println("Failed to create user: " + result.Error.Error())
		return result.Error
	}
	return nil
}

func (m *Mydb) GetUserWorkOutCardioPlanfromDB(bodyTypeID int, ageGroupID int) ([]models.Weekday, []models.Weekday) {
	var plan Workout
	// bodyTypeID := 1 // Example BodyTypeID
	// ageGroupID := 2 // Example AgeGroupID

	result := m.db.Where("body_type_id = ? AND age_group_id = ?", bodyTypeID, ageGroupID).Find(&plan)
	if result.Error != nil {
		fmt.Printf("Error fetching exercise plan: %v", result.Error)
	}

	fmt.Printf("Cardio: %d, Workout: %d\n", plan.Cardio, plan.Workout)
	cardio, workout := GetCardioWorkoutPlan(m.db, plan.Cardio, plan.Workout)
	fmt.Println(workout, cardio)
	return workout, cardio

}

// Close method to close the database connection
func (m *Mydb) Close() error {
	sqlDB, err := m.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func GetCardioWorkoutPlan(db *gorm.DB, cardio int, workout int) ([]models.Weekday, []models.Weekday) {
	var cardioSchedules []models.Weekday
	var workoutSchedules []models.Weekday
	// var cardioSchedule interface{}
	// var workoutSchedule interface{}
	// currentDay := time.Now().Weekday()
	dayColumnName := variables.DayColumnName
	// dayColumnName := "Sunday"

	switch {
	case cardio == 1:
		// var schedule WorkoutSchedule1
		// result := db.Find(&schedule)
		result := db.Model(&WorkoutSchedule1{}).Select(dayColumnName).Find(&cardioSchedules)
		if result.Error != nil {
			fmt.Printf("Error fetching WorkoutSchedule1: %v", result.Error)
		}
		// cardioSchedule = schedule
	case cardio == 2:
		// var schedule WorkoutSchedule2
		// result := db.Find(&schedule)
		result := db.Model(&WorkoutSchedule2{}).Select(dayColumnName).Find(&cardioSchedules)
		if result.Error != nil {
			fmt.Printf("Error fetching WorkoutSchedule2: %v", result.Error)
		}
		// cardioSchedule = schedule
	case cardio == 3:
		// var schedule WorkoutSchedule3
		result := db.Model(&WorkoutSchedule3{}).Select(dayColumnName).Find(&cardioSchedules)
		// result := db.Find(&schedule)
		if result.Error != nil {
			fmt.Printf("Error fetching WorkoutSchedule3: %v", result.Error)
		}
		// cardioSchedule = schedule
	default:
		log.Println("No matching workout schedule found for the given cardio and workout values.")
	}
	switch {
	case workout == 1:
		// var schedule StrengthSchedule1
		result := db.Model(&StrengthSchedule1{}).Select(dayColumnName).Find(&workoutSchedules)
		// result := db.Find(&schedule)
		if result.Error != nil {
			fmt.Printf("Error fetching WorkoutSchedule1: %v", result.Error)
		}
		// workoutSchedule = schedule
	case workout == 2:
		// var schedule StrengthSchedule2
		result := db.Model(&StrengthSchedule2{}).Select(dayColumnName).Find(&workoutSchedules)
		// result := db.Find(&schedule)
		if result.Error != nil {
			fmt.Printf("Error fetching WorkoutSchedule2: %v", result.Error)
		}
		// workoutSchedule = schedule
	case workout == 3:
		// var schedule StrengthSchedule3
		result := db.Model(&StrengthSchedule3{}).Select(dayColumnName).Find(&workoutSchedules)
		// result := db.Find(&schedule)
		if result.Error != nil {
			fmt.Printf("Error fetching WorkoutSchedule3: %v", result.Error)
		}
		// workoutSchedule = schedule
	default:
		log.Println("No matching workout schedule found for the given cardio and workout values.")
	}
	return cardioSchedules, workoutSchedules

}

func (m *Mydb) GetReps(ageGroupID int, excerciseType string) (string, string) {

	// var result DeadliftResult
	var deadlift string
	if err := m.db.Table(Exercise_table).Select(excerciseType).Where("age = ?", ageGroupID).Scan(&deadlift).Error; err != nil {
		fmt.Println("Error querying record:", err)
	} else {
		fmt.Printf("Deadlift for Age 2: %d\n", deadlift)
	}
	return deadlift, excerciseType
}

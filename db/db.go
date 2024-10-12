package db

import (
	"encoding/json"
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

const (
	Exercise_table     = "exercise_data"
	Workoutplans_table = "workoutplans"
)

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

	err = db.AutoMigrate(&DietPlan1{})
	if err != nil {
		fmt.Printf("Error migrating the Diet table in database: %v", err)
	}
	err = db.AutoMigrate(&DietPlan2{})
	if err != nil {
		fmt.Printf("Error migrating the Diet table in database: %v", err)
	}
	err = db.AutoMigrate(&DietPlan3{})
	if err != nil {
		fmt.Printf("Error migrating the Diet table in database: %v", err)
	}
	for _, diet := range DietPlans1 {
		db.Create(&diet)
	}
	for _, diet := range DietPlans2 {
		db.Create(&diet)
	}
	for _, diet := range DietPlans3 {
		db.Create(&diet)
	}
	err = db.AutoMigrate(&models.Workoutplan{})
	if err != nil {
		fmt.Printf("Error migrating the Workout table in database: %v", err)
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
	fmt.Println("Reached GetUser for ", username)

	// Use First to get the first matching row (if exists)
	result := m.db.First(&user, "username = ?", username)
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

func (m *Mydb) GetUserDietPlanfromDB(bodyTypeID int, ageGroupID int) models.DietPlan {
	// var plan Workout
	// bodyTypeID := 1 // Example BodyTypeID
	// ageGroupID := 2 // Example AgeGroupID

	// result := m.db.Where("body_type_id = ? AND age_group_id = ?", bodyTypeID, ageGroupID).Find(&plan)
	// if result.Error != nil {
	// 	fmt.Printf("Error fetching exercise plan: %v", result.Error)
	// }

	// fmt.Printf("Cardio: %d, Workout: %d\n", plan.Cardio, plan.Workout)
	diet := GetDietPlan(m.db, ageGroupID)
	// fmt.Println(workout, cardio)
	return diet

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

func GetDietPlan(db *gorm.DB, diet int) models.DietPlan {
	// var cardioSchedules []models.Weekday
	// var workoutSchedules []models.Weekday
	var dietplan models.DietPlan
	// var cardioSchedule interface{}
	// var workoutSchedule interface{}
	// currentDay := time.Now().Weekday()
	dayColumnName := variables.DayColumnName
	// dayColumnName := "Sunday"

	switch {
	case diet == 1:
		// var schedule WorkoutSchedule1
		// result := db.Find(&schedule)
		// var dietPlan DietPlan1
		if err := db.Table("diet_plan1").Where("day = ?", dayColumnName).First(&dietplan).Error; err != nil {
			log.Println("Error fetching diet plan:", err)
		} else {
			fmt.Printf("Diet Plan for %s:\n", dayColumnName)
			fmt.Printf("Breakfast: %s\n", dietplan.Breakfast)
			fmt.Printf("Lunch: %s\n", dietplan.Lunch)
			fmt.Printf("Dinner: %s\n", dietplan.Dinner)
		}

		// cardioSchedule = schedule
	case diet == 2:
		// var schedule WorkoutSchedule2
		// result := db.Find(&schedule)
		// var dietPlan DietPlan2
		if err := db.Table("diet_plan2").Where("day = ?", dayColumnName).First(&dietplan).Error; err != nil {
			log.Println("Error fetching diet plan:", err)
		} else {
			fmt.Printf("Diet Plan for %s:\n", dayColumnName)
			fmt.Printf("Breakfast: %s\n", dietplan.Breakfast)
			fmt.Printf("Lunch: %s\n", dietplan.Lunch)
			fmt.Printf("Dinner: %s\n", dietplan.Dinner)
		}
		// cardioSchedule = schedule
	case diet == 3:
		// var schedule WorkoutSchedule3
		// var dietPlan DietPlan3
		if err := db.Table("diet_plan3").Where("day = ?", dayColumnName).First(&dietplan).Error; err != nil {
			log.Println("Error fetching diet plan:", err)
		} else {
			fmt.Printf("Diet Plan for %s:\n", dayColumnName)
			fmt.Printf("Breakfast: %s\n", dietplan.Breakfast)
			fmt.Printf("Lunch: %s\n", dietplan.Lunch)
			fmt.Printf("Dinner: %s\n", dietplan.Dinner)
		}
		// cardioSchedule = schedule
	default:
		log.Println("No matching workout schedule found for the given cardio and workout values.")
	}
	fmt.Println(dietplan)

	return dietplan

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

func (m *Mydb) WriteTarget2DB(username string, workoutplan *models.Workoutplan) error {
	var existingWorkout models.Workoutplan
	today := variables.Today

	// Set workout plan values
	workoutplan.Username = username
	workoutplan.Date = variables.Today

	// Check if a record with the same username and date already exists
	err := m.db.Where("username = ? AND date = ?", username, today).First(&existingWorkout).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		// An error occurred other than not finding the record
		fmt.Println("Error fetching record:", err)
		return err
	}

	if err == gorm.ErrRecordNotFound {
		// Record does not exist, create a new one
		if err := m.db.Create(&workoutplan).Error; err != nil {
			fmt.Println("Error creating new record:", err)
			return err
		} else {
			fmt.Println("New record created successfully")
		}
	} else {
		// Record exists, update it with the new data
		if err := m.db.Model(&existingWorkout).Updates(workoutplan).Error; err != nil {
			fmt.Println("Error updating existing record:", err)
			return err
		} else {
			fmt.Println("Record updated successfully")
		}
	}
	return nil
}

func (m *Mydb) GetUserWorkoutDetails4mDB(user string, date string) models.Workoutplan {
	var details models.Workoutplan
	if err := m.db.Table(Workoutplans_table).Select("*").Where("username = ?", user).Where("date = ?", date).Scan(&details).Error; err != nil {
		fmt.Println("Error querying record:", err)
	} else {
		fmt.Printf("Recieved details  %d\n", details)
	}
	return details
}

func (m *Mydb) GetUserWorkoutDetails42day4mDB(user string, cardio []byte, workout []byte, date string) models.Workouttodaylist {
	var cardiomap map[string]interface{}
	var workoutmap map[string]interface{}
	workouttoday := []models.Workouttoday{}

	// Unmarshal the JSON byte slices into maps
	err1 := json.Unmarshal(cardio, &cardiomap)
	if err1 != nil {
		fmt.Println("Error unmarshalling cardio data:", err1)
		return models.Workouttodaylist{} // Return an empty result on error
	}
	fmt.Println("cardiomap:", cardiomap)

	err2 := json.Unmarshal(workout, &workoutmap)
	if err2 != nil {
		fmt.Println("Error unmarshalling workout data:", err2)
		return models.Workouttodaylist{} // Return an empty result on error
	}
	fmt.Println("workoutmap:", workoutmap)

	// Iterate through the cardio map to populate workouttoday
	for key, value := range cardiomap {
		var details int
		if err := m.db.Table(Workoutplans_table).Select(key+"done").Where("username = ?", user).Where("date = ?", date).Find(&details).Error; err != nil {
			fmt.Println("Error querying record:", err)
		} else {
			fmt.Printf("Received details for cardio: %d\n", details)
		}

		// Create a Workouttoday entry
		target, ok := value.(string)
		if !ok {
			fmt.Printf("Error: target for key %s is not a string\n", key)
			continue
		}

		workoutEntry := models.Workouttoday{
			Name:   key,
			Target: target,
			Done:   details,
		}

		workouttoday = append(workouttoday, workoutEntry)
	}

	// Iterate through the workout map to populate workouttoday
	for key, value := range workoutmap {
		var details int
		if err := m.db.Table(Workoutplans_table).Select(key+"done").Where("username = ?", user).Where("date = ?", date).Find(&details).Error; err != nil {
			fmt.Println("Error querying record:", err)
			return models.Workouttodaylist{}
		} else {
			fmt.Printf("Received details for workout: %d\n", details)
		}

		// Create a Workouttoday entry
		target, ok := value.(string)
		if !ok {
			fmt.Printf("Error: target for key %s is not a string\n", key)
			continue
		}

		workoutEntry := models.Workouttoday{
			Name:   key,
			Target: target,
			Done:   details,
		}

		workouttoday = append(workouttoday, workoutEntry)
	}

	// Construct the final Workouttodaylist
	workoutList := models.Workouttodaylist{
		Workout: workouttoday, // Assuming Workouttodaylist has a Workout field
	}

	fmt.Println("Final content of the list:")
	fmt.Println(workoutList)
	return workoutList
}
func (m *Mydb) CheckUserDateComboExistinDB(user string, date string) error {
	var count int64

	err := m.db.Table(Workoutplans_table).
		Where("username = ?", user).
		Where("date = ?", date).
		Count(&count).Error

	if err != nil {
		fmt.Println("Error querying the database:", err)
		return err // Handle the error appropriately
	}

	if count > 0 {
		fmt.Println("Record exists for the specified username and date.")
		return nil
	} else {
		fmt.Println("No record found for the specified username and date.")
		return fmt.Errorf("Record not found")
	}
	return nil
}

// var workoutable models.Workoutplan
// // today := time.Now().Truncate(24 * time.Hour)
// now := time.Now()
// today := now.Format("2006-01-02")

// workoutplan.Username = user.Username
// workoutplan.Date = today
// fmt.Printf("%+v", workoutplan)
// if err := m.db.Where("username = ? AND date = ?", user.Username, today).First(&workoutable).Error; err != nil {
// 	fmt.Println("theisoeio")
// 	if err == gorm.ErrRecordNotFound {
// 		// Record does not exist, create a new one
// 		workoutplan.Username = user.Username
// 		workoutplan.Date = today
// 	}
// 	if err := m.db.Create(&workoutplan).Error; err != nil {
// 		fmt.Println(err)
// 	}
// } else {
// 	fmt.Println(err)
// 	fmt.Println("record found")
// }
// fmt.Printf("%+v", workoutplan)
// if err := m.db.Where("username = ? AND date = ?", user.Username, today).Save(&workoutplan).Error; err != nil {
// 	fmt.Println(err)
// }
// }

// m.db.Save(workoutplan)
// }

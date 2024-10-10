package handler

import (
	"encoding/json"
	"fmt"
	"jwt-auth/mocks"
	"jwt-auth/models"
	"jwt-auth/utils"
	"jwt-auth/variables"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

// Login handler
func Login(db models.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("reached login")
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, ok := db.GetUser(user.Username)
		// pass_val, ok :=
		if ok.Password != user.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		expirationTime4Token := time.Now().Add(time.Hour * 5)
		claims := utils.Claims{
			Username: user.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime4Token.Unix(),
			},
		}
		newtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := newtoken.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Send the token as a JSON response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"token": tokenString,
		})
		WriteTarget2UserDB(db, user)
	}
}

// Refresh handler
func Refresh(w http.ResponseWriter, r *http.Request) {
	// Expect the token from Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims := &utils.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create a new token with a refreshed expiration time
	expirationTime4Token := time.Now().Add(time.Minute * 5)
	claims.ExpiresAt = expirationTime4Token.Unix()
	newtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = newtoken.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the refreshed token in response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

// Signup handler
func Signup(db models.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("reached signup")
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// _, ok := utils.Struct2Map()[user.Username]
		ok, out := db.GetUser(user.Username)
		fmt.Println(out)
		if ok {
			w.WriteHeader(http.StatusConflict)
			fmt.Println(user.Username)
			w.Write([]byte("User already exists, please login"))
			return
		}
		db.AddUser(user)
		mocks.Users = append(mocks.Users, user)
		w.Write([]byte("User profile created successfully"))
	}
}

// CheckAuth handler
func CheckAuth(w http.ResponseWriter, r *http.Request) {
	// Expect the token from Authorization header
	// authHeader := r.Header.Get("Authorization")
	// if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }
	// tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// claims := &utils.Claims{}
	// tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
	// 	return jwtKey, nil
	// })
	// if err != nil || !tkn.Valid {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	w.Write([]byte("Hello"))
}

// Home handler
func Home(w http.ResponseWriter, r *http.Request) {
	// Expect the token from Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims := &utils.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))
}

func GetUserById(db models.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		// fmt.Println("reached GetUserById")
		// authHeader := r.Header.Get("Authorization")
		// if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }
		// tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// claims := &utils.Claims{}
		// tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		// 	return jwtKey, nil
		// })
		// if err != nil || !tkn.Valid {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		path := r.URL.Path
		// Split the path into parts
		parts := strings.Split(path, "/")

		var id string
		// Check if the correct number of parts is present
		if len(parts) == 3 {
			id = parts[2] // "123" (the third part)
			// fmt.Fprintf(w, "User ID: %s", id)
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		}
		// _, ok := utils.Struct2Map()[user.Username]
		ok, User := db.GetUser(id)
		if !ok {
			// w.WriteHeader(http.StatusConflict)

			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("User doesnot already exists"))
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Println("##")
		fmt.Println(User)
		json.NewEncoder(w).Encode(User)
		// db.AddUser(user)
		// mocks.Users = append(mocks.Users, user)
		// w.Write([]byte("User profile created successfully"))
	}
}

func GetUserWorkOutPlan(db models.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		// fmt.Println("reached GetUserById")
		// authHeader := r.Header.Get("Authorization")
		// if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }
		// tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// claims := &utils.Claims{}
		// tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		// 	return jwtKey, nil
		// })
		// if err != nil || !tkn.Valid {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		path := r.URL.Path
		// Split the path into parts
		parts := strings.Split(path, "/")

		var id string
		// Check if the correct number of parts is present
		if len(parts) == 4 {
			id = parts[2] // "123" (the third part)
			// fmt.Fprintf(w, "User ID: %s", id)
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		}
		// _, ok := utils.Struct2Map()[user.Username]
		ok, User := db.GetUser(id)
		if !ok {
			// w.WriteHeader(http.StatusConflict)

			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("User doesnot already exists"))
			return
		}
		json_workout, _ := CreateWorkoutCardioResponse(db, User)
		// bmiID, ageID := LinkAgeBMIid(User.BMI, User.Age)
		// fmt.Println(bmiID, ageID)
		// workout, _ := db.GetUserWorkOutCardioPlanfromDB(bmiID, ageID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Println("##")
		fmt.Println(User)
		// json_cardio := FormatCardio(db, ageID, workout)
		// for _, j := range cardio {
		// 	fmt.Println("Entereed ehre")
		// 	fmt.Println(j)
		// }
		w.Header().Set("Content-Type", "application/json")
		w.Write(json_workout)
		// fmt.Println(json_cardio)

		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// fmt.Println("##")
		// fmt.Println(User)
		// json.NewEncoder(w).Encode(workout)
		// db.AddUser(user)
		// mocks.Users = append(mocks.Users, user)
		// w.Write([]byte("User profile created successfully"))
	}
}

func GetUserCardioPlan(db models.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		// fmt.Println("reached GetUserById")
		// authHeader := r.Header.Get("Authorization")
		// if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }
		// tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// claims := &utils.Claims{}
		// tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		// 	return jwtKey, nil
		// })
		// if err != nil || !tkn.Valid {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		path := r.URL.Path
		// Split the path into parts
		parts := strings.Split(path, "/")

		var id string
		// Check if the correct number of parts is present
		if len(parts) == 4 {
			id = parts[2] // "123" (the third part)
			// fmt.Fprintf(w, "User ID: %s", id)
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		}
		// _, ok := utils.Struct2Map()[user.Username]
		ok, User := db.GetUser(id)
		if !ok {
			// w.WriteHeader(http.StatusConflict)

			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("User doesnot already exists"))
			return
		}
		// type Weekday struct {
		// 	Monday    string
		// 	Tuesday   string
		// 	Wednesday string
		// 	Thursday  string
		// 	Friday    string
		// 	Saturday  string
		// 	Sunday    string
		// }
		// Today = "Sunday"
		_, json_cardio := CreateWorkoutCardioResponse(db, User)
		// bmiID, ageID := LinkAgeBMIid(User.BMI, User.Age)
		// fmt.Println(bmiID, ageID)
		// workout, _ := db.GetUserWorkOutCardioPlanfromDB(bmiID, ageID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Println("##")
		fmt.Println(User)
		// json_cardio := FormatCardio(db, ageID, workout)
		// for _, j := range cardio {
		// 	fmt.Println("Entereed ehre")
		// 	fmt.Println(j)
		// }
		w.Header().Set("Content-Type", "application/json")
		w.Write(json_cardio)

		// bmiID, ageID := LinkAgeBMIid(User.BMI, User.Age)
		// fmt.Println(bmiID, ageID)
		// _, cardio := db.GetUserWorkOutCardioPlanfromDB(bmiID, ageID)
		// // cardioData, _ := cardio.()
		// // fmt.Println(cardioData)
		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// fmt.Println("##")
		// fmt.Println(User)
		// json_cardio := FormatCardio(db, ageID, cardio)
		// // for _, j := range cardio {
		// // 	fmt.Println("Entereed ehre")
		// // 	fmt.Println(j)
		// // }
		// w.Header().Set("Content-Type", "application/json")
		// w.Write(json_cardio)
		// fmt.Println(json_cardio)

		// json.NewEncoder(w).Encode(string(json_cardio))
		// w.Header().Set("Content-Type", "application/json")
		// db.AddUser(user)
		// mocks.Users = append(mocks.Users, user)
		// w.Write([]byte("User profile created successfully"))
	}
}

func GetUserDietPlan(db models.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		// fmt.Println("reached GetUserById")
		// authHeader := r.Header.Get("Authorization")
		// if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }
		// tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// claims := &utils.Claims{}
		// tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		// 	return jwtKey, nil
		// })
		// if err != nil || !tkn.Valid {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		path := r.URL.Path
		// Split the path into parts
		parts := strings.Split(path, "/")

		var id string
		// Check if the correct number of parts is present
		if len(parts) == 4 {
			id = parts[2] // "123" (the third part)
			// fmt.Fprintf(w, "User ID: %s", id)
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		}
		// _, ok := utils.Struct2Map()[user.Username]
		ok, User := db.GetUser(id)
		if !ok {
			// w.WriteHeader(http.StatusConflict)

			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("User doesnot already exists"))
			return
		}
		// type Weekday struct {
		// 	Monday    string
		// 	Tuesday   string
		// 	Wednesday string
		// 	Thursday  string
		// 	Friday    string
		// 	Saturday  string
		// 	Sunday    string
		// }
		// Today = "Sunday"

		bmiID, ageID := LinkAgeBMIid(User.BMI, User.Age)
		fmt.Println(bmiID, ageID)
		diet := db.GetUserDietPlanfromDB(bmiID, ageID)
		// cardioData, _ := cardio.()
		// fmt.Println(cardioData)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Println("##")
		fmt.Println(diet)
		b, err := json.Marshal(diet)
		// json_cardio := FormatCardio(db, ageID, cardio)
		// for _, j := range cardio {
		// 	fmt.Println("Entereed ehre")
		// 	fmt.Println(j)
		// }
		// json.NewEncoder(w).Encode(string(diet))
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		// fmt.Println(json_cardio)

		// json.NewEncoder(w).Encode(string(json_cardio))
		// w.Header().Set("Content-Type", "application/json")
		// db.AddUser(user)
		// mocks.Users = append(mocks.Users, user)
		// w.Write([]byte("User profile created successfully"))
	}
}

func GetUserInput(db models.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		// fmt.Println("reached GetUserById")
		// authHeader := r.Header.Get("Authorization")
		// if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }
		// tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// claims := &utils.Claims{}
		// tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		// 	return jwtKey, nil
		// })
		// if err != nil || !tkn.Valid {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }
		var userinput models.Workoutplan
		fmt.Println("Recieved user input is ", r.Body)
		err := json.NewDecoder(r.Body).Decode(&userinput)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		path := r.URL.Path
		// Split the path into parts
		parts := strings.Split(path, "/")

		var id string
		// Check if the correct number of parts is present
		if len(parts) == 3 {
			id = parts[2] // "123" (the third part)
			// fmt.Fprintf(w, "User ID: %s", id)
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		}
		err = db.WriteTarget2DB(id, &userinput)
		// _, ok := utils.Struct2Map()[user.Username]
		// ok, User := db.GetUser(id)
		if err != nil {
			// w.WriteHeader(http.StatusConflict)

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Println("##")
		// fmt.Println(User)
		// json.NewEncoder(w).Encode(User)
		// db.AddUser(user)
		// mocks.Users = append(mocks.Users, user)
		w.Write([]byte("Entry updated successfully"))
	}
}

func LinkAgeBMIid(bmi, age int64) (int, int) {
	// Body type 1 Underweight = <18.5
	// Body type 2 Normal weight = 18.5–24.9
	// Body type 3 Overweight = 25–29.9
	var bmiID, ageID int
	if bmi <= 18 {
		bmiID = 1
	} else if 18 < bmi && bmi <= 25 {
		bmiID = 2
	} else {
		bmiID = 3
	}
	// Group 1 is Age less than 20
	// Group 2 is Age between 21 and  40
	// Group 3 is Age above 41 a
	if age <= 20 {
		ageID = 1
	} else if 21 < age && age <= 40 {
		ageID = 2
	} else {
		ageID = 3
	}
	return bmiID, ageID

}

func FormatCardio(db models.Database, ageID int, cardio []models.Weekday) []byte {
	dayColumnName := variables.DayColumnName
	data := map[string]interface{}{}
	for _, j := range cardio {
		fmt.Println("Entereed ehre")
		fmt.Println(j)
		// var test models.CardioList
		switch dayColumnName {
		case "Monday":
			reps, item := db.GetReps(ageID, j.Monday)

			data[item] = reps
			// var schedule WorkoutSchedule1
			// result := db.Find(&schedule)

			// cardioSchedule = schedule
		case "Tuesday":
			reps, item := db.GetReps(ageID, j.Tuesday)
			// var schedule WorkoutSchedule2
			// result := db.Find(&schedule)
			data[item] = reps
			// cardioSchedule = schedule
		case "Wednesday":
			reps, item := db.GetReps(ageID, j.Wednesday)
			// var schedule WorkoutSchedule3
			data[item] = reps
			// cardioSchedule = schedule
		case "Thursday":
			reps, item := db.GetReps(ageID, j.Thursday)
			// var schedule WorkoutSchedule3
			data[item] = reps
		case "Friday":
			reps, item := db.GetReps(ageID, j.Friday)
			// var schedule WorkoutSchedule3
			data[item] = reps
		case "Saturday":
			reps, item := db.GetReps(ageID, j.Saturday)
			// var schedule WorkoutSchedule3
			data[item] = reps
		case "Sunday":
			reps, item := db.GetReps(ageID, j.Sunday)
			// var schedule WorkoutSchedule3
			data[item] = reps
		default:
			log.Println("No matching workout schedule found for the given cardio and workout values.")
		}
	}
	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// prettyJSON, err := json.MarshalIndent(data, "", "  ")
	// if err != nil {
	// 	fmt.Println("Error marshalling data for pretty print:", err)

	// }

	// // Print the pretty JSON string
	// fmt.Println("Pretty JSON Output:")
	// fmt.Println(string(prettyJSON))

	// return string(prettyJSON)
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling data for pretty print:", err)
		// return
	}
	return jsonData

}

func CreateWorkoutCardioResponse(db models.Database, User models.User) ([]byte, []byte) {

	bmiID, ageID := LinkAgeBMIid(User.BMI, User.Age)
	fmt.Println(bmiID, ageID)
	workout, cardio := db.GetUserWorkOutCardioPlanfromDB(bmiID, ageID)
	json_cardio := FormatCardio(db, ageID, cardio)
	json_workout := FormatCardio(db, ageID, workout)
	return json_workout, json_cardio
}

func WriteTarget2UserDB(db1 models.Database, user models.User) {
	workout, cardio := CreateWorkoutCardioResponse(db1, user)

	var totalworkoutCardio models.Workoutplan
	fmt.Printf("%+v", totalworkoutCardio)

	json.Unmarshal(workout, &totalworkoutCardio)
	json.Unmarshal(cardio, &totalworkoutCardio)
	fmt.Println(string(workout))
	fmt.Println(string(cardio))
	// TotalworkoutCardio.Date = time.Now().Truncate(24 * time.Hour)
	// TotalworkoutCardio.Username = user.Username
	fmt.Printf("%+v", totalworkoutCardio)
	fmt.Println("#$#$#$#")
	db1.WriteTarget2DB(user.Username, &totalworkoutCardio)
}

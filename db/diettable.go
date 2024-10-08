package db

// type DietPlan struct {
// 	// gorm.Model
// 	// Day       string `json:"day,omitempty"`       // Day of the week (Monday, Tuesday, etc.)
// 	Breakfast string `json:"breakfast,omitempty"` // Breakfast meal
// 	Lunch     string `json:"lunch,omitempty"`     // Lunch meal
// 	Dinner    string `json:"dinner,omitempty"`    // Dinner meal
// }

type DietPlan1 struct {
	// gorm.Model
	Day       string `gorm:"not null"` // Day of the week (Monday, Tuesday, etc.)
	Breakfast string `gorm:"size:255"` // Breakfast meal
	Lunch     string `gorm:"size:255"` // Lunch meal
	Dinner    string `gorm:"size:255"` // Dinner meal
}

func (DietPlan1) TableName() string {
	return "diet_plan1"
}

var DietPlans1 = []DietPlan1{
	{Day: "Monday", Breakfast: "Oats Porridge with Fruits", Lunch: "Grilled Chicken Salad", Dinner: "Dal with Mixed Vegetables"},
	{Day: "Tuesday", Breakfast: "Egg White Omelette with Spinach", Lunch: "Quinoa with Steamed Vegetables", Dinner: "Fish Curry (Less Oil) with Brown Rice"},
	{Day: "Wednesday", Breakfast: "Sprouted Moong Salad", Lunch: "Cabbage Sabzi with Roti", Dinner: "Tofu Stir-fry (Minimal Oil)"},
	{Day: "Thursday", Breakfast: "Smoothie", Lunch: "Chickpea and Cucumber Salad", Dinner: "Palak Soup with Whole Wheat Toast"},
	{Day: "Friday", Breakfast: "Vegetable Dalia", Lunch: "Grilled Paneer with Steamed Broccoli", Dinner: "Lentil Soup with Spinach"},
	{Day: "Saturday", Breakfast: "Besan Chilla", Lunch: "Stuffed Capsicum", Dinner: "Methi Chicken (Less Oil) with Brown Rice"},
	{Day: "Sunday", Breakfast: "Cucumber and Tomato Salad", Lunch: "Vegetable Khichdi", Dinner: "Zucchini Noodles with Tomato Sauce"},
}

type DietPlan2 struct {
	// gorm.Model
	Day       string `gorm:"not null"` // Day of the week (Monday, Tuesday, etc.)
	Breakfast string `gorm:"size:255"` // Breakfast meal
	Lunch     string `gorm:"size:255"` // Lunch meal
	Dinner    string `gorm:"size:255"` // Dinner meal
}

func (DietPlan2) TableName() string {
	return "diet_plan2"
}

var DietPlans2 = []DietPlan2{
	{Day: "Monday", Breakfast: "Oats Porridge with Fruits", Lunch: "Grilled Chicken Salad", Dinner: "Dal with Mixed Vegetables"},
	{Day: "Tuesday", Breakfast: "Egg White Omelette with Spinach", Lunch: "Quinoa with Steamed Vegetables", Dinner: "Fish Curry (Less Oil) with Brown Rice"},
	{Day: "Wednesday", Breakfast: "Sprouted Moong Salad", Lunch: "Cabbage Sabzi with Roti", Dinner: "Tofu Stir-fry (Minimal Oil)"},
	{Day: "Thursday", Breakfast: "Smoothie", Lunch: "Chickpea and Cucumber Salad", Dinner: "Palak Soup with Whole Wheat Toast"},
	{Day: "Friday", Breakfast: "Vegetable Dalia", Lunch: "Grilled Paneer with Steamed Broccoli", Dinner: "Lentil Soup with Spinach"},
	{Day: "Saturday", Breakfast: "Besan Chilla", Lunch: "Stuffed Capsicum", Dinner: "Methi Chicken (Less Oil) with Brown Rice"},
	{Day: "Sunday", Breakfast: "Cucumber and Tomato Salad", Lunch: "Vegetable Khichdi", Dinner: "Zucchini Noodles with Tomato Sauce"},
}

type DietPlan3 struct {
	// gorm.Model
	Day       string `gorm:"not null"` // Day of the week (Monday, Tuesday, etc.)
	Breakfast string `gorm:"size:255"` // Breakfast meal
	Lunch     string `gorm:"size:255"` // Lunch meal
	Dinner    string `gorm:"size:255"` // Dinner meal
}

func (DietPlan3) TableName() string {
	return "diet_plan3"
}

var DietPlans3 = []DietPlan3{
	{Day: "Monday", Breakfast: "Oats Porridge with Fruits", Lunch: "Grilled Chicken Salad", Dinner: "Dal with Mixed Vegetables"},
	{Day: "Tuesday", Breakfast: "Egg White Omelette with Spinach", Lunch: "Quinoa with Steamed Vegetables", Dinner: "Fish Curry (Less Oil) with Brown Rice"},
	{Day: "Wednesday", Breakfast: "Sprouted Moong Salad", Lunch: "Cabbage Sabzi with Roti", Dinner: "Tofu Stir-fry (Minimal Oil)"},
	{Day: "Thursday", Breakfast: "Smoothie", Lunch: "Chickpea and Cucumber Salad", Dinner: "Palak Soup with Whole Wheat Toast"},
	{Day: "Friday", Breakfast: "Vegetable Dalia", Lunch: "Grilled Paneer with Steamed Broccoli", Dinner: "Lentil Soup with Spinach"},
	{Day: "Saturday", Breakfast: "Besan Chilla", Lunch: "Stuffed Capsicum", Dinner: "Methi Chicken (Less Oil) with Brown Rice"},
	{Day: "Sunday", Breakfast: "Cucumber and Tomato Salad", Lunch: "Vegetable Khichdi", Dinner: "Zucchini Noodles with Tomato Sauce"},
}

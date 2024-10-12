package variables

import "time"

var CurrentDay = time.Now().Weekday()

// var DayColumnName = CurrentDay.String()
var DayColumnName = "Wednesday"

var Today = time.Now().Format("2006-01-02")

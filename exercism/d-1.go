package exercism

import (
	"fmt"
	"strconv"
	"time"
)

func HasPassed(dateString string) bool {
	now := (time.Now())
	layout := "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	fmt.Println(t, now)
	return (t.UTC().Before(now.UTC()))
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(dateString string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, dateString)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	return 12 <= t.Hour() && t.Hour() <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// Input date string
	dateString := "7/25/2019 13:45:00"

	// Define the date layout for parsing
	layout := "1/2/2006 15:04:05"

	// Parse the date string
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		panic(err)
	}

	// Format the parsed time as required
	formattedTime := parsedTime.Format("Monday, January 2, 2006, at 15:04")
	return "You have an appointment on " + formattedTime
}

// AnniversaryDate returns a Time with this year's anniversary.
const OPENING_DATE = "September 15, 2012"

func AnniversaryDate() time.Time {
	now := (time.Now())
	openingDate, err := time.Parse("January 2, 2006", OPENING_DATE)
	if err != nil {
		panic(err)
	}
	result := openingDate.Month().String() + " " + strconv.Itoa(openingDate.Day()) + ", " + strconv.Itoa(now.Year())
	curDate, err := time.Parse("January 2, 2006", result)
	if err != nil {
		panic(err)
	}
	return curDate
}

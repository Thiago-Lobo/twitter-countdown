package logic

import (
	"time"
	"fmt"
)

func ValidateInputEventDate() {

}

func stringToTimestamp(dateString string) time.Time {
	
	result, err := time.Parse("2006-01-02 -0700", dateString)
	
	if err != nil {
		fmt.Println(err.Error())
		return time.Time{}
	}

	return result
}

func GetDaysToEvent(dateString string) int {
	eventDate := stringToTimestamp(fmt.Sprintf("%s -0300", dateString))
	
	currentDate := time.Now()
	currentDate = time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, currentDate.Location())
	
	return int(eventDate.Sub(currentDate).Hours() / 24)
}




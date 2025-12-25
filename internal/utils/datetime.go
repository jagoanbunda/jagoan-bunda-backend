package utils

import "time"

func CalculateAgeInMonths(birthday time.Time) int {
	now := time.Now()

	years := now.Year() - birthday.Year()
	months := int(now.Month()) - int(birthday.Month())

	if now.Day() < birthday.Day(){
		months--
	}

	totalMonths := years * 12 + months
	if totalMonths < 0 {
		return 0
	}

	return totalMonths
}

package moment_business

import (
	"time"
	"math"
)

func AddWeekDays(t time.Time, amount int) (time.Time) {

	if amount == 0 {
		return t
	}

	sign := 1;
	day := int(t.Weekday());
	absIncrement := int(math.Abs(float64(amount)));

	days := 0

	if day == 0 && sign == -1 {
		days = 1;
	} else if day == 6 && sign == 1 {
		days = 1;
	}

	// Add padding for weekends.
	paddedAbsIncrement := absIncrement;
	if day != 0 && day != 6 && sign > 0 {
		paddedAbsIncrement += day;
	} else if day != 0 && day != 6 && sign < 0 {
		paddedAbsIncrement += 6 - day;
	}

	p := 0
	if paddedAbsIncrement > 5 && paddedAbsIncrement % 5 > 0 {
		p = 1
	}

	weekendsInBetween := int(math.Max(math.Floor(float64(paddedAbsIncrement) / 5) - 1, 0.0)) + p

	// Add the increment and number of weekends.
	days += absIncrement + weekendsInBetween * 2;

	t = t.Add(time.Hour * 24 * time.Duration(days))

	return t
}
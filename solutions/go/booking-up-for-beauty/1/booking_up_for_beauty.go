package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layouts := []string{
        "1/2/2006 15:04:05",         // 7/13/2020 20:32:00
        "January 2, 2006 15:04:05",  // December 9, 2112 11:59:59
        "Monday, January 2, 2006 15:04:05",
    }

    for _, layout := range layouts {
        if t, err := time.Parse(layout, date); err == nil {
            return t
        }
    }
    return time.Time{}
}

func HasPassed(date string) bool {
    return time.Now().After(Schedule(date))
}


// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    t := Schedule(date)
    h := t.Hour()
    return h >= 12 && h < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    t := Schedule(date)
    h := t.Hour()
    mi := t.Minute()
    y := t.Year()
    mon := t.Month()
    day := t.Day()
    wd := t.Weekday()
    return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", wd, mon, day, y, h, mi )
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    year := time.Now().UTC().Year()
    return time.Date(year, time.September,15, 0, 0, 0, 0, time.UTC)
	panic("Please implement the AnniversaryDate function")
}

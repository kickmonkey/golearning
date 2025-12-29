package expenses
import "fmt"
// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	out := make([]Record, 0, len(in))
	for _, r := range in {
		if predicate(r) {
			out = append(out, r)
		}
	}
	return out
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
    return func(r Record) bool {
        return r.Day >= p.From && r.Day <= p.To
    }
	panic("Please implement the ByDaysPeriod function")
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(category string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == category
	}
}
// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(records []Record, p DaysPeriod) float64 {
	var total float64
	for _, r := range records {
		if r.Day >= p.From && r.Day <= p.To {
			total += r.Amount
		}
	}
	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(records []Record, p DaysPeriod, category string) (float64, error) {
	// 先判断 category 是否在“所有记录”里存在（不看 period）
	foundCategory := false
	for _, r := range records {
		if r.Category == category {
			foundCategory = true
			break
		}
	}
	if !foundCategory {
		return 0, fmt.Errorf("unknown category %s", category)
	}

	// category 存在：再计算 period 内该 category 的总和（可能为 0，但不报错）
	var total float64
	for _, r := range records {
		if r.Category == category && r.Day >= p.From && r.Day <= p.To {
			total += r.Amount
		}
	}
	return total, nil
}

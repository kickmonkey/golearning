package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return 3.213
	case balance < 1000:
		return 0.5
	case balance < 5000:
		return 1.621
	default:
		return 2.475
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    return float64(InterestRate(balance))*balance/100.0
	panic("Please implement the Interest function")
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
    return  Interest(balance) + balance
	panic("Please implement the AnnualBalanceUpdate function")
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	// already there
	if balance >= targetBalance {
		return 0
	}

	// If we're negative and target is non-negative, and negative balances only get more negative,
	// it's impossible to reach the target via interest updates alone.
	if balance < 0 && targetBalance >= 0 {
		return -1
	}

	years := 0
	for balance < targetBalance {
		next := AnnualBalanceUpdate(balance)

		// Safety: if balance doesn't increase toward target, break to avoid infinite loop
		if next <= balance {
			return -1
		}

		balance = next
		years++
	}
	return years
}

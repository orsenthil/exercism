package savings

// InterestRate calculates the interest rate for the provided balance.
func InterestRate(balance float64) float32 {

    switch {
        case balance < 0:
    		return float32(-0.03213 * 100)
        case balance >= 0 && balance < 1000:
    		return float32(0.005 * 100)
        case balance >= 1000 && balance < 5000:
    		return float32(0.01621 * 100)
        case balance >= 5000:
    		return float32(0.02475 * 100)
    }

	return float32(0.02475 * 100)
}

// InterestRate calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    var interest float32

    interest = InterestRate(balance)
    
    switch  {
        case balance < 0:
    		balance = -1 * (balance * float64(interest)) / float64(100)
        case balance >= 0 && balance < 1000:
    		balance = (balance * float64(interest)) / float64(100)
        case balance >= 1000 && balance < 5000:
    		balance = (balance * float64(interest)) / float64(100)
        case balance >= 5000:
    		balance = (balance * float64(interest)) / float64(100)
    }

	return balance
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
    var interest float64
    interest = Interest(balance)
    return balance + interest
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    var years int = 0
    for balance < targetBalance {
        balance = AnnualBalanceUpdate(balance)
        years += 1
    }

    return years
}

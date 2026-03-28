package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    unitsMap := make(map[string]int)
    unitsMap["quarter_of_a_dozen"] = 3
    unitsMap["half_of_a_dozen"] = 6
    unitsMap["dozen"] = 12
    unitsMap["small_gross"] = 120
    unitsMap["gross"] = 144
    unitsMap["great_gross"] = 1728

    return unitsMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    bill := make(map[string]int)
    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    value, exists := units[unit]
    if (exists == false) {
        return false
    }
	
	bill[item] += value
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    _, exists := bill[item]
    if (exists == false) {
        return false
    }
    value, exists := units[unit]
    if (exists == false) {
        return false
    }
	if ((bill[item] - value) < 0) {
        return false
    } else if ((bill[item] - value) == 0) {
    	delete(bill, item)
        return true
    }

    bill[item] -= value
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    value, exists := bill[item]
    return value, exists
}

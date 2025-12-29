package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	panic("Please implement the Units() function")
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    return map[string]int{}
	panic("Please implement the NewBill() function")
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    _, exists := units[unit]
    if !exists{
        return false
    } 
    bill[item] += units[unit]
    return true
	panic("Please implement the AddItem() function")
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    qty, ok := bill[item]
    if !ok{
        return false
    }
    u, ok := units[unit]
    if !ok{
        return false
    }
    if qty-u < 0{
        return false
    }
    qty -= u
    if qty == 0{
    	delete(bill, item)
    }else {
        bill[item] = qty
    }
    return true
    
	panic("Please implement the RemoveItem() function")
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    qty, ok := bill[item]
    if !ok{
        return 0, false
    }
    return qty, true
    
	panic("Please implement the GetItem() function")
}

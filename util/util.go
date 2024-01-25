package util

func CalculateEarnings(revenue float64, expenses float64, taxRate float64) (earningsBeforeTax, earningsAfterTax, ratio float64) {
	earningsBeforeTax = revenue - expenses
	earningsAfterTax = earningsBeforeTax * (taxRate / 100)
	ratio = earningsBeforeTax / earningsAfterTax
	return

}

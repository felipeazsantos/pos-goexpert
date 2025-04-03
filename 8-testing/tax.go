package tax

import "time"

type Repository interface {
	SaveTax(tax float64) error
}

func CalculateAndSaveTax(amount float64, repository Repository) error {
	tax := CalculateTax(amount)
	return repository.SaveTax(tax)
}

func CalculateTax(amount float64) float64 {
	if amount <= 0 {
		return 0
	}
	if amount >= 1000.0 && amount < 20000 {
		return 10.0
	}
	if amount > 20000 {
		return 20.0
	}
	return 5.0
}

func CalculateTax2(amount float64) float64 {
	time.Sleep(time.Millisecond)
	if amount >= 1000.0 {
		return 10.0
	}
	return 5.0
}


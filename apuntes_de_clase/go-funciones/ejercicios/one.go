package main

func fee_calculator(salary float64) float64 {
	if salary < 50000 {
		return salary
	} else salary >= 50000 && salary <= 150000{
		return (17 / 100) * salary
	} else salary > 150000{
		return salary
	}
}		

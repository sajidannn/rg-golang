package main

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name string
	BaseSalary int
	WorkingMonth int
}

type Senior struct {
	Name string
	BaseSalary int
	WorkingMonth int
	PerformanceRate float64
}

type Manager struct {
	Name string
	BaseSalary int
	WorkingMonth int
	PerformanceRate float64
	BonusManagerRate float64
}

func calculatePropWorkingMonth(workingMonth int) float64 {
	if workingMonth > 12 {
		return 1
	}
	return float64(workingMonth) / 12
}

func (j Junior) GetBonus() float64 {
	propWorkingMonth := calculatePropWorkingMonth(j.WorkingMonth)
	return float64(j.BaseSalary) * propWorkingMonth
}

func (s Senior) GetBonus() float64 {
	propWorkingMonth := calculatePropWorkingMonth(s.WorkingMonth)
	return 2 * float64(s.BaseSalary) * propWorkingMonth + (s.PerformanceRate * float64(s.BaseSalary))
}

func (m Manager) GetBonus() float64 {
	propWorkingMonth := calculatePropWorkingMonth(m.WorkingMonth)
	return 2 * float64(m.BaseSalary) * propWorkingMonth + (m.PerformanceRate * float64(m.BaseSalary)) + (m.BonusManagerRate * float64(m.BaseSalary))
}


func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus()
}

func TotalEmployeeBonus(employees []Employee) float64 {
	totalBonus := 0.0
	for _, employee := range employees {
		totalBonus += employee.GetBonus()
	}

	return totalBonus
}

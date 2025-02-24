package main

import "fmt"

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name         string
	BaseSalary   int
	WorkingMonth int
}

func (j Junior) GetBonus() float64 {
	return 1 * float64(j.BaseSalary) * float64(j.WorkingMonth) / 12
}

type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
}

func (s Senior) GetBonus() float64 {
	return 2*float64(s.BaseSalary)*float64(s.WorkingMonth)/12 + (s.PerformanceRate * float64(s.BaseSalary))
}

type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func (m Manager) GetBonus() float64 {
	return 2*float64(m.BaseSalary)*float64(m.WorkingMonth)/12 + (m.PerformanceRate * float64(m.BaseSalary)) + (m.BonusManagerRate * float64(m.BaseSalary))

}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus()
}

func TotalEmployeeBonus(employees []Employee) float64 {
	sum := 0.0
	for _, e := range employees {
		sum += e.GetBonus()
	}

	return sum
}

func main() {
	daffa := Junior{
		Name:         "Daffa",
		BaseSalary:   100000,
		WorkingMonth: 10,
	}

	afifi := Senior{
		Name:            "Afifi",
		BaseSalary:      200000,
		WorkingMonth:    20,
		PerformanceRate: 0.5,
	}

	syahrony := Manager{
		Name:             "Syahrony",
		BaseSalary:       300000,
		WorkingMonth:     2,
		PerformanceRate:  0.5,
		BonusManagerRate: 0.2,
	}

	var karyawan = []Employee{daffa, afifi, syahrony}

	fmt.Println(EmployeeBonus(daffa))
	fmt.Println(EmployeeBonus(afifi))
	fmt.Println(EmployeeBonus(syahrony))
	fmt.Println(TotalEmployeeBonus(karyawan))
}

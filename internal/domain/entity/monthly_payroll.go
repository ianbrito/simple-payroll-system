package entity

type MonthlyPayroll struct {
	ID        string
	Month     string
	Year      string
	Movements []*PayrollMovement
}

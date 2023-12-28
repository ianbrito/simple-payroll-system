package entity

type PayrollMovement struct {
	ID        string
	Movement  *PayrollMovement
	Lotation  *Lotation
	Events    []*Event
	Earnings  float64
	Discounts float64
	NetSalary float64
}

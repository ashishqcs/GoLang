package postgres

type BillEntity struct {
	Id       int64
	Name     string
	Address  string
	PlanName string `db:"plan_name"`
	Date     string
	Amount   float64
}

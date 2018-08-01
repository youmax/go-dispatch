package models

// Model Order's defination
type Order struct {
	ID          uint `gorm:"primary_key"`
	Custid      string
	Ordid       string
	CardID      uint
	Accountid   string
	Accountname string
	Bankcode    string
	Bankname    string
	Bgreturl    string
	Remark      string
	Amount      float64
	Balance     float64
	Charge      float64
	ChargeRate  float64
	AgentRate   float64
	Rakeback    float64
	Status      string
	Message     string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}

func (o *Order) Handle() {
	// the storageFolder method ensures that there are no name collision in
	// case we get same timestamp in the key name
	return
}

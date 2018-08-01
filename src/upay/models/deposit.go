package models

// Deposit model defination
type Deposit struct {
	ID          uint `gorm:"primary_key"`
	Type        string
	Custid      string
	CardID      uint
	Accountid   string
	Accountname string
	BankID      uint
	Remark      string
	Amount      float32
	Balance     float32
	Charge      float32
	ChargeRate  float32
	AgentRate   float32
	Rakeback    float32
	TransAt     string
	CreatedAt   string
	UpdatedAt   string
}

// Handle deposit payload
func (d *Deposit) Handle() {
	// the storageFolder method ensures that there are no name collision in
	// case we get same timestamp in the key name
	return
}

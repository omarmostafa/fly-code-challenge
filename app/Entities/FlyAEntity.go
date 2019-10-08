package Entities

type FlyPayAEntity struct {
	Transactions []FlyPayATransactions `json:"transactions"`
}

type FlyPayATransactions struct {
	TransactionId string `json:"transactionId"`
	Amount int `json:"amount"`
	StatusCode int `json:"statusCode"`
	OrderReference string `json:"orderReference"`
	Currency string `json:"currency"`
}

func (entity *FlyPayATransactions) GetStatusCode () string  {
	return statusCodeForFlyPayA()[entity.StatusCode]
}

func statusCodeForFlyPayA() map[int] string {
	return map[int]string{
		1 : "authorised" ,
		2 : "decline" ,
		3 : "refunded" ,
	}
}

package Entities

type FlyPayBEntity struct {
	Transactions []FlyPayBTransactions `json:"transactions"`
}

type FlyPayBTransactions struct {
	PaymentId string `json:"paymentId"`
	Value int `json:"value"`
	StatusCode int `json:"statusCode"`
	OrderInfo string `json:"orderInfo"`
	TransactionCurrency string `json:"transactionCurrency"`
}

func (entity *FlyPayBTransactions) GetStatusCode () string  {
	return statusCodeForFlyPayB()[entity.StatusCode]
}

func statusCodeForFlyPayB() map[int] string {
	return map[int]string{
		100 : "authorised" ,
		200 : "decline" ,
		300 : "refunded" ,
	}
}
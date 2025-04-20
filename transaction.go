package models

type Transaction struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Value     string `json:"value"`     
	Timestamp string `json:"timestamp"` 
	TxHash    string `json:"tx_hash"`
}

type TxDetail struct {
    TxAmount      float64 `json:"tx_amount"`
    DateTime      string  `json:"date_time"`
    TransactionID string  `json:"transaction_id"`
}

type Beneficiary struct {
    Address      string     `json:"beneficiary_address"`
    Amount       float64    `json:"amount"`
    Transactions []TxDetail `json:"transactions"`
}

type BeneficiaryResponse struct {
    Message string        `json:"message"`
    Data    []Beneficiary `json:"data"`
}

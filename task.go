package utils

import (
	"beneficiary-tracer/models"
	"fmt"
	"strconv"
	"time"
)

const weiToEther = 1e18

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

func GroupByBeneficiary(address string, txns []models.Transaction) ([]Beneficiary, error) {
	result := make(map[string]*Beneficiary)

	for _, tx := range txns {
		if tx.From == address {
			amount, err := strconv.ParseFloat(tx.Value, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse amount for transaction %s: %v", tx.TxHash, err)
			}

			timeInt, err := strconv.ParseInt(tx.Timestamp, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse timestamp for transaction %s: %v", tx.TxHash, err)
			}

			formattedTime := time.Unix(timeInt, 0).Format("2006-01-02 15:04:05")

			if _, exists := result[tx.To]; !exists {
				result[tx.To] = &Beneficiary{
					Address: tx.To,
				}
			}

			result[tx.To].Amount += amount / weiToEther
			result[tx.To].Transactions = append(result[tx.To].Transactions, TxDetail{
				TxAmount:      amount / weiToEther,
				DateTime:      formattedTime,
				TransactionID: tx.TxHash,
			})
		}
	}

	var grouped []Beneficiary
	for _, v := range result {
		grouped = append(grouped, *v)
	}
	return grouped, nil
}

func FetchNormalTransactions(address string) ([]models.Transaction, error) {
	var transactions []models.Transaction

	transactions = append(transactions, models.Transaction{
		From:      "8XDV7IIC6AF9EN5TQVHZTSRX648HN43I1",
		To:        address,
		Value:     "1000000000000000000",
		Timestamp: "1633072800",
		TxHash:    "0x1234567890abcdef",
	})

	return transactions, nil
}
//8XDV7IIC6AF9EN5TQVHZTSRX648HN43I1

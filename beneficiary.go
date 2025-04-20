package handlers

import (
	
	"net/http"

	"beneficiary-tracer/models"
	"beneficiary-tracer/utils"
	"github.com/gin-gonic/gin"
)
func GetBeneficiaries(c *gin.Context) {
	address := c.DefaultQuery("address", "") 
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing Ethereum address",
		})
		return
	}
	transactions, err := utils.FetchNormalTransactions(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	groupedBeneficiaries, err := utils.GroupByBeneficiary(address, transactions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := models.BeneficiaryResponse{
		Message: "success",
		Data:    convertToModelBeneficiaries(groupedBeneficiaries),
	}

	c.JSON(http.StatusOK, response)
}

func convertToModelBeneficiaries(input []utils.Beneficiary) []models.Beneficiary {
	var out []models.Beneficiary
	for _, b := range input {
		var txs []models.TxDetail
		for _, tx := range b.Transactions {
			txs = append(txs, models.TxDetail{
				TxAmount:      tx.TxAmount,
				DateTime:      tx.DateTime,
				TransactionID: tx.TransactionID,
			})
		}
		out = append(out, models.Beneficiary{
			Address:      b.Address,
			Amount:       b.Amount,
			Transactions: txs,
		})
	}
	return out
}

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type (
	transaction struct {
		gorm.Model
		Description   string `json:"description"`
		Amount        Price  `json:"amount"`
		Beneficiaries []user `json:"beneficiaries"`
	}
	transformedTransaction struct {
		ID            uint   `json:"id"`
		Description   string `json:"description"`
		Amount        Price  `json:"amount"`
		Beneficiaries []user `json:"beneficiaries"`
	}
)

func createTransaction(c *gin.Context) {
	description := c.PostForm("description")
	price, _ := ConvertToPrice(strconv.Atoi(c.PostForm("price")))
	// beneficiaries := c.PostForm("beneficiaries")
	beneficiaries := c.PostForm("beneficiaries")
	transaction := transaction{Description: description, Amount: price, Beneficiaries: beneficiaries}
	db.Save(&transaction)
	c.JSON(
		http.StatusCreated,
		gin.H{"status": http.StatusCreated, "message": "Transaction created successfully.", "resourceId": transaction.ID})
}

// func listTransaction(c *gin.Context) {
// 	var transactions []transaction
// 	var _transactions []transformedTransaction

// 	db.Find(&transactions)

// 	if len(transactions) <= 0 {
// 		c.JSON(
// 			http.StatusNotFound,
// 			gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
// 		return
// 	}

// 	for _, item := range transactions {
// 		_transactions = append(_transactions, transformedTransaction{ID: item.ID, Title: item.Title, Completed: completed})
// 	}
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _transactions})
// }

// func showTransaction(c *gin.Context) {
// 	var transaction transaction
// 	transactionID := c.Param("id")

// 	db.First(&transaction, transactionID)

// 	if transaction.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
// 		return
// 	}

// 	completed := false

// 	if transaction.Completed == 1 {
// 		completed = true
// 	} else {
// 		completed = false
// 	}

// 	_transaction := transformedTransaction{ID: transaction.ID, Title: transaction.Title, Completed: completed}

// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _Transaction})
// }

// func updateTransaction(c *gin.Context) {
// 	var transaction transaction
// 	transactionid := c.Param("id")

// 	db.First(&transaction, transactionID)

// 	if transaction.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
// 		return
// 	}

// 	db.Model(&transaction).Update("title", c.PostForm("title"))
// 	completed, _ := strconv.Atoi(c.PostForm("completed"))
// 	db.Model(&transaction).Update("completed", completed)
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Transaction updated successfully!"})
// }

// func deleteTransaction(c *gin.Context) {
// 	var transaction transaction
// 	transactionid := c.Param("id")
// 	db.First(&transaction, transactionID)
// 	if transaction.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Transaction found!"})
// 		return
// 	}
// 	db.Delete(&transaction)
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Transaction deleted successfully!"})
// }

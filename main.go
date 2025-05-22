package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequestDTO menggunakan map[string]interface{} agar fleksibel menerima data apapun
type RequestDTO struct {
	Data map[string]interface{} `json:"data"`
}

type RequestSignature struct {
	PartnerId string `json:"partnerId"`
	MerchantId string `json:"merchantId"`
	MerchantUser string `json:"merchantUser"`
	ReferenceNumber string `json:"referenceNumber"`
	Amount string `json:"amount"`
	ExpiryTime string `json:"expiryTime"`
}

// ResponseDTO untuk mengirim response dengan pesan dan data opsional
type ResponseDTO struct {
	QrValue string `json:"qrValue"`
	ResponseDescription string `json:"responseDescription"`
	ResponseCode string `json:"responseCode"`
	MerchantName string `json:"merchantName"`
	ExpiredDate string `json:"billNumber"`
}

type ResponseSignature struct {
	Data string `json:"data"`
}

// Validate memastikan data request tidak kosong
func (r *RequestDTO) Validate() bool {
	return r.Data != nil
}


// Handler function to process requests
func HandleData(c *gin.Context) {
	var req map[string]interface{}

	// Bind JSON to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request format",
			"error":   err.Error(),
		})
		return
	}

	// Validate request data
	// if !req.Validate() {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "Request data is required",
	// 	})
	// 	return
	// }

	// Create response
	resp := ResponseDTO{
		QrValue: "00020101021226650013ID.CO.BAG.WWW01189360003700000402680215ID20240004026880303UME51450015ID.OR.GPNQR.WWW0215ID20243564945810303UME5204653153033605405100005802ID5912NIRMANGSA A16007JAKARTA6105142506236012550000000000000000834780410703A016304A8DA",
		ResponseCode: "00",
		ResponseDescription: "Success",
		MerchantName: "EGR001",
		ExpiredDate: "2029-02-02",
	}

	// Send JSON response
	c.JSON(http.StatusOK, resp)
}

func SignatureDummy(c *gin.Context) {
	var request map[string]interface{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request format",
			"error":   err.Error(),
		})
		return
	}

	resp := "8JRGWF4MBWIT2pQZcOrt"


	// resp := ResponseSignature {
	// 	Data: "8JRGWF4MBWIT2pQZcOrt",
	// }
	c.String(http.StatusOK, resp)
}

func main() {
	router := gin.Default()
	fmt.Println("Server running on port 1414")

	// Register POST endpoint
	router.POST("/api/v3/qris/acquirer/generate", HandleData)
	router.POST("/api/generate", SignatureDummy)

	// Start server with error handling
	if err := router.Run(":1414"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

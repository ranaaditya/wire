/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Tag3600 is Business Function Code
type Tag3600 struct {
	// BusinessFunctionCode BTR: Bank Transfer (Beneficiary is a bank) DRC: Customer or Corporate Drawdown Request CKS: Check Same Day Settlement DRW: Drawdown Payment CTP: Customer Transfer Plus FFR: Fed Funds Returned CTR: Customer Transfer (Beneficiary is a not a bank) FFS: Fed Funds Sold DEP: Deposit to Sender’s Account SVC: Service Message DRB: Bank-to-Bank Drawdown Request 
	BusinessFunctionCode string `json:"businessFunctionCode"`
	// TransactionTypeCode If {3600} is CTR, an optional Transaction Type Code element is permitted; however, the Transaction Type Code 'COV' is not permitted. ToDo: Research This, I don't understand it yet
	TransactionTypeCode string `json:"transactionTypeCode,omitempty"`
}

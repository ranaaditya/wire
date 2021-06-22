/*
 * Wire API
 *
 * Moov Wire implements an HTTP API for creating, parsing, and validating Fedwire messages.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// BusinessFunctionCode struct for BusinessFunctionCode
type BusinessFunctionCode struct {
	// BusinessFunctionCode * `BTR` - Bank Transfer (Beneficiary is a bank) * `DRC` - Customer or Corporate Drawdown Request * `CKS` - Check Same Day Settlement * `DRW` - Drawdown Payment * `CTP` - Customer Transfer Plus * `FFR` - Fed Funds Returned * `CTR` - Customer Transfer (Beneficiary is a not a bank) * `FFS` - Fed Funds Sold * `DEP` - Deposit to Sender’s Account * `SVC` - Service Message * `DRB` - Bank-to-Bank Drawdown Request
	BusinessFunctionCode string `json:"businessFunctionCode"`
	// TransactionTypeCode If {3600} is CTR, an optional Transaction Type Code element is permitted; however, the Transaction Type Code 'COV' is not permitted.
	TransactionTypeCode string `json:"transactionTypeCode,omitempty"`
}

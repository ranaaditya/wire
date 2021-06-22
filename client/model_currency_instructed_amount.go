/*
 * Wire API
 *
 * Moov Wire implements an HTTP API for creating, parsing, and validating Fedwire messages.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CurrencyInstructedAmount struct for CurrencyInstructedAmount
type CurrencyInstructedAmount struct {
	// SwiftFieldTag
	SwiftFieldTag string `json:"swiftFieldTag,omitempty"`
	// Amount
	Amount string `json:"amount,omitempty"`
}

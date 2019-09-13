/*
 * Qc API
 *
 * Qc API
 *
 * API version: 3.0.0
 * Contact: cloudsupport@telestream.net
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package qc
type FieldOrderType string

// List of field_order_type
const (
	FIELDORDERTYPE_UNKNOWN_FIELD_ORDER FieldOrderType = "UnknownFieldOrder"
	FIELDORDERTYPE_TOP_FIELD_FIRST FieldOrderType = "TopFieldFirst"
	FIELDORDERTYPE_BOTTOM_FIELD_FIRST FieldOrderType = "BottomFieldFirst"
	FIELDORDERTYPE_PROGRESSIVE FieldOrderType = "Progressive"
	FIELDORDERTYPE_PULLDOWN FieldOrderType = "Pulldown"
	FIELDORDERTYPE_REPEATED FieldOrderType = "Repeated"
	FIELDORDERTYPE_CONSISTENT_FIELD_ORDER FieldOrderType = "ConsistentFieldOrder"
)
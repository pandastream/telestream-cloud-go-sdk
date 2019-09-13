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

type CorruptFrameTest struct {
	Sensitivity SensitivityType `json:"sensitivity,omitempty"`
	RejectOnError bool `json:"reject_on_error,omitempty"`
	DoCorrection bool `json:"do_correction,omitempty"`
	Checked bool `json:"checked,omitempty"`
}

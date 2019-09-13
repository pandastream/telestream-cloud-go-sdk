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

type BlackLevelTest struct {
	LevelDefaultOrCustom DefaultOrCustomType `json:"level_default_or_custom,omitempty"`
	Level int32 `json:"level,omitempty"`
	LevelMaxOutsideRange float32 `json:"level_max_outside_range,omitempty"`
	RejectOnError bool `json:"reject_on_error,omitempty"`
	DoCorrection bool `json:"do_correction,omitempty"`
	Checked bool `json:"checked,omitempty"`
}

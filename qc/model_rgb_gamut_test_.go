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

type RgbGamutTest struct {
	LevelDefaultOrCustom DefaultOrCustomType `json:"level_default_or_custom,omitempty"`
	LevelLower int32 `json:"level_lower,omitempty"`
	LevelUpper int32 `json:"level_upper,omitempty"`
	LevelMaxOutsideRange float32 `json:"level_max_outside_range,omitempty"`
	LevelTolerance float32 `json:"level_tolerance,omitempty"`
	LowPassFilter LowPassFilterType `json:"low_pass_filter,omitempty"`
	RejectOnError bool `json:"reject_on_error,omitempty"`
	DoCorrection bool `json:"do_correction,omitempty"`
	Checked bool `json:"checked,omitempty"`
}

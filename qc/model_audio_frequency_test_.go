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

type AudioFrequencyTest struct {
	ToneType ToneType `json:"tone_type,omitempty"`
	Frequency float64 `json:"frequency,omitempty"`
	TimeRangeEnabled bool `json:"time_range_enabled,omitempty"`
	Power float64 `json:"power,omitempty"`
	Tolerance float64 `json:"tolerance,omitempty"`
	StartTime float64 `json:"start_time,omitempty"`
	TimeSecsOrFrames SecsOrFramesType `json:"time_secs_or_frames,omitempty"`
	Duration float64 `json:"duration,omitempty"`
	DurationSecsOrFrames SecsOrFramesType `json:"duration_secs_or_frames,omitempty"`
	NotAtAnyOtherTime bool `json:"not_at_any_other_time,omitempty"`
	Channels Channels `json:"channels,omitempty"`
	RejectOnError bool `json:"reject_on_error,omitempty"`
	Checked bool `json:"checked,omitempty"`
}

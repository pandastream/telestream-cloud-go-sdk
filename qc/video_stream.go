/*
 * Qc API
 *
 * QC API
 *
 * API version: 2.0.0
 * Contact: cloudsupport@telestream.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package qc

type VideoStream struct {

	// Video stream duration measured in seconds.
	Duration float32 `json:"duration,omitempty"`

	Codec string `json:"codec,omitempty"`

	Width int32 `json:"width,omitempty"`

	Height int32 `json:"height,omitempty"`

	// Video stream bitrate measured in bps
	Bitrate int32 `json:"bitrate,omitempty"`

	Fps float32 `json:"fps,omitempty"`
}

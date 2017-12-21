/* 
 * Flip API
 *
 * Description
 *
 * OpenAPI spec version: 3.1.0
 * Contact: cloudsupport@telestream.net
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package flip

type VideoUploadBody struct {

	// Size of the file that will be uploaded in `bytes`.
	FileSize int64 `json:"file_size"`

	// Name of the file that will be uploaded.
	FileName string `json:"file_name"`

	// A list of names of additional files that will be uploaded.
	ExtraFiles []ExtraFile `json:"extra_files,omitempty"`

	// A comma-separated list of profile names or IDs to be used during encoding. Alternatively, specify none so no encodings will created right away.
	Profiles string `json:"profiles,omitempty"`

	PathFormat string `json:"path_format,omitempty"`

	Payload map[string]string `json:"payload,omitempty"`

	ExtraVariables map[string]string `json:"extra_variables,omitempty"`

	// URL pointing to an image that will be used asa watermark.
	WatermarkUrl string `json:"watermark_url,omitempty"`

	// Determines distance between the left edge of a video and the left edge of a watermark image. Can be specified in pixels or percents. This parameter can be set only if watermark_right is not.
	WatermarkLeft string `json:"watermark_left,omitempty"`

	// Determines distance between the top edge of a video and the top edge of a watermark image. Can be specified in pixels or percents. This parameter can be set only if watermark_bottom is not.
	WatermarkTop string `json:"watermark_top,omitempty"`

	// Determines distance between the right edge of a video and the right edge of a watermark image. Can be specified in pixels or percents. This parameter can be set only if watermark_left is not.
	WatermarkRight string `json:"watermark_right,omitempty"`

	// Determines distance between the bottom edge of a video and the bottom edge of a watermark image. Can be specified in pixels or percents. This parameter can be set only if watermark_top is not.
	WatermarkBottom string `json:"watermark_bottom,omitempty"`

	// Determines width of the watermark image. Should be specified in pixels.
	WatermarkWidth string `json:"watermark_width,omitempty"`

	// Determines width of the watermark image. Should be specified in pixels.
	WatermarkHeight string `json:"watermark_height,omitempty"`

	// Length of the uploaded video. Should be formatted as follows: HH:MM:SS
	ClipLength string `json:"clip_length,omitempty"`

	// Clip starts at a specific offset.
	ClipOffset string `json:"clip_offset,omitempty"`

	MultiChunk bool `json:"multi_chunk,omitempty"`
}

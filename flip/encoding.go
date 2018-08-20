/*
 * Flip API
 *
 * Description
 *
 * API version: 2.0.1
 * Contact: cloudsupport@telestream.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flip

type Encoding struct {

	// A unique identifier of an Encoding.
	Id string `json:"id,omitempty"`

	// Audio bitrate (in bits/s).
	AudioBitrate int32 `json:"audio_bitrate,omitempty"`

	// A number of audio channels.
	AudioChannels int32 `json:"audio_channels,omitempty"`

	// A codec that is used to encode audio streams.
	AudioCodec string `json:"audio_codec,omitempty"`

	// A number of samples of audio carried per second.
	AudioSampleRate int32 `json:"audio_sample_rate,omitempty"`

	// A date and time when the Encoding has been created.
	CreatedAt string `json:"created_at,omitempty"`

	Duration int32 `json:"duration,omitempty"`

	EncodingProgress int32 `json:"encoding_progress,omitempty"`

	EncodingTime int32 `json:"encoding_time,omitempty"`

	// A class of an error that has occurred during the encoding process. It is present only if the encoding status is equal to `fail`.
	ErrorClass string `json:"error_class,omitempty"`

	// A message that explains why the encoding process has resulted in an error. It is present only if the encoding status is equal to `fail`.
	ErrorMessage string `json:"error_message,omitempty"`

	ExternalId string `json:"external_id,omitempty"`

	// Extension of the output file.
	Extname string `json:"extname,omitempty"`

	// A size of the output file.
	FileSize int64 `json:"file_size,omitempty"`

	// An array of output file names.
	Files []string `json:"files,omitempty"`

	// Number of frames per second.
	Fps float32 `json:"fps,omitempty"`

	// Height of the output video.
	Height int32 `json:"height,omitempty"`

	// Width of the output video.
	Width int32 `json:"width,omitempty"`

	// An URL pointing to a logfile.
	LogfileUrl string `json:"logfile_url,omitempty"`

	// A mime type of the encoded file.
	MimeType string `json:"mime_type,omitempty"`

	ParentEncodingId string `json:"parent_encoding_id,omitempty"`

	Path string `json:"path,omitempty"`

	// An id of a related Profile.
	ProfileId string `json:"profile_id,omitempty"`

	// A name of the used Profile.
	ProfileName string `json:"profile_name,omitempty"`

	Screenshots []string `json:"screenshots,omitempty"`

	// A date and time when the encoding process has been started
	StartedEncodingAt string `json:"started_encoding_at,omitempty"`

	// Determines at what stage the encoding process is at the moment.
	Status string `json:"status,omitempty"`

	// A date and time when a Encoding has been updated last time.
	UpdatedAt string `json:"updated_at,omitempty"`

	// video bitrate (in bits/s)
	VideoBitrate int32 `json:"video_bitrate,omitempty"`

	// A codec that is used to encode video streams.
	VideoCodec string `json:"video_codec,omitempty"`

	// An id of a related Video object
	VideoId string `json:"video_id,omitempty"`
}

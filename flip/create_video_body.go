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

type CreateVideoBody struct {

	// An URL pointing to a source file.
	SourceUrl string `json:"source_url,omitempty"`

	// Comma-separated list of profile names or IDs to be used during encoding. Alternatively, specify none so no encodings are created yet.
	Profiles string `json:"profiles,omitempty"`

	// Arbitrary string stored along the Video object.
	Payload string `json:"payload,omitempty"`

	// String-encoded JSON describing profiles pipeline.
	Pipeline string `json:"pipeline,omitempty"`

	// A list of urls pointing to remote subtitle files.
	SubtitleFiles []string `json:"subtitle_files,omitempty"`

	ExtraFiles map[string][]string `json:"extra_files,omitempty"`

	ExtraVariables map[string]string `json:"extra_variables,omitempty"`

	PathFormat string `json:"path_format,omitempty"`
}

/*
 * Tts API
 *
 * Description
 *
 * API version: 2.0.0
 * Contact: cloudsupport@telestream.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package tts

type Project struct {

	// The ID of the Project.
	Id string `json:"id,omitempty"`

	// The name of the Project.
	Name string `json:"name,omitempty"`

	// The description of the Project.
	Description string `json:"description,omitempty"`

	// Determines a stage of training.
	Status string `json:"status,omitempty"`

	// The language code of model.
	Language string `json:"language,omitempty"`

	// The sample rate of model.
	SampleRate int32 `json:"sample_rate,omitempty"`

	// If true, the service replaces profanity from output with asterisks.
	ProfanityFilter bool `json:"profanity_filter,omitempty"`

	// Indicates whether video preview should be generated.
	GenerateProxy bool `json:"generate_proxy,omitempty"`

	// Words used for model training, separated by space.
	CustomWords string `json:"custom_words,omitempty"`

	// A date and time when the project was created
	CreatedAt string `json:"created_at,omitempty"`

	// A date and time when the project was updated
	UpdatedAt string `json:"updated_at,omitempty"`
}

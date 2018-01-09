/*
 * Flip API
 *
 * Description
 *
 * API version: 3.1.0
 * Contact: cloudsupport@telestream.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flip

type ProfileBody struct {

	// a name of a preset that a profile will use.
	PresetName string `json:"preset_name"`

	// todo
	AdvancedFpsConversion string `json:"advanced_fps_conversion,omitempty"`

	// Default is \"letterbox\".
	AspectMode string `json:"aspect_mode,omitempty"`

	// Sets the desired display aspect ratio. By default it is not set.
	AspectRatio string `json:"aspect_ratio,omitempty"`

	// audio bitrate (in bits/s)
	AudioBitrate int32 `json:"audio_bitrate,omitempty"`

	// Sets the number of audio channels. By default it is not set.
	AudioChannels int32 `json:"audio_channels,omitempty"`

	// A channel layout specifies the spatial disposition of the channels in a multi-channel audio stream.
	AudioChannelsLayout string `json:"audio_channels_layout,omitempty"`

	// Sets the number of audio channels per track.
	AudioChannelsPerTrack string `json:"audio_channels_per_track,omitempty"`

	// Audio codec that will be used by the profile.
	AudioCodec string `json:"audio_codec,omitempty"`

	// Specifies an audio container.
	AudioFormat string `json:"audio_format,omitempty"`

	// Packet identifier used by MPEG formats.
	AudioPid string `json:"audio_pid,omitempty"`

	// Sets an audio profile.
	AudioProfile string `json:"audio_profile,omitempty"`

	// The number of samples of audio carried per second.
	AudioSampleRate int32 `json:"audio_sample_rate,omitempty"`

	// Sets the number of audio streams.
	AudioStreams int32 `json:"audio_streams,omitempty"`

	// class of the AVC-Intra video coding.
	AvcintraClass string `json:"avcintra_class,omitempty"`

	// Sets the buffer size, and can be 1-2 seconds for most gaming screencasts, and up to 5 seconds for more static content. You will have to experiment to see what looks best for your content.
	BufferSize string `json:"buffer_size,omitempty"`

	// todo
	BufferSizeInPackets string `json:"buffer_size_in_packets,omitempty"`

	// Sets the clip’s duration.
	ClipLength string `json:"clip_length,omitempty"`

	// Clip starts at a specific offset.
	ClipOffset string `json:"clip_offset,omitempty"`

	// One of add (adds captions as a separate streams) or burn (burns captions on video stream using the first subtitle file). By default it is not set.
	ClosedCaptions string `json:"closed_captions,omitempty"`

	DashProfile string `json:"dash_profile,omitempty"`

	// One of `keep_fps` or `double_fps`. By default it is not set.
	Deinterlace string `json:"deinterlace,omitempty"`

	DeinterlaceFrames string `json:"deinterlace_frames,omitempty"`

	DnxhdType string `json:"dnxhd_type,omitempty"`

	Encryption bool `json:"encryption,omitempty"`

	// File extension.
	Extname string `json:"extname,omitempty"`

	// Null value copy the original fps. By default it is not set.
	Fps float32 `json:"fps,omitempty"`

	// Array of offsets (Frames or seconds).
	FrameOffsets string `json:"frame_offsets,omitempty"`

	// Thumbnail interval (Frames or seconds).
	FrameInterval string `json:"frame_interval,omitempty"`

	// Evenly spaced number of generated screenshots. By default it is not set.
	FrameCount int32 `json:"frame_count,omitempty"`

	// A specified set of constraints that indicate a degree of required decoder performance for a profile.
	H264Level string `json:"h264_level,omitempty"`

	// Profiles represent a sub-set of the encoding techniques available in H.264.
	H264Profile string `json:"h264_profile,omitempty"`

	// Use this option to change settings based upon the specifics of your input
	H264Tune string `json:"h264_tune,omitempty"`

	// Height in pixels.
	Height int32 `json:"height,omitempty"`

	ImxType string `json:"imx_type,omitempty"`

	Interlace string `json:"interlace,omitempty"`

	// Adds a key frame every N frames. Default is 250, adds a key frame every 250 frames.
	KeyframeInterval int32 `json:"keyframe_interval,omitempty"`

	// todo
	KeyframeRate float32 `json:"keyframe_rate,omitempty"`

	// Set max bitrate tolerance (in bits/s). By default this is not set
	MaxRate int32 `json:"max_rate,omitempty"`

	MergeAudioStreams string `json:"merge_audio_streams,omitempty"`

	// Unique machine-readable name that will identify the profile. Helpful later on for filtering encodings by profile.
	Name string `json:"name,omitempty"`

	// Specify the directory where the output files should be stored. By default it is not set. More information about this [here](https://cloud.telestream.net/docs#path-format---know-how).
	OutputsPathFormat string `json:"outputs_path_format,omitempty"`

	PmtPid string `json:"pmt_pid,omitempty"`

	ProresFormat string `json:"prores_format,omitempty"`

	SegmentTime string `json:"segment_time,omitempty"`

	Size string `json:"size,omitempty"`

	Tar bool `json:"tar,omitempty"`

	TransportRate string `json:"transport_rate,omitempty"`

	TsPids string `json:"ts_pids,omitempty"`

	// Upscale the video resolution to match your profile. Default is `true`.
	Upscale bool `json:"upscale,omitempty"`

	// Pattern utilised to match HLS.Variant presets by name. Default is hls.*.
	Variants string `json:"variants,omitempty"`

	VideoBitrate int32 `json:"video_bitrate,omitempty"`

	VideoPid string `json:"video_pid,omitempty"`

	// Distance from the bottom of the video frame in pixels or percentage of video frame height. Works like CSS. Default is `0`.
	WatermarkBottom string `json:"watermark_bottom,omitempty"`

	// Height of the watermark image in pixels or percentage of video frame height. Default is no resizing
	WatermarkHeight string `json:"watermark_height,omitempty"`

	// Distance from the left of the video frame in pixels or percentage of video frame width. Works like CSS. Default is `0`.
	WatermarkLeft string `json:"watermark_left,omitempty"`

	// Distance from the right of the video frame in pixels or percentage of video frame width. Works like CSS. Default is `0`.
	WatermarkRight string `json:"watermark_right,omitempty"`

	// Distance from the top of the video frame in pixels or percentage of video frame height. Works like CSS. Default is `0`.
	WatermarkTop string `json:"watermark_top,omitempty"`

	// Url of a watermark image.
	WatermarkUrl string `json:"watermark_url,omitempty"`

	// Width of the watermark image in pixels or percentage of video frame width. Default is `no resizing`.
	WatermarkWidth string `json:"watermark_width,omitempty"`

	// Width in pixels.
	Width int32 `json:"width,omitempty"`

	X264Options string `json:"x264_options,omitempty"`

	X265Options string `json:"x265_options,omitempty"`

	XdcamFormat string `json:"xdcam_format,omitempty"`

	// Remove audio from input video file. By default it is set to `false`.
	MuteAudioTracks bool `json:"mute_audio_tracks,omitempty"`

	ByteRangeRequests string `json:"byte_range_requests,omitempty"`

	Lang string `json:"lang,omitempty"`

	UseEditlist string `json:"use_editlist,omitempty"`

	AudioMap string `json:"audio_map,omitempty"`

	AudioStreamId string `json:"audio_stream_id,omitempty"`

	Bumpers string `json:"bumpers,omitempty"`

	// Determines a preset that is used by encoders.
	CodecPreset string `json:"codec_preset,omitempty"`

	ColorMetadata string `json:"color_metadata,omitempty"`

	// Distance (in pixels) from the bottom edge of the screen from which you want your crop to be done.
	CropInputBottom string `json:"crop_input_bottom,omitempty"`

	// Width of the cropped image in pixels.
	CropInputHeight string `json:"crop_input_height,omitempty"`

	// Distance (in pixels) from the left edge of the screen from which you want your crop to be done.
	CropInputLeft string `json:"crop_input_left,omitempty"`

	// Distance (in pixels) from the right edge of the screen from which you want your crop to be done.
	CropInputRight string `json:"crop_input_right,omitempty"`

	// Distance (in pixels) from the top edge of the screen from which you want your crop to be done.
	CropInputTop string `json:"crop_input_top,omitempty"`

	// Height of the cropped image in pixels.
	CropInputWidth string `json:"crop_input_width,omitempty"`

	DynamicRecipe string `json:"dynamic_recipe,omitempty"`

	PlaylistType string `json:"playlist_type,omitempty"`

	PresetVersion string `json:"preset_version,omitempty"`

	SegmentDelimiter string `json:"segment_delimiter,omitempty"`

	SwsFlags string `json:"sws_flags,omitempty"`

	TelestreamBlockSize string `json:"telestream_block_size,omitempty"`

	// Minimum value is 0, maximum is 4.
	TelestreamBlurScaler string `json:"telestream_blur_scaler,omitempty"`

	// Minimum value is 0, maximum is 4.
	TelestreamCostScaler string `json:"telestream_cost_scaler,omitempty"`

	// Minimum value is 0, maximum is 2.
	TelestreamSearchLengthScaler string `json:"telestream_search_length_scaler,omitempty"`

	TelestreamSubpelMode string `json:"telestream_subpel_mode,omitempty"`

	Trailers string `json:"trailers,omitempty"`

	VantageGroupId string `json:"vantage_group_id,omitempty"`

	WatermarkBumpers string `json:"watermark_bumpers,omitempty"`

	WatermarkTrailers string `json:"watermark_trailers,omitempty"`

	WorkorderCriteria *interface{} `json:"workorder_criteria,omitempty"`

	// Enable more sensitive pulldown removal algorithm.
	TachyonAllowRemovePulldown bool `json:"tachyon_allow_remove_pulldown,omitempty"`

	// If the images you are converting are composited 29.976, but the pulldown pattern was not adhered to when performing the composite, this setting is required to remove combing artifacts. It will also remove combing artifacts related to very poor 3:2 cadence.
	TachyonEnablePostPulldownFilter bool `json:"tachyon_enable_post_pulldown_filter,omitempty"`

	TachyonMediaHintIsCartoon bool `json:"tachyon_media_hint_is_cartoon,omitempty"`

	// Remove chroma noise during the analysis of a video.
	TachyonMediaHintHasChromaNoise bool `json:"tachyon_media_hint_has_chroma_noise,omitempty"`

	// When pulldown is not achieved due to extremely broken cadence, or other factors like highly mixed content or if chroma noise masks motion, the pulldown engine may fall back to de-interlacing rather than removing telecine. If that's a case, a more sensitive pulldown pattern can be used. This algorithm favors inverse telecine and with lower thresholds for triggering pulldown identification, will maximize the number of progressive frames created from the video.
	TachyonMoreSensitiveRemovePulldown bool `json:"tachyon_more_sensitive_remove_pulldown,omitempty"`

	TachyonAllowAddStandardPd bool `json:"tachyon_allow_add_standard_pd,omitempty"`

	// Allows 2:2 (PSF) Insertion. Creates a new series of frames which are based on duplicating the field an interlacing it into top/bottom field. Maintains a film-look.
	TachyonAllowAdd22pd bool `json:"tachyon_allow_add_2_2pd,omitempty"`

	// Allows 4:4 Insertion. Repeats each progressive frame twice on output (motion rate is halved). This setting is used when you want to convert to high progressive frame rates (i.e. 50p/59.94p/60p) but want to preserve film qualities (low motion rate, such as 24p).
	TachyonAllowAdd44pd bool `json:"tachyon_allow_add_4_4pd,omitempty"`

	// 2:3 Insertion. inserts a standard 2:3 telecine pattern to 23.976p video stream to achieve a 29.97i frame rate
	TachyonAllowAdd46pd bool `json:"tachyon_allow_add_4_6pd,omitempty"`

	// Allows Euro Insertion. For field based interpolation rather than pixel-based. This is designed for interlaced or progressive integer frame rate conversions that are being converted to interlaced outputs. This method is valid for 24p to 50i conversions only.
	TachyonAllowAddEuroPd bool `json:"tachyon_allow_add_euro_pd,omitempty"`

	// Allows Adaptive Insertion. For field-based interpolation rather than using pixel-based interpolation. This algorithm is designed for both integer and non-integer frame rate conversion targets - as long as one of them is a non-integer rate (23.976, 29.97, 59.94, etc). This creates NTSC-PAL conversions clean of motion artifacts at the expense of potential slight stutter. Stutter is most noticeable with material that has smooth and uniform motion.
	TachyonAllowAddAdaptivePd bool `json:"tachyon_allow_add_adaptive_pd,omitempty"`

	// This setting determines how much Tachyon will trust motion vectors in the creation of new images
	TachyonMotionAmount string `json:"tachyon_motion_amount,omitempty"`

	// This option specifies the transition region size between fallback areas and motion compensated areas. A larger fallback size allows more blending (feathering) to occur between the regions.
	TachyonFallbackSize string `json:"tachyon_fallback_size,omitempty"`

	// This option specifies the size of a motion block.
	TachyonMblockSize string `json:"tachyon_mblock_size,omitempty"`

	TachyonCutDetectionSensitivity float32 `json:"tachyon_cut_detection_sensitivity,omitempty"`

	// Enables the trusted metadata framework.
	Eac3EvolutionEnable bool `json:"eac3_evolution_enable,omitempty"`

	// Selects the type of audio service. **For 1/0 Voiceover will be used when Voiceover/Karaoke is selected. For 2/0 and above Karaoke will be used.
	Eac3BitstreamMode string `json:"eac3_bitstream_mode,omitempty"`

	// Applies a 90-degree phase shift to the surround channels; necessary if the output file is being decoded by a Dolby Surround Pro Logic or Pro Logic II decoder.
	Eac3NinetyDegreePhaseShift bool `json:"eac3_ninety_degree_phase_shift,omitempty"`

	// Attenuates the surround channels by 3 dB before encoding.
	Eac3ThreeDecibelAttenuation bool `json:"eac3_three_decibel_attenuation,omitempty"`

	// Applies a 120 Hz eighth order lowpass filter to the LFE input prior to encoding.
	Eac3EnableLfeLowPassFilter bool `json:"eac3_enable_lfe_low_pass_filter,omitempty"`

	// Allows audio that has passed through an A/D conversion stage to be marked as such.
	Eac3AnalogToDigitalConverterType string `json:"eac3_analog_to_digital_converter_type,omitempty"`

	Eac3StereoDownmixPreference string `json:"eac3_stereo_downmix_preference,omitempty"`

	// Indicates the level shift applied to the center channel when adding to the left and right outputs during a downmix to a Lt/Rt output.
	Eac3LtRtCenterMixLevel string `json:"eac3_lt_rt_center_mix_level,omitempty"`

	// Indicates the level shift applied to the surround channel when adding to the left and right outputs during a downmix to a Lt/Rt output.
	Eac3LtRtSurroundMixLevel string `json:"eac3_lt_rt_surround_mix_level,omitempty"`

	// Indicates the level shift applied to the center channel when adding to the left and right outputs during a downmix to a Lo/Ro output.
	Eac3LoRoCenterMixLevel string `json:"eac3_lo_ro_center_mix_level,omitempty"`

	// Indicates the level shift applied to the surround channel when adding to the left and right outputs during a downmix to a Lo/Ro output.
	Eac3LoRoSurroundMixLevel string `json:"eac3_lo_ro_surround_mix_level,omitempty"`

	// Indicates whether the audio stream was encoded using Dolby EX.
	Eac3SurroundExMode string `json:"eac3_surround_ex_mode,omitempty"`

	// Dynamic Range Control for Line Mode.
	Eac3DrcLineModeProfile string `json:"eac3_drc_line_mode_profile,omitempty"`

	// Dynamic Range Control for RF Mode.
	Eac3DrcRfModeProfile string `json:"eac3_drc_rf_mode_profile,omitempty"`

	// Represents the volume level of dialog in the audio stream which can be used by a Dolby Digital decoder. This aids the decoder in matching volume between program sources. Minimum value is 1, maximum is 31.
	Eac3DialogNormalization int32 `json:"eac3_dialog_normalization,omitempty"`

	Eac3RoomType string `json:"eac3_room_type,omitempty"`

	// Minimum value is 80, maximum is 111.
	Eac3MixingLevel int32 `json:"eac3_mixing_level,omitempty"`

	// Indicates whether the encoded bitstream is copyright protected.
	Eac3CopyrightProtected bool `json:"eac3_copyright_protected,omitempty"`

	// Indicates whether the encoded bitstream is the master version, or a copy.
	Eac3OriginalBitstream bool `json:"eac3_original_bitstream,omitempty"`

	// Human-readable name.
	Title string `json:"title,omitempty"`

	// If set, timestamps will be added to your videos. By default this is not set.
	TimeCode bool `json:"time_code,omitempty"`
}

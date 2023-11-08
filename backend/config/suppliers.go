package config

// UploadVideoRequest combines universal and platform-specific settings
type UploadVideoRequest struct {
	Title           string            `json:"title,omitempty"`
	PrivacyLevel    string            `json:"privacy_level,omitempty"`
	TikTokConfig    TikTokSettings    `json:"tiktok,omitempty"`
	YouTubeConfig   YouTubeSettings   `json:"youtube,omitempty"`
	InstagramConfig InstagramSettings `json:"instagram,omitempty"`
}

type TikTokSettings struct {
	DisableDuet bool `json:"disable_duet,omitempty"`
}

type YouTubeSettings struct {
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

type InstagramSettings struct {
}

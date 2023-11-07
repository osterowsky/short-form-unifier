package config

// UploadVideoRequest combines universal and platform-specific settings
type UploadVideoRequest struct {
	Title         string          `json:"title,omitempty"`
	PrivacyLevel  string          `json:"privacy_level,omitempty"`
	TikTokConfig  TikTokSettings  `json:"tiktok,omitempty"`
	YouTubeConfig YouTubeSettings `json:"youtube,omitempty"`
}

// Specific fields for TikTok uploads
type TikTokSettings struct {
	DisableDuet bool `json:"disable_duet,omitempty"`
}

// Specific fields for YouTube uploads
type YouTubeSettings struct {
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

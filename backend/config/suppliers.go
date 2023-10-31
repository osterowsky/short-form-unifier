package config

// Paylod which we will need to handle
type VideoRequest struct {
	// Universal settings for all suppliers
	Title        string `json:"title"`
	PrivacyLevel string `json:"privacy_level"`

	// Youtube specific settings
	YoutubeConfig *YoutubeConfig `json:"youtube_config,omitempty"`
	TiktokConfig  *TiktokConfig  `json:"tiktok_config,omitempty"`
}

type YoutubeConfig struct {
	Description             string   `json:"description"`
	Tags                    []string `json:"tags"`
	SelfDeclaredMadeForKids bool     `json:"self_declared_made_for_kids"`
}

type TiktokConfig struct {
	DisableDuet    bool `json:"disable_duet,omitempty"`
	DisableComment bool `json:"disable_comment,omitempty"`
}

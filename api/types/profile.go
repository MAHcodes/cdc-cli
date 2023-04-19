package types

type Profile struct {
	ID         string `json:"@id"`
	Country    string `json:"country"`
	Followers  int    `json:"followers"`
	IsStreamer bool   `json:"is_streamer"`
	Joined     int    `json:"joined"`
	LastOnline int    `json:"last_online"`
	League     string `json:"league"`
	PlayerID   int    `json:"player_id"`
	Status     string `json:"status"`
	URL        string `json:"url"`
	Username   string `json:"username"`
	Verified   bool   `json:"verified"`
}

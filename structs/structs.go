package structs

import (
	"os"
)

type DiscordConfig struct {
	WelcomeID string `json: "welcome_id"`
}

var Discord DiscordConfig

func InitDiscordConfig() error {
	file, err := os.Open("./config/discord.json")
}

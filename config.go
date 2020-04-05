package main

type MixerConfig struct {
	XAirDevice	string 			`toml:"xair_device"`
	Level 		int 			`toml:"level"`

	Scenes		[]SceneConfig 	`toml:"scene"`
}

type SceneConfig struct {
	Channels	[]ChannelConfig `toml:"channel"`
}

type ChannelConfig struct {
	Channel 	int 	`toml:"channel"`
	Bus 		int 	`toml:"bus"`
	Level 		int 	`toml:"level"`
}
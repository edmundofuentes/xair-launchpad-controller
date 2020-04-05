package main

import (
	"log"

	"github.com/BurntSushi/toml"

	launchpad "edmundofuentes/launchpad/mini/mk1"
	//xair "edmundofuentes/xair/xr12"
)

const ConfigPath = "config.toml"


func main() {

	// Read Configurations from file
	var config MixerConfig
	if _, err := toml.DecodeFile(ConfigPath, &config); err != nil {
		// handle error
		log.Fatalf("Invalid Configuration file provided: %v", err)
	}


	// Connect to the Launchpad device via MIDI
	pad, err := launchpad.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer pad.Close()

	pad.Clear()


	/*
	// Connect to the XAir mixer via MIDI
	xair, err := xair.Open(config.XAirDevice)
	if err != nil {
		log.Fatal(err)
	}
	defer xair.Close()
	*/


	// Initialize the Mixer app
	mixer, err := NewMixer(config, nil, pad)
	if err != nil {
		log.Fatal(err)
	}


	err = mixer.Run()

	if err != nil {
		log.Fatal(err)
	}
}
package main

import (
	launchpad "edmundofuentes/launchpad/mini/mk1"
)

func (m *Mixer) HandleButtonHit(hit launchpad.Hit) {
	// Check if the button is a scene button
	if hit.Y == 8 {
		m.HandleSceneButtonHit(hit.X)
		return
	}

	// Check if the button is a MainChannel
	if hit.X == 8 {
		m.HandleMainChannelButtonHit(hit.Y)
		return
	}

	m.HandleChannelButtonHit(hit.X, hit.Y)
}

func (m *Mixer) HandleSceneButtonHit(x int) {
	m.checkSceneRange(x)

	// check if the scene is enabled
	if m.Scenes[x].Enabled == false {
		// do nothing
		return
	}

	// if the scene button is the current active scene, do nothing
	if m.ActiveScene == x {
		return
	}

	m.StartChangeToScene(x)
}

func (m *Mixer) HandleMainChannelButtonHit(y int) {
	// check if the hit was on the mute/unmute row
	if y == 7 {
		m.ToggleMainChannelMute()
		m.DrawMainChannel()
		return
	}

	// set the new level according to the hit
	l := 7 - y
	m.SetMainChannelLevel(l)
	m.DrawMainChannel()
}

func (m *Mixer) HandleChannelButtonHit(x, y int) {
	m.checkChannelRange(x)

	// check if the channel is enabled
	if m.Scenes[m.ActiveScene].Channels[x].Enabled == false {
		// do nothing
		return
	}

	// check if the hit was on the mute/unmute row
	if y == 7 {
		m.ToggleChannelMute(x)
		m.DrawChannel(x)
		return
	}

	// set the new level according to the hit
	l := 7 - y
	m.SetChannelLevel(x, l)
	m.DrawChannel(x)
}

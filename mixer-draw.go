package main

import (
	launchpad "edmundofuentes/launchpad/mini/mk1"
)

func (m *Mixer) Draw() {
	m.DrawSceneBar()

	for i := 0; i < 8; i++ {
		m.DrawChannel(i)
	}

	m.DrawMainChannel()
}

func (m *Mixer) DrawSceneBar() {
	for s := 0; s < 8; s++ {

		if m.Scenes[s].Enabled == false {
			// the scene is not enabled, the indicator should be off
			m.Pad.LightColor(s, 8, launchpad.Off) // scenes are the top row (y = 8)
		} else {
			if s == m.ActiveScene {
				// this is the currently enabled scene
				m.Pad.LightColor(s, 8, launchpad.GreenHigh) // scenes are the top row (y = 8)
			} else {
				// this is another available scene, but not the current
				m.Pad.LightColor(s, 8, launchpad.GreenLow) // scenes are the top row (y = 8)
			}
		}

	}
}

func (m *Mixer) DrawMainChannel() {
	if m.MainActive {
		// the chanel is active (unmuted)
		m.Pad.LightColor(8, 7, launchpad.GreenHigh) // the bottom row (y=7) displays the mute status

		// display the current level
		for r := 0; r < 7; r++ { // r: row
			if (7 - m.MainLevel) <= r {
				// active rows, the top 2 pads should be red, the 3rd green, the rest yellow
				if r == 0 || r == 1 {
					m.Pad.LightColor(8, r, launchpad.OrangeHigh)
				} else if r == 2 {
					m.Pad.LightColor(8, r, launchpad.YellowGreen)
				} else {
					m.Pad.LightColor(8, r, launchpad.YellowHigh)
				}

			} else {
				m.Pad.LightColor(8, r, launchpad.Off)
			}
		}
	} else {
		// the chanel is not active (muted)
		m.Pad.LightColor(8, 7, launchpad.RedLow) // the bottom row (y=7) displays the mute status

		// display the current level
		for r := 0; r < 7; r++ { // r: row
			if (7 - m.MainLevel) <= r {
				m.Pad.LightColor(8, r, launchpad.YellowLow)
			} else {
				m.Pad.LightColor(8, r, launchpad.Off)
			}
		}
	}
}

func (m *Mixer) DrawChannel(c int) {
	m.checkChannelRange(c)

	if m.Scenes[m.ActiveScene].Channels[c].Enabled == false {
		// the current channel is not enabled, set the column to 0
		for y := 7; y >= 0; y-- {
			m.Pad.LightColor(c, y, launchpad.Off)
		}
	} else {
		// the channel is enabled
		if m.Scenes[m.ActiveScene].Channels[c].Active {
			// the chanel is active (unmuted)
			m.Pad.LightColor(c, 7, launchpad.GreenHigh) // the bottom row (y=7) displays the mute status

			// display the current level
			for r := 0; r < 7; r++ { // r: row
				if (7 - m.Scenes[m.ActiveScene].Channels[c].Level) <= r {

					// active rows, the top 2 pads should be red, the 3rd green, the rest yellow
					if r == 0 || r == 1 {
						m.Pad.LightColor(c, r, launchpad.OrangeHigh)
					} else if r == 2 {
						m.Pad.LightColor(c, r, launchpad.YellowGreen)
					} else {
						m.Pad.LightColor(c, r, launchpad.YellowHigh)
					}

				} else {
					m.Pad.LightColor(c, r, launchpad.Off)
				}
			}
		} else {
			// the chanel is not active (muted)
			m.Pad.LightColor(c, 7, launchpad.RedLow) // the bottom row (y=7) displays the mute status

			// display the current level
			for r := 0; r < 7; r++ { // r: row
				if (7 - m.Scenes[m.ActiveScene].Channels[c].Level) <= r {
					m.Pad.LightColor(c, r, launchpad.YellowLow)
				} else {
					m.Pad.LightColor(c, r, launchpad.Off)
				}
			}
		}
	}
}

func (m *Mixer) DebugColors() {
	m.Pad.LightColor(0, 0, launchpad.RedHigh)
	m.Pad.LightColor(1, 0, launchpad.RedMid)
	m.Pad.LightColor(2, 0, launchpad.RedLow)

	m.Pad.LightColor(0, 0, launchpad.GreenHigh)
	m.Pad.LightColor(1, 0, launchpad.GreenMid)
	m.Pad.LightColor(2, 0, launchpad.GreenLow)
}

package main

import (
	"errors"
	"fmt"
	"time"

	launchpad "edmundofuentes/launchpad/mini/mk1"
	xair "edmundofuentes/xair/xr12"
)

type Mixer struct {
	XAir 	*xair.XAir
	Pad 	*launchpad.Launchpad

	Scenes [8]Scene

	ActiveScene int

	MainLevel int // Main Channel volume level
	MainActive bool // Main Channel mute status

	Refresh chan bool

	LoadingState bool
	LoadingStateDone chan bool

	LoadingAnimation LoadingAnimation
}

func NewMixer(config MixerConfig, xair *xair.XAir, pad *launchpad.Launchpad) (*Mixer, error) {
	mixer := &Mixer{
		XAir: xair,
		Pad: pad,
	}

	if len(config.Scenes) > 8 {
		return mixer, errors.New(fmt.Sprintf("The maximum number of scenes that can be configured is 8, got %d", len(config.Scenes)))
	}

	if len(config.Scenes) == 0 {
		return mixer, errors.New(fmt.Sprintf("At least one scene configuration is required to initialize"))
	}

	for s := 0; s < len(config.Scenes); s++ {
		// Initialize the Scene
		mixer.Scenes[s].Enabled = true

		if len(config.Scenes[s].Channels) == 0 {
			return mixer, errors.New(fmt.Sprintf("Scene %d does not have any configured channels", s))
		}
		if len(config.Scenes[s].Channels) > 8 {
			return mixer, errors.New(fmt.Sprintf("Scene %d exceeds the maximum number of channels per scene is (max 8), got %d", s, len(config.Scenes[s].Channels)))
		}

		for c := 0; c < len(config.Scenes[s].Channels); c++ {
			mixer.Scenes[s].Channels[c].Enabled = true
			mixer.Scenes[s].Channels[c].Channel = config.Scenes[s].Channels[c].Channel
			mixer.Scenes[s].Channels[c].Bus 	= config.Scenes[s].Channels[c].Bus
			mixer.Scenes[s].Channels[c].Level 	= config.Scenes[s].Channels[c].Level
		}
	}

	// default to the first scene
	mixer.ActiveScene = 0

	// load default level
	mixer.MainLevel = config.Level

	// create a buffered channel for Screen Refresh requests
	mixer.Refresh = make(chan bool, 1)

	return mixer, nil
}

func (m *Mixer) Run() error {
	// by default, start on the scene 0
	m.StartChangeToScene(0)

	buttonHits := m.Pad.Listen()

	for {
		if m.LoadingState {
			select {
			case <-m.LoadingStateDone:
				break
			case <-m.LoadingAnimation.ticker.C:
				m.LoadingAnimation.DrawTick()
			case <-buttonHits:
				// hits should be discarded when loading
			}
		} else {
			select {
			case <-m.Refresh:
				m.Draw()
			case hit := <-buttonHits:
				//color := rand.Intn(8) + 1 // 1 thru 9
				//pad.LightColor(hit.X, hit.Y, color)
				//fmt.Printf("Button: %d %d\n", hit.X, hit.Y)

				m.HandleButtonHit(hit)
			}
		}
	}

	return nil
}

func (m *Mixer) StartChangeToScene(s int) {
	m.checkSceneRange(s)

	/* TODO: Decide this.. if we restrict it, we cannot reload our current scene.. think about errors and refreshes?
	if m.Scenes[s].Enabled == false {
		// do nothing, the scene is not even enabled, we shouldn't even be here
		return
	}
	*/

	fmt.Printf("Changing to Scene %d..\n", s)
	m.ActiveScene = s

	m.StartLoadingState(s)

	go m.LoadSceneValues(s)
}

func (m *Mixer) LoadSceneValues(s int) {
	time.Sleep(3 * time.Second)
	fmt.Printf("Scene %d loaded!\n", s)

	m.EndLoadingState()
	m.RequestRefresh()
}

func (m *Mixer) ToggleMainChannelMute() {
	if m.MainActive == false {
		// the channel was muted, unmute
		m.MainActive = true
	} else {
		// the channel was active, mute it
		m.MainActive = false
	}
}

func (m *Mixer) ToggleChannelMute(c int) {
	m.checkChannelRange(c)

	if m.Scenes[m.ActiveScene].Channels[c].Enabled == false {
		// do nothing, the channel is not even enabled, we shouldn't even be here
		return
	}

	if m.Scenes[m.ActiveScene].Channels[c].Active == false {
		// the channel was muted, unmute
		m.Scenes[m.ActiveScene].Channels[c].Active = true
	} else {
		// the channel was active, mute it
		m.Scenes[m.ActiveScene].Channels[c].Active = false
	}
}

func (m *Mixer) SetMainChannelLevel(l int) {
	m.checkChannelLevel(l)

	m.MainLevel = l
}

func (m *Mixer) SetChannelLevel(c, l int) {
	m.checkChannelRange(c)
	m.checkChannelLevel(l)

	if m.Scenes[m.ActiveScene].Channels[c].Enabled == false {
		// do nothing, the channel is not even enabled, we shouldn't even be here
		return
	}

	m.Scenes[m.ActiveScene].Channels[c].Level = l
}


func (m *Mixer) RequestRefresh() {
	select {
	case m.Refresh <- true:
		// refresh requested
	default:
		// another refresh was queued in the buffer
	}
}


//// VALIDATIONS ////

func (m *Mixer) checkSceneRange(c int) {
	if c < 0 || c > 7 {
		m.Die("scene out of bounds")
	}
}

func (m *Mixer) checkChannelRange(c int) {
	if c < 0 || c > 7 {
		m.Die("channel out of bounds")
	}
}

func (m *Mixer) checkChannelLevel(l int) {
	if l < MinLevel || l > MaxLevel {
		m.Die("level out of bounds")
	}
}

func (m *Mixer) Die(s string) {
	m.Pad.Clear()

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if i == j || i == (7-j) {
				m.Pad.LightColor(i, j, launchpad.RedHigh)
			}
		}
	}

	panic(s)
}
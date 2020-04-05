package main

import (
	"time"
)

func (m *Mixer) StartLoadingState(scene int) {
	m.LoadingAnimation = LoadingAnimation{
		pad: m.Pad,
		ticker: time.NewTicker(200 * time.Millisecond),
		scene: scene,
		step: 0,
	}

	m.LoadingState = true
	m.LoadingStateDone = make(chan bool)

	m.Pad.Clear()
}

func (m *Mixer) EndLoadingState() {
	m.LoadingState = false
	close(m.LoadingStateDone)
}

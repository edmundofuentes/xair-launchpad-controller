package main

import (
	launchpad "edmundofuentes/launchpad/mini/mk1"
	"math/rand"
	"time"
)

type LoadingAnimation struct {
	step 	int
	ticker 	*time.Ticker
	pad 	*launchpad.Launchpad
	scene 	int
}

func (a *LoadingAnimation) Reset() {
	a.step = 0
}

func (a *LoadingAnimation) DrawTick() {
	if a.step > 3 {
		a.step = 0
	}

	var matrix [8][8]int

	switch a.step {
	case 0:
		matrix = [8][8]int{
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
		}
	case 1:
		matrix = [8][8]int{
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 1, 1, 1, 1, 0, 0},
			{0, 0, 1, 0, 0, 1, 0, 0},
			{0, 0, 1, 0, 0, 1, 0, 0},
			{0, 0, 1, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
		}
	case 2:
		matrix = [8][8]int{
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 1, 1, 1, 1, 1, 0},
			{0, 1, 0, 0, 0, 0, 1, 0},
			{0, 1, 0, 0, 0, 0, 1, 0},
			{0, 1, 0, 0, 0, 0, 1, 0},
			{0, 1, 0, 0, 0, 0, 1, 0},
			{0, 1, 1, 1, 1, 1, 1, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
		}
	case 3:
		matrix = [8][8]int{
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 0, 0, 0, 0, 0, 0, 1},
			{1, 0, 0, 0, 0, 0, 0, 1},
			{1, 0, 0, 0, 0, 0, 0, 1},
			{1, 0, 0, 0, 0, 0, 0, 1},
			{1, 0, 0, 0, 0, 0, 0, 1},
			{1, 0, 0, 0, 0, 0, 0, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
		}
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {

			d := matrix[i][j]

			if d == 1 {
				color := rand.Intn(2) + 5 // 4 thru 5
				a.pad.LightColor(i, j, color)
			} else {
				a.pad.LightColor(i, j, launchpad.Off)
			}

		}
	}

	// Finally, we'll update the status of the currently selected Scene
	if a.step == 0  {
		a.pad.LightColor(a.scene, 8, launchpad.YellowHigh)
	} else if a.step == 2 {
		a.pad.LightColor(a.scene, 8, launchpad.YellowLow)
	}


	a.step++
}

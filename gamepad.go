package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Check if a start button is pressed (buttons 1,2,3,4,9)
func startButtonPressed(gp ebiten.GamepadID) bool {
	buttons := []ebiten.GamepadButton{1, 2, 3, 4, 9}
	for _, b := range buttons {
		if ebiten.IsGamepadButtonPressed(gp, b) {
			return true
		}
	}
	return false
}

// --------------------
// Unified ball start
// --------------------

// ShouldStartBall returns true if keyboard or any gamepad start buttons are pressed
func ShouldStartBall() bool {
	// Keyboard
	if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyEnter) {
		return true
	}

	// Gamepads 0 and 1
	for _, gp := range []ebiten.GamepadID{0, 1, 2, 3, 4} {
		if startButtonPressed(gp) {
			return true
		}
	}
	return false
}

const deadzone = 0.2

func gamepadUp(id ebiten.GamepadID) bool {
	if !ebiten.IsStandardGamepadLayoutAvailable(id) {
		return false
	}
	// Left pad vertical: -1 = up
	return ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisLeftStickVertical) < -deadzone
}

func gamepadDown(id ebiten.GamepadID) bool {
	if !ebiten.IsStandardGamepadLayoutAvailable(id) {
		return false
	}
	// Left pad vertical: +1 = down
	return ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisLeftStickVertical) > deadzone
}

package textinput

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// New creates a new text input object.
func New() objectspec.TextInput {
	return &textInput{}
}

type textInput struct {
	// Settings.

	// Echo being set to true causes the provided input simply to be echoed back.
	// The provided input goes through the whole stack and is streamed back and
	// forth, but bypasses neural network. This is useful to test the
	// client/server integration of the gRPC stream implementation.
	Echo bool
	// Expectation represents the expectation object. This is used to match
	// against the calculated output. In case there is an expectation given, the
	// neural network tries to calculate an output until it generated one that
	// matches the given expectation.
	Expectation objectspec.Expectation
	// Input represents the input being fed into the neural network. There must
	// be a none empty input given when requesting calculations from the neural
	// network.
	Input string
	// SessionID represents the session the current text request is associated
	// with. This is provided to differentiate streams between different users.
	SessionID string
}

func (ti *textInput) GetEcho() bool {
	return ti.Echo
}

func (ti *textInput) GetExpectation() objectspec.Expectation {
	return ti.Expectation
}

func (ti *textInput) GetInput() string {
	return ti.Input
}

func (ti *textInput) GetSessionID() string {
	return ti.SessionID
}

func (ti *textInput) SetEcho(echo bool) {
	ti.Echo = echo
}

func (ti *textInput) SetExpectation(expectation objectspec.Expectation) {
	ti.Expectation = expectation
}

func (ti *textInput) SetInput(input string) {
	ti.Input = input
}

func (ti *textInput) SetSessionID(sessionID string) {
	ti.SessionID = sessionID
}

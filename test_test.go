package main

import (
	"covertest/state"
	"covertest/syncstate"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetSmall(t *testing.T) {
	s := syncstate.New(new(state.State))
	require.NoError(t, s.SetInt(5))
	require.Equal(t, state.Small, s.Get())
}

func TestSetAverage(t *testing.T) {
	s := syncstate.New(new(state.State))
	require.NoError(t, s.SetInt(15))
	require.Equal(t, state.Average, s.Get())
}

func TestSetGreat(t *testing.T) {
	s := syncstate.New(new(state.State))
	require.NoError(t, s.SetInt(25))
	require.Equal(t, state.Great, s.Get())
}

func TestSetUnsupported(t *testing.T) {
	s := syncstate.New(new(state.State))
	previousState := s.Get()
	require.Error(t, s.SetInt(11))
	require.Equal(t, previousState, s.Get())
}

package day11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStoneZeroChangeToOneOnBlink(t *testing.T) {
	var stone = Stone{engravedNumber: 0}

	var stonesAfterBlink = stone.OnBlink()

	assert.Len(t, stonesAfterBlink, 1)
	assert.Equal(t, Stone{engravedNumber: 1}, stonesAfterBlink[0])
}



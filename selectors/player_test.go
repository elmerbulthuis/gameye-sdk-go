package selectors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectPlayerList(t *testing.T) {
	playerList := SelectPlayerList(statisticStateMock)
	assert.Equal(t, 2, len(playerList))
	for _, playerItem := range playerList {
		switch playerItem.PlayerKey {
		case "3":
			assert.Equal(t, "denise", playerItem.Name)
		case "4":
			assert.Equal(t, "Smashmint", playerItem.Name)
		default:
			assert.Fail(t, playerItem.PlayerKey)
		}
	}
}

func TestSelectPlayerListForTeam(t *testing.T) {
	playerList := SelectPlayerListForTeam(statisticStateMock, "1")
	assert.Equal(t, 1, len(playerList))
	for _, playerItem := range playerList {
		switch playerItem.PlayerKey {
		case "3":
			assert.Equal(t, "denise", playerItem.Name)
		default:
			assert.Fail(t, playerItem.PlayerKey)
		}
	}
}

func TestSelectPlayerItem(t *testing.T) {
	playerItem := SelectPlayerItem(statisticStateMock, "4")
	assert.NotNil(t, playerItem)
	if playerItem != nil {
		assert.Equal(t, "4", playerItem.PlayerKey)
		assert.Equal(t, "Smashmint", playerItem.Name)
	}
}

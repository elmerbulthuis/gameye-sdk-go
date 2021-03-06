package selectors

import "github.com/Gameye/gameye-sdk-go/models"

var gameStateMock = &models.GameQueryState{
	Game: models.GameQueryGameIndex{
		"csgo": &models.GameQueryGameItem{GameKey: "csgo", Location: map[string]bool{}},
		"tf2":  &models.GameQueryGameItem{GameKey: "tf2", Location: map[string]bool{}},
		"css":  &models.GameQueryGameItem{GameKey: "css", Location: map[string]bool{}},
		"l4d2": &models.GameQueryGameItem{GameKey: "l4d2", Location: map[string]bool{}},
		"kf2":  &models.GameQueryGameItem{GameKey: "kf2", Location: map[string]bool{}},
		"test": &models.GameQueryGameItem{GameKey: "test", Location: map[string]bool{"local": true}},
	},
	Location: models.GameQueryLocationIndex{
		"rotterdam":     &models.GameQueryLocationItem{LocationKey: "rotterdam"},
		"ireland":       &models.GameQueryLocationItem{LocationKey: "ireland"},
		"dubai":         &models.GameQueryLocationItem{LocationKey: "dubai"},
		"tokyo":         &models.GameQueryLocationItem{LocationKey: "tokyo"},
		"los_angeles":   &models.GameQueryLocationItem{LocationKey: "los_angeles"},
		"washington_dc": &models.GameQueryLocationItem{LocationKey: "washington_dc"},
		"local":         &models.GameQueryLocationItem{LocationKey: "local"},
	},
}

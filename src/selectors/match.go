package selectors

import "github.com/Gameye/gameye-sdk-go/src/models"

type MatchItem = models.MatchQueryMatchItem

/**
 * Select a list of active matches.
 * @param matchState match state
 */
func SelectMatchList(
	matchState *models.MatchQueryState,
) (
	matchList []*MatchItem,
) {
	matchList = make([]*MatchItem, 0)

	for _, matchItem := range matchState.Match {
		if matchItem == nil {
			continue
		}

		matchList = append(matchList, matchItem)
	}

	return
}

/**
 * Select a list of active matches for a game.
 * @param matchState match state
 * @param gameKey identifier of the game to select matches for
 */
func SelectMatchListForGame(
	matchState *models.MatchQueryState,
	gameKey string,
) (
	matchList []*MatchItem,
) {
	matchList = make([]*MatchItem, 0)

	for _, matchItem := range matchState.Match {
		if matchItem == nil {
			continue
		}
		if matchItem.GameKey != gameKey {
			continue
		}

		matchList = append(matchList, matchItem)
	}

	return
}

/**
 * Get details about a single match from a match-state as returned by
 * the gameye api.
 * @param matchState match state
 * @param matchKey identifier of the match to get the details for
 */
func SelectMatchItem(
	matchState *models.MatchQueryState,
	matchKey string,
) (
	matchItem *MatchItem,
) {
	matchItem = matchState.Match[matchKey]
	return
}

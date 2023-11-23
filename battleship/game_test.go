package battleship

import (
	"reflect"
	"testing"
)

func TestNewGame(t *testing.T) {
	testSize := 5
	game := NewGame(testSize)

	if len(game.Players) != testSize {
		t.Errorf("Received %v, expected %v", len(game.Players), testSize)
	}
}

func TestSetUpPlayer(t *testing.T) {
	tests := map[string]struct {
		playerNumber   int
		gridSize       int
		totalShips     int
		shipPositions  []Coordinates
		totalMissiles  int
		missileMoves   []Coordinates
		expectedPlayer Player
	}{
		"test set 1": {
			playerNumber: 0,
			gridSize:     10,
			totalShips:   1,
			shipPositions: []Coordinates{
				Coordinates{
					X: 0,
					Y: 0,
				},
			},
			totalMissiles: 1,
			missileMoves: []Coordinates{
				Coordinates{
					X: 0,
					Y: 0,
				},
			},
			expectedPlayer: Player{
				Board:          NewBoard(10),
				TotalMissilies: 1,
				TotalShips:     1,
				ShipCoordinates: []Coordinates{
					Coordinates{
						X: 0,
						Y: 0,
					},
				},
				HitCoordinates: []Coordinates{
					Coordinates{
						X: 0,
						Y: 0,
					},
				},
				TotalPoints: 0,
			},
		},
		"test set 2": {
			playerNumber: 1,
			gridSize:     5,
			totalShips:   2,
			shipPositions: []Coordinates{
				Coordinates{
					X: 0,
					Y: 0,
				},
				Coordinates{
					X: 1,
					Y: 1,
				},
			},
			totalMissiles: 2,
			missileMoves: []Coordinates{
				Coordinates{
					X: 0,
					Y: 0,
				},
				Coordinates{
					X: 1,
					Y: 1,
				},
			},
			expectedPlayer: Player{
				Board:          NewBoard(5),
				TotalMissilies: 2,
				TotalShips:     2,
				ShipCoordinates: []Coordinates{
					Coordinates{
						X: 0,
						Y: 0,
					},
					Coordinates{
						X: 1,
						Y: 1,
					},
				},
				HitCoordinates: []Coordinates{
					Coordinates{
						X: 0,
						Y: 0,
					},
					Coordinates{
						X: 1,
						Y: 1,
					},
				},
				TotalPoints: 0,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			testSize := 5
			game := NewGame(testSize)
			game.SetUpPlayer(test.playerNumber, test.gridSize, test.totalShips, test.totalMissiles, test.shipPositions, test.missileMoves)

			player := game.Players[test.playerNumber]
			if reflect.DeepEqual(player, test.playerNumber) {
				t.Errorf("Received %v, expected %v", player, test.expectedPlayer)
			}
		})
	}
}

func TestFireMissiles(t *testing.T) {
	var (
		gridSize      int = 5
		totalShips    int = 5
		totalMissiles int = 2
	)
	missileMoves := []Coordinates{
		Coordinates{
			X: 0,
			Y: 0,
		},
		Coordinates{
			X: 1,
			Y: 1,
		},
	}
	//p2MissileMoves  []battleship.Coordinates
	game := NewGame(2)
	game.SetUpPlayer(0, gridSize, totalShips, totalMissiles, []Coordinates{}, missileMoves)
	game.SetUpPlayer(1, gridSize, totalShips, totalMissiles, []Coordinates{}, []Coordinates{})

	// Player 0 fires at player 1.
	game.FireMissiles(0, 1)

	opponent := game.Players[1]
	// We check the coordinated on player 1's board.
	for _, coordinate := range missileMoves {
		if opponent.Board.Board[coordinate.X][coordinate.Y] != MARK_HIT && opponent.Board.Board[coordinate.X][coordinate.Y] != MARK_MISS {
			t.Errorf("Test failed: Expected a hit or miss marked here")
		}
	}

}

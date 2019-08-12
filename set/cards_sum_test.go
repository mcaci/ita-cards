package set

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
)

func assertScoreCount(t *testing.T, expectedScore uint8, cards Cards) {
	if score := cards.Sum(func(card.Item) uint8 { return 1 }); expectedScore != score {
		t.Fatalf("Score expected is %d but %d was computed", expectedScore, score)
	}
}

func TestEmptyPileSums0(t *testing.T) {
	assertScoreCount(t, 0, nil)
}

func TestPileWithOneCardSums1(t *testing.T) {
	assertScoreCount(t, 1, Cards{*card.MustID(1)})
}

func TestPileWithThreeCardsSums3(t *testing.T) {
	assertScoreCount(t, 3, Cards{*card.MustID(1), *card.MustID(2), *card.MustID(3)})
}

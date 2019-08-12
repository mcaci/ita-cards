package set

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
)

func TestCreateSet(t *testing.T) {
	c, _ := card.New("5", "Cup")
	set := Cards{*c}
	if index := set.Find(*c); index == -1 {
		t.Fatalf("There should be the 5 of Cup card in the set")
	}
}

func TestAddCardToSet(t *testing.T) {
	c, _ := card.New("3", "Cudgel")
	set := Cards{*c}
	if index := set.Find(*c); index == -1 {
		t.Fatal("There should be the 3 of Cudgel card in the set")
	}
}

func TestClearRemovesAllCards(t *testing.T) {
	c, _ := card.New("2", "Coin")
	set := Cards{*c}
	set.Clear()
	if len(set) != 0 {
		t.Fatalf("Cards were not removed from set")
	}
}

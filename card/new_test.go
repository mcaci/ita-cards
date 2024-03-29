package card_test

import (
	"strconv"
	"testing"

	"github.com/mcaci/ita-cards/card"
)

type cardCreationCheck struct {
	msg  string
	errF func(err error) bool
	chkF func(card *card.Item, number, seed string) bool
}

var noErrCheck = cardCreationCheck{
	msg:  "An unexpected error was raised",
	errF: func(err error) bool { return err == nil },
}
var errCheck = cardCreationCheck{
	msg:  "The %q of %q is not a valid card",
	errF: func(err error) bool { return err != nil },
}
var cCheck = cardCreationCheck{
	msg: "Card's number is not created well from %q and %q",
	chkF: func(card *card.Item, number, seed string) bool {
		return strconv.Itoa(int(card.Number())) == number && card.Seed().String() == string(seed)
	},
}

func TestCardCreation(t *testing.T) {
	testcases := []struct {
		name   string
		number string
		seed   string
		check  cardCreationCheck
	}{
		{"1OfCoinIsCreatedCorrectly_NoError", "1", "Coin", noErrCheck},
		{"1OfCoinIsCreatedCorrectly_NumberIs1", "1", "Coin", cCheck},
		{"Test1OfCoinIsCreatedCorrectly_SeedIsCoin", "1", "Coin", cCheck},
		{"Test2OfSwordIsCreatedCorrectly_NoError", "2", "Sword", noErrCheck},
		{"Test2OfSwordIsCreatedCorrectly_NumberIs2", "2", "Sword", cCheck},
		{"Test2OfSwordIsCreatedCorrectly_SeedIsSword", "2", "Sword", cCheck},
		{"Test8OfCupIsCreatedCorrectly_NoError", "8", "Cup", noErrCheck},
		{"Test8OfCupIsCreatedCorrectly_NumberIs8", "8", "Cup", cCheck},
		{"Test8OfCupIsCreatedCorrectly_SeedIsCup", "8", "Cup", cCheck},
		{"Test10OfCudgelIsCreatedCorrectly_NoError", "10", "Cudgel", noErrCheck},
		{"Test10OfCudgelIsCreatedCorrectly_NumberIs10", "10", "Cudgel", cCheck},
		{"Test10OfCudgelIsCreatedCorrectly_SeedIsCudgel", "10", "Cudgel", cCheck},
		{"Test15OfCupDoesntExist", "15", "Cup", errCheck},
		{"Test8OfSpadesDoesntExist", "8", "Spades", errCheck},
		{"TestTwoOfCudgelIsIncorrect", "Two", "Cudgel", errCheck},
		{"TestEmptyNumberIsIncorrect", "", "Cudgel", errCheck},
		{"TestEmptySeedIsIncorrect", "6", "", errCheck},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			item, err := card.New(tc.number, tc.seed)
			if tc.check.errF != nil && !tc.check.errF(err) {
				t.Fatalf(tc.check.msg, tc.number, tc.seed)
			}
			if tc.check.chkF != nil && !tc.check.chkF(item, tc.number, tc.seed) {
				t.Fatalf(tc.check.msg, tc.number, tc.seed, item.Seed().String())
			}
		})
	}
}

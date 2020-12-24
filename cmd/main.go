package main

import (
	"encoding/json"
	"log"
	"net/http"

	poker "github.com/zshearin/poker/cmd/poker"
)

func main() {

	http.HandleFunc("/", pokerHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func pokerHandler(w http.ResponseWriter, r *http.Request) {
	deal1 := shuffleAndDeal(4)

	//	deal1.PrintBoardAndHands()
	//	deal1.Players.Print()

	type result struct {
		Board       poker.Cards
		Hands       poker.Hands
		HandResults []poker.HandResult
	}
	res := result{
		Board:       deal1.GetBoard(),
		Hands:       deal1.Hands,
		HandResults: deal1.HandResults,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	//Now format the body with all the objects we want

}

func shuffleAndDeal(players int) poker.Deal {

	deck := poker.GetDeck()
	deck.Shuffle()
	deck.Shuffle()
	//deck.PrintOrder()
	return deck.GetDeal(players)
}

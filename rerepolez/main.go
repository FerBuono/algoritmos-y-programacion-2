package main

import (
	"fmt"
	"os"
	"rerepolez/errors"
	"rerepolez/libs/models"
	"rerepolez/libs/utils"
)

func main() {
	var args = os.Args[1:]
	if len(args) < 2 {
		fmt.Fprintln(os.Stdout, errors.ErrorParams{}.Error())
		os.Exit(0)
	}

	parties := utils.OpenFile(args[0])
	partyList := utils.SaveParties(parties)
	parties.Close()

	voterList := utils.OpenFile(args[1])
	voterArray := utils.SortVoterList(utils.SaveVoterList(voterList))
	voterList.Close()

	blankVotes := models.CreateBlankVotes()
	var impugnedVotes *int = new(int)

	utils.Voting(voterArray, partyList, blankVotes, impugnedVotes)
	utils.Finalize(blankVotes, partyList, impugnedVotes)
}

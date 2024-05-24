package models

import "fmt"

type partyImplementation struct {
	name       string
	candidates [3]string
	votes      [3]int
}

type blankParty struct {
	votes [3]int
}

func CreateParty(name string, candidates [3]string) Party {
	party := new(partyImplementation)
	party.name = name
	party.candidates = candidates
	return party
}

func CreateBlankVotes() Party {
	blank := new(blankParty)
	return blank
}

func (party *partyImplementation) VotedFor(position int) {
	party.votes[position]++
}

func (party partyImplementation) GetResult(position int) string {
	if party.votes[position] == 1 {
		return fmt.Sprintf("%s - %s: %d voto", party.name, party.candidates[position], party.votes[position])
	}
	return fmt.Sprintf("%s - %s: %d votos", party.name, party.candidates[position], party.votes[position])
}

func (blank *blankParty) VotedFor(position int) {
	blank.votes[position]++
}

func (blank blankParty) GetResult(position int) string {
	if blank.votes[position] == 1 {
		return fmt.Sprintf("Votos en Blanco: %d voto", blank.votes[position])
	}
	return fmt.Sprintf("Votos en Blanco: %d votos", blank.votes[position])
}

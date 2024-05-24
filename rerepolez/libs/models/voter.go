package models

import (
	"rerepolez/errors"

	dynamic_stack "github.com/FerBuono/go-data-structures/dynamic-stack"
)

const (
	VOTING = true
)

type vote struct {
	position    int
	alternative int
}

type voterImplementation struct {
	dni    int
	votes  dynamic_stack.Stack[vote]
	voting bool
}

func CreateVoter(dni int) Voter {
	v := new(voterImplementation)
	v.dni = dni
	v.votes = dynamic_stack.NewDynamicStack[vote]()
	v.voting = VOTING
	return v
}

func (voter voterImplementation) ReadDNI() int {
	return voter.dni
}

func (voter *voterImplementation) Vote(position int, alternative int) error {
	if !voter.voting {
		return voter.voterFraud()
	}
	vote := vote{position, alternative}
	voter.votes.Push(vote)
	return nil
}

func (voter *voterImplementation) Undo() error {
	if !voter.voting {
		return voter.voterFraud()
	}
	if voter.votes.IsEmpty() {
		newError := new(errors.ErrorNoPreviousVotes)
		return newError
	}
	voter.votes.Pop()
	return nil
}

func (voter *voterImplementation) EndVote() (Vote, error) {
	if !voter.voting {
		return Vote{}, voter.voterFraud()
	}

	finalVote := [3]int{BLANK_VOTE, BLANK_VOTE, BLANK_VOTE}

	for !voter.votes.IsEmpty() {
		vote := voter.votes.Pop()

		if finalVote[vote.position] == BLANK_VOTE || vote.alternative == IMPUGNED_LIST {
			finalVote[vote.position] = vote.alternative
		}
	}
	voter.voting = false
	for _, vote := range finalVote {
		if vote == IMPUGNED_LIST {
			return Vote{finalVote, true}, nil
		}
	}
	return Vote{finalVote, false}, nil
}

func (voter *voterImplementation) voterFraud() error {
	newError := new(errors.ErrorVoterFraud)
	newError.Dni = voter.dni
	return newError
}

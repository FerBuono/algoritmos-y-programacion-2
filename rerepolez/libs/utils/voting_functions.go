package utils

import (
	"bufio"
	"fmt"
	"os"
	"rerepolez/errors"
	"rerepolez/libs/models"
	"strconv"
	"strings"

	linked_queue "github.com/FerBuono/go-data-structures/linked-queue"
)

var positions = [3]string{"Presidente", "Gobernador", "Intendente"}

func Voting(voters []models.Voter, parties []models.Party, blankVotes models.Party, impugnedVotes *int) {
	voterQueue := linked_queue.NewLinkedQueue[models.Voter]()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		input := strings.Split(scanner.Text(), " ")
		action := input[0]

		switch action {

		case "ingresar":
			pos, err := enterVoter(input, voters)

			if err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
			} else {
				voterQueue.Enqueue(voters[pos])
				fmt.Println("OK")
			}

		case "votar":
			err := vote(input, voterQueue, parties)

			if err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
			} else {
				fmt.Println("OK")
			}

		case "deshacer":
			err := undoVote(voterQueue)

			if err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
			} else {
				fmt.Println("OK")
			}

		case "fin-votar":
			vote, err := finalizeVote(voterQueue)

			if err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
			} else {
				voterQueue.Dequeue()
				fmt.Println("OK")
				distributeVotes(vote, impugnedVotes, blankVotes, parties)
			}

		default:
			fmt.Fprintln(os.Stdout, "Input incorrecto")
		}
	}

	if !voterQueue.IsEmpty() {
		fmt.Fprintln(os.Stdout, errors.ErrorVotersWithoutVote{}.Error())
	}
}

func enterVoter(input []string, voters []models.Voter) (int, error) {
	dni, err := strconv.Atoi(input[1])

	if err != nil || dni < 0 {
		return -1, errors.DNIError{}
	}

	pos := findVoter(voters, 0, len(voters)-1, dni)
	if pos == -1 {
		return -1, errors.DNIOutOfRegister{}
	}

	return pos, nil
}

func vote(input []string, voterQueue linked_queue.Queue[models.Voter], parties []models.Party) error {
	if voterQueue.IsEmpty() {
		return errors.EmptyQueue{}
	} else if len(input) < 3 {
		return errors.ErrorVoteType{}
	}

	position := input[1]
	if position != positions[0] && position != positions[1] && position != positions[2] {
		return errors.ErrorVoteType{}
	}

	list, err := strconv.Atoi(input[2])
	numParties := len(parties)
	if err != nil || list > numParties {
		return errors.ErrorInvalidAlternative{}
	}

	var voterFraudError error
	for i, pos := range positions {
		if position == pos {
			voterFraudError = voterQueue.Peek().Vote(i, list)
		}
	}

	if voterFraudError != nil {
		voterQueue.Dequeue()
		return voterFraudError
	}

	return nil
}

func undoVote(voterQueue linked_queue.Queue[models.Voter]) error {
	if voterQueue.IsEmpty() {
		return errors.EmptyQueue{}
	}

	newError := voterQueue.Peek().Undo()
	if newError != nil {
		if newError.Error() != "ERROR: Sin voto a deshacer" {
			voterQueue.Dequeue()
		}
		return newError
	}

	return nil
}

func finalizeVote(voterQueue linked_queue.Queue[models.Voter]) (models.Vote, error) {
	if voterQueue.IsEmpty() {
		return models.Vote{}, errors.EmptyQueue{}
	}

	finalVote, newError := voterQueue.Peek().EndVote()

	if newError != nil {
		return models.Vote{}, newError
	}

	return finalVote, nil
}

func distributeVotes(vote models.Vote, impugnedVotes *int, blankVotes models.Party, parties []models.Party) {
	if vote.Impugned {
		*impugnedVotes++
		return
	}
	for i, vote := range vote.VotesByType {
		if vote == models.BLANK_VOTE {
			blankVotes.VotedFor(i)
		} else {
			parties[vote-1].VotedFor(i)
		}
	}
}

func Finalize(blankVotes models.Party, parties []models.Party, impugnedVotes *int) {
	for i, pos := range positions {
		if i == 0 {
			fmt.Fprintf(os.Stdout, "%s:\n", pos)
		} else {
			fmt.Fprintf(os.Stdout, "\n%s:\n", pos)
		}
		fmt.Fprintln(os.Stdout, blankVotes.GetResult(i))
		for _, party := range parties {
			fmt.Fprintln(os.Stdout, party.GetResult(i))
		}
	}

	if *impugnedVotes == 1 {
		fmt.Fprintf(os.Stdout, "\nVotos Impugnados: %d voto\n", *impugnedVotes)
	} else {
		fmt.Fprintf(os.Stdout, "\nVotos Impugnados: %d votos\n", *impugnedVotes)
	}
}

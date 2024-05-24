package models

const (
	PRESIDENT int = 0
	GOVERNOR  int = 1
	MAYOR     int = 2
)

const (
	NUM_POSITIONS     = MAYOR + 1
	IMPUGNED_LIST int = 0
	BLANK_VOTE    int = -1
)

// Vote stores the information of a cast vote for each possible type of vote.
// For example, the position GOVERNOR will have the chosen alternative for Governor.
// If it is 0, it is a blank vote.
// If Impugned is 'true', then none of the chosen alternatives should be considered.
type Vote struct {
	VotesByType [int(NUM_POSITIONS)]int
	Impugned    bool
}

// Voter models a voter in our voting system.
type Voter interface {
	// ReadDNI gives us the DNI of the voter.
	ReadDNI() int

	// Vote registers the chosen alternative for the specified position. If the voter has already finished voting,
	// it will return the corresponding error. Otherwise, nil.
	Vote(position int, alternative int) error

	// Undo undoes the last action performed. It should be possible to undo until the initial state of the vote
	// (equivalent to a completely blank vote). If there has been an action to undo, it will return nil. If there is no action
	// to undo, it will return the corresponding error. It can also return an error if the voter has already finished
	// their voting process.
	Undo() error

	// EndVote completes the voting process for this voter. If the voter has already finished voting,
	// it will return the corresponding error. Otherwise, it will return the final vote state obtained from the various
	// applications of Vote and Undo.
	EndVote() (Vote, error)
}

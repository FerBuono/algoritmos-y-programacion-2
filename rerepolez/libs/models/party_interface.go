package models

// Party models a political party with its candidates for each type of election.
type Party interface {
	// VotedFor indicates that this Party has received a vote for the specified position.
	VotedFor(position int)

	// GetResult retrieves the result for this Party for the specified position. The format will be suitable for display.
	GetResult(position int) string
}

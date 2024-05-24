package utils

import (
	"bufio"
	"fmt"
	"os"
	"rerepolez/errors"
	"rerepolez/libs/models"
	"strconv"
	"strings"
)

const (
	_NUM_DIGITS = 10
	_DNI_DIGITS = 8
)

func OpenFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		newError := new(errors.ErrorReadingFile)
		fmt.Fprintln(os.Stdout, newError.Error())
		os.Exit(0)
	}
	return file
}

func SaveParties(parties *os.File) []models.Party {
	partyList := []models.Party{}

	scanner := bufio.NewScanner(parties)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		candidates := [3]string{data[1], data[2], data[3]}
		party := models.CreateParty(data[0], candidates)
		partyList = append(partyList, party)
	}
	return partyList
}

func SaveVoterList(voterList *os.File) []models.Voter {
	voters := []models.Voter{}

	scanner := bufio.NewScanner(voterList)
	for scanner.Scan() {
		dni, _ := strconv.Atoi(scanner.Text())
		voter := models.CreateVoter(dni)
		voters = append(voters, voter)
	}
	return voters
}

func intPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func countByCriterion(voters []models.Voter, criterion func(int) int) []models.Voter {
	frequencies := make([]int, _NUM_DIGITS)
	for _, voter := range voters {
		frequencies[criterion(voter.ReadDNI())]++
	}

	freqSum := make([]int, _NUM_DIGITS)
	sum := 0
	for i := range freqSum {
		freqSum[i] += sum
		sum += frequencies[i]
	}

	sortedArray := make([]models.Voter, len(voters))
	for _, voter := range voters {
		sortedArray[freqSum[criterion(voter.ReadDNI())]] = voter
		freqSum[criterion(voter.ReadDNI())]++
	}

	return sortedArray
}

func SortVoterList(voters []models.Voter) []models.Voter {
	sortedArray := voters
	for i := 0; i < _DNI_DIGITS; i++ {
		sortedArray = countByCriterion(sortedArray, func(num int) int { return (num / intPow(10, i)) % 10 })
	}

	return sortedArray
}

func findVoter(voters []models.Voter, start, end, element int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if voters[mid].ReadDNI() == element {
		return mid
	}
	if voters[mid].ReadDNI() < element {
		return findVoter(voters, mid+1, end, element)
	} else {
		return findVoter(voters, start, mid-1, element)
	}
}

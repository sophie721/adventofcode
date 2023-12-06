package _1202

import (
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func Run01(input []string) int {
	constraint := game{sets: map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}}

	wg := new(sync.WaitGroup)
	wg.Add(len(input))
	lock := new(sync.Mutex)
	result := 0

	for _, s := range input {
		checkOneGame(s, constraint, &result, lock, wg)
	}

	wg.Wait()
	return result
}

func checkOneGame(s string, constraint game, result *int, lock *sync.Mutex, wg *sync.WaitGroup) {
	info := getGameInfo(s)
	if check(constraint, info) {
		lock.Lock()
		*result += info.id
		lock.Unlock()
	}
	wg.Done()
}

func check(constraint game, info gameInfo) bool {

	for _, set := range info.sets {
		for color, count := range set {
			if count > constraint.sets[color] {
				return false
			}
		}
	}
	return true
}

type gameInfo struct {
	id   int
	sets []map[string]int
}

type game struct {
	sets gameSets
}

type gameSets map[string]int

func getGameInfo(s string) gameInfo {
	r, _ := regexp.Compile("^Game (\\d+): (.+)$")
	matches := r.FindStringSubmatch(s)
	gameId, _ := strconv.Atoi(matches[1])
	info := gameInfo{
		id:   gameId,
		sets: []map[string]int{},
	}

	colorInfoRegex, _ := regexp.Compile("^(\\d+) (\\w+)$")
	for _, setInfo := range strings.Split(matches[2], "; ") {
		for _, colorInfo := range strings.Split(setInfo, ", ") {
			matches := colorInfoRegex.FindStringSubmatch(colorInfo)
			colorCount, _ := strconv.Atoi(matches[1])
			info.sets = append(info.sets, map[string]int{
				matches[2]: colorCount,
			})
		}
	}

	return info
}

package _1202

import "sync"

func Run02(input []string) int {
	wg := new(sync.WaitGroup)
	wg.Add(len(input))
	lock := new(sync.Mutex)
	result := 0

	for _, s := range input {
		sumGamePower(s, &result, lock, wg)
	}

	wg.Wait()
	return result
}

func sumGamePower(s string, result *int, lock *sync.Mutex, wg *sync.WaitGroup) {
	info := getGameInfo(s)
	game := getMinGameSet(info)
	mul := 1
	for _, count := range game {
		mul *= count
	}

	lock.Lock()
	*result += mul
	lock.Unlock()
	wg.Done()
}

func getMinGameSet(info gameInfo) gameSets {
	min := gameSets{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, set := range info.sets {
		for color, count := range set {
			min[color] = max(min[color], count)
		}
	}
	return min
}

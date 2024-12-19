package days

import (
	"adventofcode/src/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Day11 struct {
	utils.DayN
	stones []int
	maxIt  int
	memory sync.Map
	wg     sync.WaitGroup
	score  *atomic.Int64
}

var _ utils.Day = (*Day11)(nil)

func (d *Day11) Preprocess(path string) error {
	d.maxIt = 75 // Increased iterations
	d.score = &atomic.Int64{}
	d.wg = sync.WaitGroup{}
	d.memory = sync.Map{}
	d.score.Store(0)
	line := utils.ReadFileContent(path)
	strVals := strings.Split(line, " ")
	for i := 0; i < len(strVals); i++ {
		v, _ := strconv.Atoi(strVals[i])
		d.stones = append(d.stones, v)
	}
	return nil
}

func (d *Day11) Solve(path string) {
	startTime := time.Now()

	fmt.Printf("Stones: %v\n", d.stones)

	const batchSize = 100
	numWorkers := 8
	ch := make(chan int, batchSize)
	var workerWg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		workerWg.Add(1)
		go func() {
			defer workerWg.Done()
			for stone := range ch {
				d.processIteration(stone)
			}
		}()
	}

	for i := 0; i < len(d.stones); i++ {
		ch <- d.stones[i]
	}

	close(ch)
	workerWg.Wait()

	elapsedTime := time.Since(startTime)

	fmt.Printf("Elapsed Time: %s\n", elapsedTime)

	d.Pt1Sol = int(d.score.Load())
}

func (d *Day11) processIteration(stone int) {
	stoneCounts := make(map[int]int)
	stoneCounts[stone] = 1
	for it := 0; it < d.maxIt; it++ {
		nextStoneCounts := make(map[int]int)

		for s, count := range stoneCounts {
			s1, s2, useSecond := d.getStones(s)

			if !useSecond {
				nextStoneCounts[s1] += count
			} else {
				nextStoneCounts[s1] += count
				nextStoneCounts[s2] += count
			}
		}

		stoneCounts = nextStoneCounts
	}

	totalStones := 0
	for _, count := range stoneCounts {
		totalStones += count
	}

	d.score.Add(int64(totalStones))
}

func (d *Day11) getStones(stone int) (s1, s2 int, useSecond bool) {
	if stone == 0 {
		return 1, 0, false
	}
	digits := int(math.Log10(float64(stone))) + 1
	if digits%2 != 0 {
		return stone * 2024, 0, false
	}

	divisor := int(math.Pow(10, float64(digits/2)))
	return stone / divisor, stone % divisor, true
}

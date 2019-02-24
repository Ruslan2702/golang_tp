package main

import (
	"runtime"
	"sort"
	"strconv"
	"sync"
)

var mu = &sync.Mutex{}
var list = &sync.Mutex{}


func SingleHash(in, out chan interface{}) {
	shGroup := &sync.WaitGroup{}
	for input := range in {
		data := strconv.Itoa(input.(int))

		shGroup.Add(1)
		go SingleHashCalculate(data, out, shGroup)
		runtime.Gosched()
	}

	shGroup.Wait()
}

func SingleMd5(result chan interface{}, data string, compareGroup *sync.WaitGroup) {
	defer compareGroup.Done()

	mu.Lock()
	result <- DataSignerMd5(data)
	mu.Unlock()
}

func SingleCrc32(result chan interface{}, data string, compareGroup *sync.WaitGroup) {
	defer compareGroup.Done()

	result <- DataSignerCrc32(data)
}

func SingleHashCalculate(data string, out chan interface{}, singleWait *sync.WaitGroup) {
	defer singleWait.Done()

	compareGroup := &sync.WaitGroup{}

	val1 := make(chan interface{})
	compareGroup.Add(1)
	go SingleMd5(val1, data, compareGroup)

	val2 := make(chan interface{})
	compareGroup.Add(1)
	go SingleCrc32(val2, data, compareGroup)

	val3 := make(chan interface{})
	compareGroup.Add(1)
	go SingleCrc32(val3, (<-val1).(string), compareGroup)

	out <- (<-val2).(string) + "~" + (<-val3).(string)
}

func MultiHash(in, out chan interface{}) {

	multiGroup := &sync.WaitGroup{}

	for input := range in {
		data := (input).(string)

		multiGroup.Add(1)
		go MultiHashCalculate(data, out, multiGroup)
	}

	multiGroup.Wait()
}

func MultiHashCalculate(data string, out chan interface{}, multiGroup *sync.WaitGroup) {
	defer multiGroup.Done()

	sixItGroup := &sync.WaitGroup{}

	chansForChanks := make([]chan interface{}, 6)
	for i := 0; i < 6; i++ {
		chansForChanks[i] = make(chan interface{})
	}

	for i := 0; i < 6; i++ {
		sixItGroup.Add(1)
		go MultiCrc32(chansForChanks[i], data, i, sixItGroup)
	}

	resultHash := ""

	for i := 0; i < 6; i++ {
		resultHash += (<-chansForChanks[i]).(string)
	}

	out <- resultHash
}

func MultiCrc32(result chan interface{}, data string, iteration int, group *sync.WaitGroup) {
	defer group.Done()

	result <- DataSignerCrc32(strconv.Itoa(iteration) + data)
}

func CombineResults(in, out chan interface{}) {

	multiResults := make([]string, 0)

	for input := range in {
		list.Lock()
		multiResults =  append(multiResults, input.(string))
		list.Unlock()
	}

	sort.Strings(multiResults)

	result := ""
	for idx, data := range multiResults {
		result += data
		if idx != len(multiResults) - 1 {
			result += "_"
		}
	}

	out <- result
}

func ExecutePipeline(tasks ...job) {
	out := make(chan interface{})
	in := make(chan interface{})

	wg := &sync.WaitGroup{}

	for _, task := range tasks {
		in = out
		out = make(chan interface{})

		wg.Add(1)
		go DoTask(task, in, out, wg)
	}

	wg.Wait()
}

func DoTask(task job, in chan interface{}, out chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)

	task(in, out)
}

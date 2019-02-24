package main

import (
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var mu = &sync.Mutex{}


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

func SingleMd5(result chan string, data string) {
	mu.Lock()
	result <- DataSignerMd5(data)
	mu.Unlock()
}

func SingleCrc32(result chan string, data string) {
	result <- DataSignerCrc32(data)
}

func SingleHashCalculate(data string, out chan interface{}, singleWait *sync.WaitGroup) {
	defer singleWait.Done()

	val1 := make(chan string)
	go SingleMd5(val1, data)

	val2 := make(chan string)
	go SingleCrc32(val2, data)

	val3 := make(chan string)
	go SingleCrc32(val3, <-val1)

	out <- (<-val2) + "~" + (<-val3)
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

	chansForChanks := make([]chan string, 6)
	for i := 0; i < 6; i++ {
		chansForChanks[i] = make(chan string)
	}

	for i := 0; i < 6; i++ {
		sixItGroup.Add(1)
		go MultiCrc32(chansForChanks[i], data, i, sixItGroup)
	}

	resultHash := ""

	for i := 0; i < 6; i++ {
		resultHash += <-chansForChanks[i]
	}

	out <- resultHash
}

func MultiCrc32(result chan string, data string, iteration int, group *sync.WaitGroup) {
	defer group.Done()

	result <- DataSignerCrc32(strconv.Itoa(iteration) + data)
}

func CombineResults(in, out chan interface{}) {

	multiResults := make([]string, 0)

	for input := range in {
		multiResults =  append(multiResults, input.(string))
	}

	sort.Strings(multiResults)
	result := strings.Join(multiResults, "_")

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

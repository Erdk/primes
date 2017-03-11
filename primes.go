package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"sync"
)

func main() {
	fd, err := os.Create("primes.out")
	defer fd.Close()

	index := uint64(1)

	if err != nil {
		panic(err)
	}

	for potentialPrime := uint64(1); potentialPrime < 100000000; potentialPrime++ {
		// validate input, also return for first 10 values
		switch potentialPrime {
		case 1:
			{
				fd.Write([]byte(fmt.Sprintf("%v\t%v\n", index, potentialPrime)))
				index++
				continue
			}
		case 2:
			{
				fd.Write([]byte(fmt.Sprintf("%v\t%v\n", index, potentialPrime)))
				index++
				continue
			}
		case 3:
			{
				fd.Write([]byte(fmt.Sprintf("%v\t%v\n", index, potentialPrime)))
				index++
				continue
			}
		case 4:
			{
				continue
			}
		case 5:
			{
				fd.Write([]byte(fmt.Sprintf("%v\t%v\n", index, potentialPrime)))
				index++
				continue
			}
		case 6:
			{
				continue
			}
		case 7:
			{
				fd.Write([]byte(fmt.Sprintf("%v\t%v\n", index, potentialPrime)))
				index++
				continue
			}
		case 8:
			{
				continue
			}
		case 9:
			{
				continue
			}
		case 10:
			{
				continue
			}
		}

		boundary := uint64(math.Sqrt(float64(potentialPrime)))
		// check if at the end is 0, 2, 4, 6, 8
		lastDigit := potentialPrime % 10
		if lastDigit == 2 || lastDigit == 4 || lastDigit == 6 || lastDigit == 8 || lastDigit == 0 || lastDigit == 5 {
			continue
		}

		// check rest
		if potentialPrime < 257 {
			isPrime := true
			for i := uint64(3); i < boundary; i += 2 {
				if potentialPrime%i == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				fd.Write([]byte(fmt.Sprintf("%v\t%v\n", index, potentialPrime)))
				index++
			}
		} else {
			numThreads := uint64(runtime.NumCPU())

			var wg sync.WaitGroup
			wg.Add(int(numThreads))

			isPrime := true

			f := func(thidx int, start, end uint64, cancel chan bool) {
				defer wg.Done()
				if start%2 == 0 {
					start++
				}
				for i := start; i < end; i += 2 {
					select {
					case <-cancel:
						return
					default:
					}
					if potentialPrime%i == 0 {
						isPrime = false
						for j := 0; j < runtime.NumCPU(); j++ {
							cancel <- true
						}
						return
					}
				}
			}

			cancel := make(chan bool, runtime.NumCPU()*runtime.NumCPU())

			stripe := (boundary - uint64(2)) / numThreads
			for i := uint64(0); i < numThreads; i++ {
				go f(int(i), 2+i*stripe, 2+(i+1)*stripe, cancel)
			}
			wg.Wait()
			close(cancel)
			if isPrime {
				fd.Write([]byte(fmt.Sprintf("%v\t%v\n", index, potentialPrime)))
				index++
			}
		}
	}
}

package hw06pipelineexecution

import (

	"sync"
)

var mutex = &sync.Mutex{}

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func execStages(data interface{}, stages ...Stage) Out {
	out := make(Bi)

	go func() {
		fIn := make(Bi)
		if len(stages) > 0 {
			fOut := stages[0](fIn)
			for i := 1; i < len(stages); i++ {
				fOut = stages[i](fOut)
			}
			fIn <- data
			out <- <-fOut
		} else {
			out <- data
		}


		close(out)
	}()
	return out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := make(Bi)

	go func() {
		values := make(map[int]Out)
		var wg sync.WaitGroup
		i := -1
		for data := range in {
			wg.Add(1)
			i++
			go func(values map[int]Out, index int, data interface{}, stages ...Stage) {
				defer wg.Done()
				value := execStages(data, stages...)
				mutex.Lock()
				values[index] = value
				mutex.Unlock()
			}(values, i, data, stages...)
		}
		wg.Wait()

		for j := 0; j <= i; j++ {
			v := <-values[j]
			out <- v
		}

		close(out)
	}()
	return out
}

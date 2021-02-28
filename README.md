<img src="https://github.com/tejzpr/ordered-concurrently/actions/workflows/tests.yml/badge.svg" />

# Ordered Concurrently
A go module that processes work concurrently and returns output in a channel with a predefined order

# Usage 
## Import package
```go
import concurrently "github.com/tejzpr/ordered-concurrently" 
```
## Create a work function
```go
// The work that needs to be performed
func workFn(val interface{}) interface{} {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	return val
}
```
## Run
```go
func main() {
	max := 100
	inputChan := make(chan *concurrently.OrderedInput)
	wg := &sync.WaitGroup{}
	go func() {
		outChan := concurrently.Process(inputChan, workFn, 10)
		for out := range outChan {
			fmt.Println(out.Value)
			wg.Done()
		}
	}()

	// Create work and the associated order
	for work, order := 0, 0; work < max; work, order = work+1, order+1 {
		wg.Add(1)
		input := &concurrently.OrderedInput{work, order}
		inputChan <- input
	}
	wg.Wait()
}

```
# Credits
Thanks to [u/justinisrael](https://www.reddit.com/user/justinisrael/) for inputs on improving resource usage.

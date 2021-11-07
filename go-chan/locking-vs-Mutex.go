'''
Using locking (mutexes)#
Use mutex when:

Caching information in a shared data structure
Holding state information, i.e., the context or status of a running application

Using channels#
Use channels when:

Communicating asynchronous results
Distributing units of work
Passing ownership of data
'''

// Using Mutex
type Pool struct {
  Mu sync.Mutex
  Tasks []Task
}

func Worker(pool *Pool) {
  for {
    pool.Mu.Lock()
    // begin critical section:
    task := pool.Tasks[0] // take the first task
    pool.Tasks = pool.Tasks[1:] // update the pool of tasks
    // end critical section
    pool.Mu.Unlock()
    process(task)
  }
}

//Using Channels

func main() {
  pending, done := make(chan *Task), make(chan *Task)
  go sendWork(pending) // put tasks to do on the channel
  for i := 0; i < N; i++ { // start N goroutines to do work
    go Worker(pending, done)
  }
  consumeWork(done) // continue with the processed tasks
}

func Worker(in, out chan *Task) {
  for {
    t := <-in
    process(t)
    out <- t
  }
}

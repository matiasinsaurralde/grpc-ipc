package grpc_ipc

import (
	"log"
	"math/rand"
	"time"
)

// Pool holds the "worker" grpc servers
type Pool struct {
	Workers []*Worker
	Cmd     []string
}

// NewPool initializes a new pool of workers.
func NewPool(poolSize int, poolCmd []string, portRange []int) *Pool {
	log.Println("Initializing pool")

	rand.Seed(time.Now().Unix())

	pool := &Pool{Cmd: poolCmd}
	pool.Workers = make([]*Worker, poolSize)
	for i := 0; i < poolSize; i++ {
		pool.Workers[i] = &Worker{pool: pool}
	}
	return pool
}

// Start will start all the workers.
func (p *Pool) Start() {
	log.Println("Starting pool", p)

	for index, worker := range p.Workers {
		log.Println("  Starting worker:", worker, index)
		worker.Prepare()
		worker.Start()
	}
}

// PickWorker will return a random worker from the pool.
func (p *Pool) PickWorker() *Worker {
	index := rand.Int() % len(p.Workers)
	log.Println("Pick worker #", index)
	return p.Workers[index]
}

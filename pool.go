package main

import(
  "log"
)

// Pool holds the "worker" grpc servers
type Pool struct {
  Workers []*Worker
}

func NewPool(poolSize int, poolCmd []string, portRange []int) *Pool {
  log.Println("Initializing pool")
  pool := &Pool{}
  pool.Workers = make([]*Worker, poolSize)
  for i := 0; i < poolSize; i++ {
    pool.Workers[i] = &Worker{ pool: pool }
  }
  return pool
}

func(p *Pool) Start() {
  log.Println("Starting pool", p)

  for index, worker := range p.Workers {
    log.Println("  Starting worker:", worker, index)
  }
}

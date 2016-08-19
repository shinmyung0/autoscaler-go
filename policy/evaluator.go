package policy

import (
	"sync"
	"time"

	log "github.com/shinmyung0/loglite"
)

type EvaluationState int

const (
	Stable EvaluationState = 1
	ScalingOut
	ScalingIn
	Observing
)

type Evaluator struct {
	sync.RWMutex
	evaluations map[string]*EvaluationJob
}

type EvaluationJob struct {
	Timer         *time.Ticker
	Policy        Policy
	CurrentState  EvaluationState
	PreviousState EvaluationState
}

var evaluator *Evaluator

func init() {
	evaluator = &Evaluator{evaluations: make(map[string]*EvaluationJob)}
}

func GetEvaluator() *Evaluator {
	return evaluator
}

func (e *Evaluator) AddEvaluation(p Policy) {
	log.Debug("Adding evaluation for ", p.Service)
	e.Lock()
	j := e.initializeEvaluation(p)
	e.Unlock()

	j.run()
}

func (e *Evaluator) RemoveEvaluation(serviceName string) {
	log.Debug("Removing evaluation job for ", serviceName)
	e.Lock()
	job, ok := e.evaluations[serviceName]
	if ok {
		job.stop()
	}
	delete(e.evaluations, serviceName)
	e.Unlock()
}

func (e *Evaluator) initializeEvaluation(p Policy) *EvaluationJob {
	ticker := time.NewTicker(time.Second * time.Duration(p.EvaluationInterval))
	job := &EvaluationJob{ticker, p, Stable, Stable}
	e.evaluations[p.Service] = job
	return job

}

func (j *EvaluationJob) run() {
	go func() {
		for t := range j.Timer.C {
			log.Debug("Running collection for", j.Policy.Service, t)
			collectingPlaceholder(j.Policy)
			log.Debugf("%s Finished!", j.Policy.Service)
		}
	}()
}

func (j *EvaluationJob) stop() {
	j.Timer.Stop()
}

func collectingPlaceholder(p Policy) {
	time.Sleep(time.Second * 1)
}

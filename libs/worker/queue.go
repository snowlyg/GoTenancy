package worker

// Queue is an interface defined methods need for a job queue
type Queue interface {
	Add(QorJobInterface) error
	Run(QorJobInterface) error
	Kill(QorJobInterface) error
	Remove(QorJobInterface) error
}

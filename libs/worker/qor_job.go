package worker

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/audited"
	"GoTenancy/libs/serializable_meta"
	"github.com/jinzhu/gorm"
)

// QorJobInterface is a interface, defined methods that needs for a qor job
type QorJobInterface interface {
	GetJobID() string
	GetJobName() string
	GetStatus() string
	SetStatus(string) error
	GetJob() *Job
	SetJob(*Job)

	GetProgress() uint
	SetProgress(uint) error
	GetProgressText() string
	SetProgressText(string) error
	GetLogs() []string
	AddLog(string) error
	GetResultsTable() ResultsTable
	AddResultsRow(...TableCell) error

	StartReferesh()
	StopReferesh()

	GetArgument() interface{}
	serializable_meta.SerializableMetaInterface
}

// ResultsTable is a struct, including importing/exporting results
type ResultsTable struct {
	Name       string `json:"-"` // only used for generate string column in database
	TableCells [][]TableCell
}

// Scan used to scan value from database into itself
func (resultsTable *ResultsTable) Scan(data interface{}) error {
	switch values := data.(type) {
	case []byte:
		return json.Unmarshal(values, resultsTable)
	case string:
		return resultsTable.Scan([]byte(values))
	default:
		return errors.New("unsupported data type for Qor Job error table")
	}
}

// Value used to read value from itself and save it into databae
func (resultsTable ResultsTable) Value() (driver.Value, error) {
	result, err := json.Marshal(resultsTable)
	return string(result), err
}

// TableCell including Value, Error for a data cell
type TableCell struct {
	Value string
	Error string
}

// QorJob predefined qor job struct, which will be used for Worker, if it doesn't include a job resource
type QorJob struct {
	gorm.Model
	Status       string `sql:"default:'new'"`
	Progress     uint
	ProgressText string
	Log          string       `sql:"size:65532"`
	ResultsTable ResultsTable `sql:"size:65532"`

	mutex sync.Mutex `sql:"-"`

	stopReferesh bool `sql:"-"`
	inReferesh   bool `sql:"-"`

	// Add `valid:"-"`` to make the QorJob work well with qor/validations
	// When the qor/validations auto exec the validate struct callback we get error
	// runtime: goroutine stack exceeds 1000000000-byte limit
	// fatal error: stack overflow
	Job *Job `sql:"-" valid:"-"`

	audited.AuditedModel
	serializable_meta.SerializableMeta
}

// GetJobID get job's ID from a qor job
func (job *QorJob) GetJobID() string {
	return fmt.Sprint(job.ID)
}

// GetJobName get job's name from a qor job
func (job *QorJob) GetJobName() string {
	return job.Kind
}

// GetStatus get job's status from a qor job
func (job *QorJob) GetStatus() string {
	return job.Status
}

// SetStatus set job's status to a qor job instance
func (job *QorJob) SetStatus(status string) error {
	job.mutex.Lock()
	defer job.mutex.Unlock()

	job.Status = status
	if status == JobStatusDone {
		job.Progress = 100
	}

	if job.shouldCallSave() {
		return job.callSave()
	}

	return nil
}

func (job *QorJob) shouldCallSave() bool {
	return !job.inReferesh || job.stopReferesh
}

func (job *QorJob) StartReferesh() {
	job.mutex.Lock()
	defer job.mutex.Unlock()
	if !job.inReferesh {
		job.inReferesh = true
		job.stopReferesh = false

		go func() {
			job.referesh()
		}()
	}
}

func (job *QorJob) StopReferesh() {
	job.mutex.Lock()
	defer job.mutex.Unlock()

	err := job.callSave()
	if err != nil {
		log.Println(err)
	}

	job.stopReferesh = true
}

func (job *QorJob) referesh() {
	job.mutex.Lock()
	defer job.mutex.Unlock()

	err := job.callSave()
	if err != nil {
		log.Println(err)
	}

	if job.stopReferesh {
		job.inReferesh = false
		job.stopReferesh = false
	} else {
		time.AfterFunc(5*time.Second, job.referesh)
	}
}

func (job *QorJob) callSave() error {
	worker := job.GetJob().Worker
	context := worker.Admin.NewContext(nil, nil).Context
	return worker.JobResource.CallSave(job, context)
}

// SetJob set `Job` for a qor job instance
func (job *QorJob) SetJob(j *Job) {
	job.Kind = j.Name
	job.Job = j
}

// GetJob get predefined job for a qor job instance
func (job *QorJob) GetJob() *Job {
	if job.Job != nil {
		return job.Job
	}
	return nil
}

// GetArgument get job's argument
func (job *QorJob) GetArgument() interface{} {
	return job.GetSerializableArgument(job)
}

// GetSerializableArgumentResource get job's argument's resource
func (job *QorJob) GetSerializableArgumentResource() *admin.Resource {
	if j := job.GetJob(); j != nil {
		return j.Resource
	}
	return nil
}

// GetProgress get qor job's progress
func (job *QorJob) GetProgress() uint {
	return job.Progress
}

// SetProgress set qor job's progress
func (job *QorJob) SetProgress(progress uint) error {
	job.mutex.Lock()
	defer job.mutex.Unlock()

	if progress > 100 {
		progress = 100
	}
	job.Progress = progress

	if job.shouldCallSave() {
		return job.callSave()
	}

	return nil
}

// GetProgressText get qor job's progress text
func (job *QorJob) GetProgressText() string {
	return job.ProgressText
}

// SetProgressText set qor job's progress text
func (job *QorJob) SetProgressText(str string) error {
	job.mutex.Lock()
	defer job.mutex.Unlock()

	job.ProgressText = str
	if job.shouldCallSave() {
		return job.callSave()
	}

	return nil
}

// GetLogs get qor job's logs
func (job *QorJob) GetLogs() []string {
	return strings.Split(job.Log, "\n")
}

// AddLog add a log to qor job
func (job *QorJob) AddLog(log string) error {
	job.mutex.Lock()
	defer job.mutex.Unlock()

	fmt.Println(log)
	job.Log += "\n" + log
	if job.shouldCallSave() {
		return job.callSave()
	}

	return nil
}

// GetResultsTable get the job's process logs
func (job *QorJob) GetResultsTable() ResultsTable {
	return job.ResultsTable
}

// AddResultsRow add a row of process results to a job
func (job *QorJob) AddResultsRow(cells ...TableCell) error {
	job.mutex.Lock()
	defer job.mutex.Unlock()

	job.ResultsTable.TableCells = append(job.ResultsTable.TableCells, cells)
	if job.shouldCallSave() {
		return job.callSave()
	}

	return nil
}

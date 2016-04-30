package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Shyp/rickover/Godeps/_workspace/src/github.com/Shyp/go-types"
)

// A Job is an in-memory representation of a record in the jobs table.
//
// Once you create a Job, you can enqueue new jobs using the Job name.
type Job struct {
	Name             string           `json:"name"`
	DeliveryStrategy DeliveryStrategy `json:"delivery_strategy"`
	Attempts         uint8            `json:"attempts"`
	Concurrency      uint8            `json:"concurrency"`
	CreatedAt        time.Time        `json:"created_at"`
}

// A QueuedJob is a job to be run at a point in the future.
//
// QueuedJobs can have the status "queued" (to be run at some point), or
// "in-progress" (a dequeuer is acting on them).
type QueuedJob struct {
	Id        types.PrefixUUID `json:"id"`
	Name      string           `json:"name"`
	Attempts  uint8            `json:"attempts"`
	RunAfter  time.Time        `json:"run_after"`
	ExpiresAt types.NullTime   `json:"expires_at"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	Status    JobStatus        `json:"status"`
	Data      json.RawMessage  `json:"data"`
}

// DeliveryStrategy describes how a job should be run. If it's safe to run a
// job more than once (updating a cache), use StrategyAtLeastOnce for your Job.
// If it's not safe to run a job more than once (sending an email or SMS), use
// StrategyAtMostOnce.
type DeliveryStrategy string

// StrategyAtLeastOnce should be used for jobs that can be retried in the event
// of failure.
const StrategyAtLeastOnce = DeliveryStrategy("at_least_once")

// StrategyAtMostOnce should be used for jobs that should not be retried in
// the event of failure.
const StrategyAtMostOnce = DeliveryStrategy("at_most_once")

// Value implements the driver.Valuer interface.
func (d DeliveryStrategy) Value() (driver.Value, error) {
	return string(d), nil
}

type JobStatus string

// StatusQueued indicates a QueuedJob is scheduled to be run at some point in
// the future.
const StatusQueued = JobStatus("queued")

// StatusInProgress indicates a QueuedJob has been dequeued, and is being
// worked on.
const StatusInProgress = JobStatus("in-progress")

// StatusSucceeded indicates a job has been completed successfully and then
// archived.
const StatusSucceeded = JobStatus("succeeded")

// StatusFailed indicates the job completed, but an error occurred.
const StatusFailed = JobStatus("failed")

// StatusExpired indicates the job was dequeued after its ExpiresAt date.
const StatusExpired = JobStatus("expired")

// Scan implements the Scanner interface.
func (j *JobStatus) Scan(src interface{}) error {
	if src == nil {
		return nil
	} else if txt, ok := src.(string); ok {
		*j = JobStatus(txt)
		return nil
	} else if txt, ok := src.([]byte); ok {
		*j = JobStatus(string(txt))
		return nil
	}
	return fmt.Errorf("Unsupported JobStatus: %#v", src)
}

// Value implements the driver.Valuer interface.
func (j JobStatus) Value() (driver.Value, error) {
	return string(j), nil
}

// Scan implements the Scanner interface.
func (d *DeliveryStrategy) Scan(src interface{}) error {
	if src == nil {
		return nil
	} else if txt, ok := src.(string); ok {
		*d = DeliveryStrategy(txt)
		return nil
	} else if txt, ok := src.([]byte); ok {
		*d = DeliveryStrategy(string(txt))
		return nil
	}
	return fmt.Errorf("Unsupported DeliveryStrategy: %#v", src)
}

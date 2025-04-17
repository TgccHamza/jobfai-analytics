package services

import (
	"os"
	"sync"
	"time"

	"github.com/hibiken/asynq"
)

// TaskService handles task queue operations
type TaskService struct {
	client *asynq.Client
}

var (
	taskServiceInstance *TaskService
	taskServiceOnce     sync.Once
)

// NewTaskService creates a new TaskService instance
func NewTaskService() *TaskService {
	taskServiceOnce.Do(func() {
		// Get Redis address from environment variable with fallback
		redisAddr := os.Getenv("REDIS_ADDR")
		if redisAddr == "" {
			redisAddr = "127.0.0.1:6379"
		}

		// Create Redis client options
		redisOpt := asynq.RedisClientOpt{
			Addr: redisAddr,
		}

		// Optional: Add Redis password if provided
		redisPassword := os.Getenv("REDIS_PASSWORD")
		if redisPassword != "" {
			redisOpt.Password = redisPassword
		}

		// Create the client
		client := asynq.NewClient(redisOpt)
		
		taskServiceInstance = &TaskService{
			client: client,
		}
	})

	return taskServiceInstance
}

// EnqueueTask enqueues a task to be processed immediately
func (s *TaskService) EnqueueTask(task *asynq.Task) (*asynq.TaskInfo, error) {
	return s.client.Enqueue(task)
}

// EnqueueTaskWithDelay enqueues a task to be processed after the specified delay
func (s *TaskService) EnqueueTaskWithDelay(task *asynq.Task, delay time.Duration) (*asynq.TaskInfo, error) {
	return s.client.Enqueue(task, asynq.ProcessIn(delay))
}

// EnqueueTaskWithOptions enqueues a task with the specified options
func (s *TaskService) EnqueueTaskWithOptions(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	return s.client.Enqueue(task, opts...)
}

// Close closes the underlying Redis connection
func (s *TaskService) Close() error {
	return s.client.Close()
}
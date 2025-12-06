package pve

import (
	"fmt"
)

// TasksService handles task-related API operations
type TasksService struct {
	client *Client
}

// List retrieves all tasks
func (s *TasksService) List(options *TaskListOptions) ([]*Task, error) {
	req, err := s.client.NewRequest("GET", "cluster/tasks", options)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []*Task
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetTask retrieves a specific task by UPID
func (s *TasksService) GetTask(upid string) (*Task, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("cluster/tasks/%s", upid), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data *Task
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// StopTask stops a running task
func (s *TasksService) StopTask(upid string) error {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("cluster/tasks/%s/stop", upid), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// StopNodeTask stops a task on a specific node
func (s *TasksService) StopNodeTask(nodeName, upid string) error {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/tasks/%s/stop", nodeName, upid), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// GetTaskLog retrieves task log
func (s *TasksService) GetTaskLog(upid string) ([]string, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("cluster/tasks/%s/log", upid), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []string
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetTaskLogWithPaging retrieves task log with paging
func (s *TasksService) GetTaskLogWithPaging(upid string, start, limit int) ([]string, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("cluster/tasks/%s/log", upid), map[string]any{
		"start": start,
		"limit": limit,
	})
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []string
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetNodeTasks retrieves tasks for a specific node
func (s *NodesService) GetNodeTasks(name string, options *TaskListOptions) ([]*Task, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/tasks", name), options)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []*Task
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetNodeTaskLog retrieves task log for a node task
func (s *TasksService) GetNodeTaskLog(nodeName, upid string) ([]string, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/tasks/%s/log", nodeName, upid), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []string
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetNodeTaskStatus retrieves task status for a node task
func (s *TasksService) GetNodeTaskStatus(nodeName, upid string) (map[string]any, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/tasks/%s/status", nodeName, upid), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data map[string]any
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// WaitForTask waits for a task to complete
func (s *TasksService) WaitForTask(upid string, timeout int) (map[string]any, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("cluster/tasks/%s/status", upid), map[string]any{
		"timeout": timeout,
	})
	if err != nil {
		return nil, err
	}

	var result struct {
		Data map[string]any
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

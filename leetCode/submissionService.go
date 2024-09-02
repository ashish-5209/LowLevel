package main

import (
	"errors"
	"time"
)

type Submission struct {
	ID        string
	UserID    string
	ProblemID string
	Code      string
	Language  string
	Status    string
	CreatedAt time.Time
}

type SubmissionService struct {
	submission map[string]*Submission
	judge      *JudgeService
}

func NewSubmissionService(judge *JudgeService) *SubmissionService {
	return &SubmissionService{
		submission: make(map[string]*Submission),
		judge:      judge,
	}
}

func (s *SubmissionService) SubmitCode(userID, problemID, code, language string) (*Submission, error) {
	id := generateID()
	submission := &Submission{
		ID:        id,
		UserID:    userID,
		ProblemID: problemID,
		Code:      code,
		Language:  language,
		Status:    "Pending",
		CreatedAt: time.Now(),
	}
	s.submission[id] = submission

	go func() {
		status, err := s.judge.EvaluateSubmission(submission)
		if err == nil {
			submission.Status = status
		} else {
			submission.Status = "Error"
		}
	}()
	return submission, nil
}

func (s *SubmissionService) GetSubmission(submissionID string) (*Submission, error) {
	submission, exists := s.submission[submissionID]
	if !exists {
		return nil, errors.New("submission not found")
	}
	return submission, nil
}

func (s *SubmissionService) ListSubmissionByUser(userID string) ([]*Submission, error) {
	submissions := []*Submission{}
	for _, submission := range s.submission {
		if submission.UserID == userID {
			submissions = append(submissions, submission)
		}
	}

	return submissions, nil
}

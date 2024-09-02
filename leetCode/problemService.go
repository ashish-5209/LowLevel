package main

import (
	"errors"
	"time"
)

type TestCase struct {
	Input  string
	Output string
}

type Problem struct {
	Id          string
	Title       string
	Description string
	Difficulty  string
	TestCases   []TestCase
	CreatedAt   time.Time
}

type ProblemService struct {
	problems map[string]*Problem
}

func NewProblemService() *ProblemService {
	return &ProblemService{problems: make(map[string]*Problem)}
}

func (s *ProblemService) CreateProblem(title, description, difficulty string, testCases []TestCase) (*Problem, error) {
	id := generateID()
	if _, exists := s.problems[id]; exists {
		return nil, errors.New("problem already exixts")
	}

	problem := &Problem{
		Id:          id,
		Title:       title,
		Description: description,
		Difficulty:  difficulty,
		TestCases:   testCases,
		CreatedAt:   time.Now(),
	}
	s.problems[id] = problem
	return problem, nil
}

func (s *ProblemService) GetProblem(problemID string) (*Problem, error) {
	problem, exists := s.problems[problemID]
	if !exists {
		return nil, errors.New("problem not found")
	}
	return problem, nil
}

func (s *ProblemService) ListProblem() ([]*Problem, error) {
	problems := []*Problem{}
	for _, problem := range s.problems {
		problems = append(problems, problem)
	}
	return problems, nil
}

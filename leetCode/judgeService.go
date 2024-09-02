package main

import (
	"errors"
	"fmt"
)

type JudgeService struct {
	problemService *ProblemService
}

func NewJudgeService(problemService *ProblemService) *JudgeService {
	return &JudgeService{
		problemService: problemService,
	}
}

func (j *JudgeService) EvaluateSubmission(submission *Submission) (string, error) {
	problem, err := j.problemService.GetProblem(submission.ProblemID)
	if err != nil {
		return "", errors.New("problem not found")
	}

	for _, testcase := range problem.TestCases {
		output, err := runCode(submission.Code, testcase.Input, submission.Language)
		if err != nil || output != testcase.Output {
			return "Failed", nil
		}
	}
	return "Passed", nil
}

func runCode(code, input, language string) (string, error) {
	fmt.Printf("Running code: %s with input: %s\n", code, input)
	return input, nil
}

// +----------------------------------+
// |            UserService           |
// +----------------------------------+
// | - CreateUser(username, email,    |
// |   password string): *User, error |
// | - AuthenticateUser(username,     |
// |   password string): string, error|
// | - GetUser(userID string): *User, |
// |   error                          |
// +----------------------------------+

// +----------------------------------+
// |            ProblemService        |
// +----------------------------------+
// | - CreateProblem(title, desc,     |
// |   difficulty string,             |
// |   testCases []TestCase):         |
// |   *Problem, error                |
// | - GetProblem(problemID string):  |
// |   *Problem, error                |
// | - ListProblems(): []Problem,     |
// |   error                          |
// +----------------------------------+

// +----------------------------------+
// |           SubmissionService      |
// +----------------------------------+
// | - SubmitCode(userID, problemID,  |
// |   code, language string):        |
// |   *Submission, error             |
// | - GetSubmission(submissionID     |
// |   string): *Submission, error    |
// | - ListSubmissionsByUser(userID   |
// |   string): []Submission, error   |
// +----------------------------------+

// +----------------------------------+
// |            JudgeService          |
// +----------------------------------+
// | - EvaluateSubmission(submissionID|
// |   string): string, error         |
// +----------------------------------+

// +----------------------------------+
// |         LeaderboardService       |
// +----------------------------------+
// | - GetLeaderboard(): []Leaderboard|
// | - UpdateLeaderboard(userID string|
// |   score int): error              |
// +----------------------------------+

// +----------------------------------+
// |            DiscussionForum       |
// +----------------------------------+
// | - PostComment(userID, problemID, |
// |   comment string): error         |
// | - GetComments(problemID string): |
// |   []Comment, error               |
// +----------------------------------+

// +-----------------+       +-----------------+
// |      User       |       |    Problem      |
// +-----------------+       +-----------------+
// | - ID            |       | - ID            |
// | - Username      |       | - Title         |
// | - Email         |       | - Description   |
// | - Password      |       | - Difficulty    |
// | - CreatedAt     |       | - TestCases     |
// +-----------------+       +-----------------+

// +-----------------+       +-----------------+
// |   Submission    |       |   TestCase      |
// +-----------------+       +-----------------+
// | - ID            |       | - Input         |
// | - UserID        |       | - Output        |
// | - ProblemID     |       |                 |
// | - Code          |       +-----------------+
// | - Language      |
// | - Status        |       +-----------------+
// | - CreatedAt     |       |  Leaderboard    |
// +-----------------+       +-----------------+
//                             | - UserID       |
//                             | - Score        |
//                             | - Rank         |
//                             +----------------+

package main

import (
	"fmt"
	"time"
)

func main() {
	userService := NewUserService()
	problemService := NewProblemService()
	judgeService := NewJudgeService(problemService)
	submissionService := NewSubmissionService(judgeService)
	discussionService := NewDiscussionForum()
	leaderBoardService := NewLeaderBoardService()

	user, _ := userService.CreateUser("john_doe", "john@example.com", "password123")
	fmt.Printf("User Created: %+v\n", user)

	testCases := []TestCase{
		{Input: "2\n3", Output: "5"},
		{Input: "4\n6", Output: "10"},
	}

	problem, _ := problemService.CreateProblem("Add Two Numbers", "Given two numbers, return their sum.", "Easy", testCases)
	fmt.Printf("Problem Created: %+v\n", problem)

	submission, _ := submissionService.SubmitCode(user.Id, problem.Id, "sum = a + b", "python")
	fmt.Printf("Submission Created: %+v\n", submission)

	time.Sleep(1 * time.Second)

	updatedSubmission, _ := submissionService.GetSubmission(submission.ID)
	fmt.Printf("Submission Status: %s\n", updatedSubmission.Status)

	if updatedSubmission.Status == "Passed" {
		leaderBoardService.UpdateLeaderBoard(user.Id, 10)
	}

	comment, _ := discussionService.GetComments(problem.Id)
	fmt.Printf("Comments for Problem: %+v\n", comment)

	leaderBoard := leaderBoardService.GetLeaderBoard()
	fmt.Printf("Leaderboard: %+v\n", leaderBoard)

}

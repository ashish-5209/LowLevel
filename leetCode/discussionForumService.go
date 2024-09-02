package main

import (
	"errors"
	"time"
)

type Comment struct {
	ID        string
	UserID    string
	ProblemID string
	Content   string
	CreatedAt time.Time
}

type DiscussionForum struct {
	comments map[string][]*Comment
}

func NewDiscussionForum() *DiscussionForum {
	return &DiscussionForum{
		comments: make(map[string][]*Comment),
	}
}

func (d *DiscussionForum) PostComment(userID, problemID, content string) (*Comment, error) {
	id := generateID()
	comment := &Comment{
		ID:        id,
		UserID:    userID,
		ProblemID: problemID,
		Content:   content,
		CreatedAt: time.Now(),
	}
	d.comments[problemID] = append(d.comments[problemID], comment)
	return comment, nil
}

func (d *DiscussionForum) GetComments(ProblemID string) ([]*Comment, error) {
	comments, exists := d.comments[ProblemID]
	if !exists {
		return nil, errors.New("no comments found for this problem")
	}
	return comments, nil
}

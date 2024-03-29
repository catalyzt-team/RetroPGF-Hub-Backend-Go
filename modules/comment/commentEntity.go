package comment

import (
	"RetroPGF-Hub/RetroPGF-Hub-Backend-Go/modules/users"
	"time"
)

type (
	PushCommentReq struct {
		Title     string    `json:"title" validate:"required"`
		Content   string    `json:"content" validate:"required"`
		CreatedBy string    `json:"created_by"`
		CreateAt  time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	CommentRes struct {
		ProjectId string        `json:"projectId"`
		Comments  []CommentARes `json:"comments"`
		CreateAt  time.Time     `json:"createdAt"`
		UpdatedAt time.Time     `json:"updatedAt"`
	}

	CommentARes struct {
		CommentId string    `json:"commentId"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		CreatedBy string    `json:"createdBy"`
		CreateAt  time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	CommentAResWithUser struct {
		CommentId string               `json:"commentId"`
		Title     string               `json:"title"`
		Content   string               `json:"content"`
		CreatedBy users.UserProfileRes `json:"createdBy"`
		CreateAt  time.Time            `json:"createdAt"`
		UpdatedAt time.Time            `json:"updatedAt"`
	}
)

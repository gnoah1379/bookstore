package model

import "time"

type Review struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	BookID     int       `json:"book_id"`
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment"`
	CountReply int       `json:"count_reply"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (r Review) tableName() string {
	return "reviews"
}

type ReplyReview struct {
	ID        int       `json:"id"`
	ReviewID  int       `json:"review_id"`
	UserID    int       `json:"user_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r ReplyReview) tableName() string {
	return "reply_reviews"
}

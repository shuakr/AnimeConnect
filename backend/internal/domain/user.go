package domain

import "time"

type User struct {
	ID        string    `json:"id" firestore:"id"`
	Username  string    `json:"username" firestore:"username"`
	Email     string    `json:"email" firestore:"email"`
	Password  string    `json:"password,omitempty" firestore:"password"`
	AvatarURL string    `json:"avatarUrl,omitempty" firestore:"avatarUrl"`
	CreatedAt time.Time `json:"createdAt" firestore:"createdAt"`
}

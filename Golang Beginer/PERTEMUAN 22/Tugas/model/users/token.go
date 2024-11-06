package UserModel

import "time"

type Token struct {
	Token_id   int
	User_id    int
	Token      string
	Expired_at time.Time
}

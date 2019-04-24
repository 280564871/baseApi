package moudel

import "time"

type User struct {
	username     string `json:"username"`
	password     string `json:"password"`
	open_id      string `json:"open_id"`
	nickname     string `json:"nickname"`
	is_subscribe int    `js`
	push_time    int
	created_at   time.Time
	updated_at   time.Time
}

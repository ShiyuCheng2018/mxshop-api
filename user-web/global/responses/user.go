package responses

import (
	"fmt"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(j).Format("2022-01-01"))
	return []byte(stamp), nil
}

type UserInfoResponse struct {
	Id       int64    `json:"id"`
	Name     string   `json:"name"`
	Gender   string   `json:"gender"`
	Birthday JsonTime `json:"birthday"`
	Nickname string   `json:"nickname"`
	Mobile   string   `json:"mobile"`
	Role     int32    `json:"role"`
}

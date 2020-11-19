package util

import uuid "github.com/satori/go.uuid"

func NewTraceID() string {
	u := uuid.NewV4()
	return u.String()
}

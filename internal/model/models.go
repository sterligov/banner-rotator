package model

import (
	"fmt"
	"time"
)

const (
	EventClick = iota
	EventSelect
)

var (
	ErrNotFound = fmt.Errorf("entity not found")
)

type (
	Banner struct {
		ID          int64
		Description string
	}

	Group struct {
		ID          int64
		Description string
	}

	Slot struct {
		ID          int64
		Description string
	}

	Event struct {
		Type     byte      `json:"type"`
		SlotID   int64     `json:"slot_id"`
		BannerID int64     `json:"banner_id"`
		GroupID  int64     `json:"group_id"`
		Date     time.Time `json:"date"`
	}
)

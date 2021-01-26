package model

import "time"

const (
	Click = iota
	Select
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

	Event struct {
		Type          byte      `json:"type"`
		SlotID        int64     `json:"slot_id"`
		BannerID      int64     `json:"banner_id"`
		SocialGroupID int64     `json:"social_group_id"`
		Date          time.Time `json:"date"`
	}
)

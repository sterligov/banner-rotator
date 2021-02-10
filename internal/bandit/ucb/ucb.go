package ucb

import (
	"math"

	"github.com/sterligov/banner-rotator/internal/model"
)

type UCB struct{}

func New() *UCB {
	return &UCB{}
}

func (u *UCB) SelectBanner(stats []model.Statistic) int64 {
	var totalShows int64

	for _, s := range stats {
		if s.Shows == 0 {
			return s.BannerID
		}

		totalShows += s.Shows
	}

	var (
		maxWeight float64
		bannerID  int64
	)

	for _, s := range stats {
		weight := float64(s.Clicks)/float64(s.Shows) + math.Sqrt(2*math.Log(float64(totalShows))/float64(s.Shows))
		if weight > maxWeight {
			maxWeight = weight
			bannerID = s.BannerID
		}
	}

	return bannerID
}

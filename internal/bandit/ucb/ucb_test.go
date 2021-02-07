package ucb

import (
	"testing"

	"github.com/sterligov/banner-rotator/internal/model"
	"github.com/stretchr/testify/require"
)

func TestSelectBanner(t *testing.T) {
	t.Run("select banner with 0 shows", func(t *testing.T) {
		var expectedBannerID int64 = 2
		stats := []model.Statistic{
			{
				BannerID: 1,
				SlotID:   1,
				GroupID:  1,
				Clicks:   10,
				Shows:    10,
			},
			{
				BannerID: expectedBannerID,
				SlotID:   1,
				GroupID:  1,
				Clicks:   1,
				Shows:    0,
			},
		}

		ucb := New()
		actualBannerID := ucb.SelectBanner(stats)
		require.Equal(t, expectedBannerID, actualBannerID)
	})

	t.Run("select banner with max weight", func(t *testing.T) {
		var expectedBannerID int64 = 2
		stats := []model.Statistic{
			{
				BannerID: 1,
				SlotID:   1,
				GroupID:  1,
				Clicks:   19,
				Shows:    45,
			},
			{
				BannerID: expectedBannerID,
				SlotID:   1,
				GroupID:  1,
				Clicks:   30,
				Shows:    40,
			},
			{
				BannerID: 3,
				SlotID:   1,
				GroupID:  1,
				Clicks:   7,
				Shows:    50,
			},
		}

		ucb := New()
		actualBannerID := ucb.SelectBanner(stats)
		require.Equal(t, expectedBannerID, actualBannerID)
	})
}

package dto

import "time"

type FarmAnalyticsQuery struct {
	StartDate   *string `form:"start_date" binding:"omitempty,datetime=2006-01-02"`
	EndDate     *string `form:"end_date" binding:"omitempty,datetime=2006-01-02"`
	SectorID    *uint   `form:"sector_id" binding:"omitempty,gt=0"`
	Aggregation string  `form:"aggregation" binding:"omitempty,oneof=daily weekly monthly"`
}

func (query *FarmAnalyticsQuery) SetDefaults() {
	if query.Aggregation == "" {
		query.Aggregation = "daily"
	}

	if query.SectorID == nil {
		query.SectorID = new(uint)
	}

	now := time.Now().Format("2006-01-02")

	if query.StartDate == nil {
		query.StartDate = &now
	}

	if query.EndDate == nil {
		query.EndDate = &now
	}
}

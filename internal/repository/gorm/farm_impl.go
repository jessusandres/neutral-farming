package gorm

import (
	"time"

	"gorm.io/gorm"

	"looker.com/neutral-farming/internal/domain"
	"looker.com/neutral-farming/internal/domain/read_models"
	"looker.com/neutral-farming/internal/model"
	"looker.com/neutral-farming/internal/repository"
)

type FarmRepo struct {
	db *gorm.DB
}

func NewFarmRepo(db *gorm.DB) repository.FarmRepository {
	return &FarmRepo{db}
}

func (repo *FarmRepo) FindByID(ID uint) (*domain.Farm, error) {

	domainInstance := &model.Farm{}
	result := repo.db.First(domainInstance, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &domain.Farm{
		ID:   domainInstance.ID,
		Name: domainInstance.Name,
	}, nil
}

func (repo *FarmRepo) YearOverYearAnalytics(farmID uint, sectorID uint, startDate, endDate time.Time) ([]*read_models.IrrigationAnalytics, error) {
	var results []*read_models.IrrigationAnalytics

	query := `
WITH periods AS (
    SELECT ?::TIMESTAMPTZ                        AS c_start,
           ?::TIMESTAMPTZ                          AS c_end,
           (?::TIMESTAMPTZ - INTERVAL '1 year')  AS p1_start,
           (?::TIMESTAMPTZ - INTERVAL '1 year')    AS p1_end,
           (?::TIMESTAMPTZ - INTERVAL '2 years') AS p2_start,
           (?::TIMESTAMPTZ - INTERVAL '2 years')   AS p2_end
),
filtered_data AS (
    SELECT id,
           nominal_amount,
           real_amount,
           CASE
               WHEN start_time BETWEEN (SELECT c_start FROM periods) AND (SELECT c_end FROM periods) THEN 'current'
               WHEN start_time BETWEEN (SELECT p1_start FROM periods) AND (SELECT p1_end FROM periods) THEN 'prev_1'
               WHEN start_time BETWEEN (SELECT p2_start FROM periods) AND (SELECT p2_end FROM periods) THEN 'prev_2'
           END AS period_tag
    FROM irrigation_data
    WHERE farm_id = ?
      AND (? = 0 OR irrigation_sector_id = ?)
)
SELECT period_tag,
       SUM(real_amount)                             AS total_volume_mm,
       COUNT(id)                                    AS total_events,
       AVG(real_amount / NULLIF(nominal_amount, 0)) AS avg_efficiency,
       MIN(real_amount / NULLIF(nominal_amount, 0)) AS min_efficiency,
       MAX(real_amount / NULLIF(nominal_amount, 0)) AS max_efficiency
FROM filtered_data
WHERE period_tag IS NOT NULL
GROUP BY period_tag;
`

	err := repo.db.Raw(query,
		startDate, endDate, // c_start, c_end
		startDate, endDate, // p1_start, p1_end
		startDate, endDate, // p2_start, p2_end
		farmID,             // farm_id
		sectorID, sectorID, // irrigation_sector_id
	).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}

func (repo *FarmRepo) TimeSeriesByAggregation(farmID uint, sectorID uint, startDate, endDate time.Time, aggregation string) ([]*read_models.TimeSeriesAnalytics, error) {
	var results []*read_models.TimeSeriesAnalytics

	query := `
SELECT
    DATE_TRUNC(?, start_time) AS date_period,
    SUM(nominal_amount) AS nominal_amount_mm,
    SUM(real_amount) AS real_amount_mm,
    SUM(real_amount) / NULLIF(SUM(nominal_amount), 0) AS efficiency,
    COUNT(id) AS event_count
FROM irrigation_data
WHERE farm_id = ? AND (? = 0 OR irrigation_sector_id = ?)
  AND start_time BETWEEN ? AND ?
GROUP BY date_period
ORDER BY date_period;
`

	err := repo.db.Raw(query,
		aggregation,
		farmID,
		sectorID, sectorID,
		startDate,
		endDate,
	).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}

func (repo *FarmRepo) SectorBreakdownAnalytics(farmID uint, sectorID uint, startDate, endDate time.Time) ([]*read_models.BreakdownAnalytics, error) {
	var results []*read_models.BreakdownAnalytics

	query := `
SELECT
    s.id AS sector_id,
    s.name AS sector_name,
    SUM(d.real_amount) AS total_volume_mm,
    AVG(d.real_amount / NULLIF(d.nominal_amount, 0)) AS average_efficiency
FROM irrigation_sectors s
         JOIN irrigation_data d ON s.id = d.irrigation_sector_id
WHERE d.farm_id = ? AND (? = 0 OR s.id = ?)
  AND d.start_time BETWEEN ? AND ?
GROUP BY s.id, s.name;
`

	err := repo.db.Raw(query,
		farmID,             // farm_id
		sectorID, sectorID, // irrigation_sector_id
		startDate, // start_time
		endDate,   // end_time
	).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}

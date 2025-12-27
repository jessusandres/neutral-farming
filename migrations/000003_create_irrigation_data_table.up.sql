CREATE TABLE IF NOT EXISTS irrigation_data
(
    id                   BIGSERIAL PRIMARY KEY,
    farm_id              BIGINT      NOT NULL,
    irrigation_sector_id BIGINT      NOT NULL,
    start_time           TIMESTAMPTZ NOT NULL,
    end_time             TIMESTAMPTZ NOT NULL,
    nominal_amount       NUMERIC(10, 2)       DEFAULT 0.00,
    real_amount          NUMERIC(10, 2)       DEFAULT 0.00,
    created_at           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE irrigation_data
    ADD CONSTRAINT fk_farm FOREIGN KEY (farm_id) REFERENCES farms (id) ON DELETE RESTRICT;

ALTER TABLE irrigation_data
    ADD CONSTRAINT fk_sector FOREIGN KEY (irrigation_sector_id) REFERENCES irrigation_sectors (id) ON DELETE RESTRICT;

CREATE INDEX IF NOT EXISTS idx_irrigation_farm_time
    ON irrigation_data (farm_id, start_time);

CREATE INDEX IF NOT EXISTS idx_irrigation_sector_time
    ON irrigation_data (irrigation_sector_id, start_time);


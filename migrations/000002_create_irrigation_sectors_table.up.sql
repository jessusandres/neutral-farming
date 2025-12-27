CREATE TABLE IF NOT EXISTS irrigation_sectors
(
    id         BIGSERIAL PRIMARY KEY,
    farm_id    BIGINT       NOT NULL,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

ALTER table irrigation_sectors
    ADD CONSTRAINT fk_farm FOREIGN KEY (farm_id) REFERENCES farms (id) ON DELETE RESTRICT;

CREATE INDEX IF NOT EXISTS idx_irrigation_sectors_farm_id ON irrigation_sectors (farm_id);

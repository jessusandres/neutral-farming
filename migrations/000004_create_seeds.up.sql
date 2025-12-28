TRUNCATE irrigation_data, irrigation_sectors, farms RESTART IDENTITY CASCADE;

INSERT INTO farms (name)
VALUES ('Hacienda Perú'),
       ('Hacienda Chile');

INSERT INTO irrigation_sectors (farm_id, name)
VALUES (1, 'Sector Norte - Maíz'),
       (1, 'Sector Sur - Soja'),
       (1, 'Sector Este - Trigo');


INSERT INTO irrigation_sectors (farm_id, name)
VALUES (2, 'Lote A - Uva'),
       (2, 'Lote B - Cerezo'),
       (2, 'Lote C - Olivos');

INSERT INTO irrigation_data (farm_id,
                             irrigation_sector_id,
                             start_time,
                             end_time,
                             nominal_amount,
                             real_amount)
SELECT s.farm_id,
       s.id,
       ts_date + TIME '06:00:00'                                                         as start_time,
       ts_date + TIME '06:00:00' + (random() * (interval '2 hours') + interval '1 hour') as end_time,
       (random() * 15 + 10)::numeric(10, 2)                                              as nominal_amount,
       ((random() * 15 + 10) * (random() * 0.4 + 0.7))::numeric(10, 2)                   as real_amount
FROM
    generate_series('2021-01-01'::timestamp, '2025-12-31'::timestamp, '1 day') as ts_date
        CROSS JOIN irrigation_sectors s;

-- 5. Edge Cases

-- CASO A: Riego Nominal Cero
UPDATE irrigation_data
SET nominal_amount = 0.00,
    real_amount    = 5.00
WHERE id = (SELECT id
            FROM irrigation_data
            WHERE start_time BETWEEN '2024-01-15 00:00' AND '2024-01-15 23:59'
            LIMIT 1);

-- CASO B: Riego Real Cero (Eficiencia cero)
UPDATE irrigation_data
SET nominal_amount = 15.00,
    real_amount    = 0.00
WHERE id = (SELECT id
            FROM irrigation_data
            WHERE start_time BETWEEN '2024-01-16 00:00' AND '2024-01-16 23:59'
            LIMIT 1);

-- CASO C: Sobre-riego extremo (eficiencia > 200%)
UPDATE irrigation_data
SET nominal_amount = 10.00,
    real_amount    = 35.00
WHERE id = (SELECT id
            FROM irrigation_data
            WHERE start_time BETWEEN '2024-01-17 00:00' AND '2024-01-17 23:59'
            LIMIT 1);

-- CASO D: Datos faltantes en el pasado (Gap en la historia)
DELETE
FROM irrigation_data
WHERE start_time BETWEEN '2024-01-20 00:00' AND '2024-01-25 23:59';

-- 6. Conteo de eventos
SELECT COUNT(*)        as total_records,
       MIN(start_time) as first_record,
       MAX(start_time) as last_record
FROM irrigation_data;


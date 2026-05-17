CREATE TABLE ref.location (
  id BIGSERIAL PRIMARY KEY,
  parent_id BIGINT REFERENCES ref.location(id),
  type VARCHAR(30) NOT NULL,
  title VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);

-- WITH RECURSIVE location_tree AS (
--     SELECT id
--     FROM ref.location
--     WHERE id = $1
--     UNION ALL
--     SELECT l.id
--     FROM ref.location l
--     JOIN location_tree lt ON l.parent_id = lt.id
-- )
-- SELECT a.*
-- FROM app.announcement a
-- WHERE a.location_id IN (
--     SELECT id FROM location_tree
-- );
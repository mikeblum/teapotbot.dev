-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION created_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.created = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION updated_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
DROP FUNCTION IF EXISTS created_timestamp;
DROP FUNCTION IF EXISTS updated_timestamp;

-- +goose Up
CREATE TABLE IF NOT EXISTS public.crawls (
    id TEXT PRIMARY KEY DEFAULT generate_id( 'public', 'crawls', 'id', 10 ) CONSTRAINT idchk CHECK (char_length(id) >= 10),
    url TEXT NOT NULL,
    status_code SMALLINT NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ
);

CREATE OR REPLACE TRIGGER created_timestamp
BEFORE INSERT ON public.crawls
FOR EACH ROW
EXECUTE PROCEDURE created_timestamp();

CREATE OR REPLACE TRIGGER updated_timestamp
BEFORE UPDATE ON public.crawls
FOR EACH ROW
EXECUTE PROCEDURE updated_timestamp();

-- +goose Down
DROP TABLE IF EXISTS crawls CASCADE;

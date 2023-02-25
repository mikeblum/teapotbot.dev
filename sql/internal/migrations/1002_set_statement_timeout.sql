-- +goose Up

ALTER DATABASE teapotbot SET statement_timeout = '60s';
ALTER DATABASE teapotbot SET idle_in_transaction_session_timeout = '60s';

-- +goose Down

ALTER DATABASE teapotbot SET statement_timeout = 0;
ALTER DATABASE teapotbot SET idle_in_transaction_session_timeout = 0;

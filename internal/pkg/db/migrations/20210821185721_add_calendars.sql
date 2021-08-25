-- +goose Up
-- +goose StatementBegin
CREATE TABLE "calendars"
(
    Id     serial primary key,
    UserId INT     NOT NULL,
    Type   INT     NOT NULL,
    Link   VARCHAR NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "calendars"
-- +goose StatementEnd

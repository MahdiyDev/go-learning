-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table todos(
	id serial primary key not null,
	title varchar(255) not null,
	context varchar(255) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table todos
-- +goose StatementEnd

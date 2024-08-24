-- +goose Up
-- +goose StatementBegin
alter table inventory_items
    rename ammount to amount;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table inventory_items
    rename amount to ammount;
-- +goose StatementEnd

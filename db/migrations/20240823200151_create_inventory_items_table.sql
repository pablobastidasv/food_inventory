-- +goose Up
-- +goose StatementBegin
CREATE TABLE inventory_items (
    id uuid not null,
    product_id uuid not null,
    ammount smallserial,
    constraint pk_inventory PRIMARY KEY(id),
    constraint fk_inventory_product FOREIGN KEY(product_id) REFERENCES products(id)
)
;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE inventory_items
     drop constraint fk_inventory_product
;

DROP TABLE inventory_items
;
-- +goose StatementEnd

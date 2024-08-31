-- +goose Up
ALTER TABLE inventory_items
    drop constraint fk_inventory_product,
    add constraint fk_inventory_product 
        FOREIGN KEY (product_id) REFERENCES products(id) 
            on delete cascade
;

-- +goose Down
ALTER TABLE inventory_items
    drop constraint fk_inventory_product,
    add constraint fk_inventory_product 
        FOREIGN KEY (product_id) REFERENCES products(id) 
;

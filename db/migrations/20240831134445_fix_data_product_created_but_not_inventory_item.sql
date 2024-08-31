-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

insert into inventory_items (id, product_id, amount)
select uuid_generate_v4(), p.id, 0 
from products p
where p.id not in (select ii.product_id from inventory_items ii);

-- +goose Down
select 'hello world!!'

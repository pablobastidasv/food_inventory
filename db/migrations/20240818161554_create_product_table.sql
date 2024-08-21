-- +goose Up
create table categories (
  code varchar(32) not null,
  name text not null,
  parent varchar(32),
  constraint pk_categories PRIMARY KEY(code),
  constraint fk_categories_parent FOREIGN KEY(parent) REFERENCES categories(code)
);

create table products(
  id uuid not null,
  name text not null,
  category_code varchar(32) not null,
  constraint pk_products PRIMARY KEY(id),
  constraint fk_products_categories FOREIGN KEY(category_code) REFERENCES categories(code)
);

-- +goose Down
ALTER TABLE products
     drop constraint fk_products_categories;

DROP TABLE categories;
DROP TABLE products;


-- +goose Up
-- +goose StatementBegin
insert into categories(code, name, parent) values
    ('VEGETABLES', 'Vegetales', null),
    ('MEAT', 'Carne', null),
    ('FRIED_FOOD', 'Frituras', null),
    ('FRUITS', 'Frutas', null),
    ('PORK', 'Cerdo', 'MEAT'),
    ('BEEF', 'Res', 'MEAT'),
    ('CHICKEN', 'Pollo', 'MEAT')
;

INSERT INTO products(id, name, category_code) VALUES 
    ('e8f2d7fe-ce0b-4ebf-8380-3599106b997b','Cerdo Wok','PORK'),
    ('3dea6634-7bb5-4ee6-be0c-537eeeb6d795','Res Wok','BEEF'),
    ('28bee164-4600-4cdc-aa55-2fc467df4691','Pollo Wok','CHICKEN'),
    ('16ba7af5-3023-4657-b374-7182839b24f4','Costilla fritar','PORK'),
    ('33ba85b1-2374-48f2-85cd-8f3907d93993','Higadillas','CHICKEN'),
    ('cbb68ac9-f638-436d-99d3-b103aecc73ad','Cubitos Cerdo','PORK'),
    ('6ffb90bb-d879-4376-a537-4a1798de6b3e','Cubitos Pollo','CHICKEN'),
    ('d144aecf-25b1-4e8c-ae3c-ad3fa73a718d','Filete cerdo','PORK'),
    ('ceb8fe52-4970-4d3c-b91d-e46e9a98785d','Filete Res','BEEF'),
    ('62ac76b5-718e-49bb-b452-1ada5050d9ae','Filete Pollo','CHICKEN'),
    ('544cd14c-f2bf-461b-aeb0-6c99159ae07f','Horno Cerdo','PORK'),
    ('865896b8-1d23-488c-b45a-145e32a4905e','Horno Res','BEEF'),
    ('aa4a59aa-850b-4ef8-a2be-890febc622be','Sobrebarriga','BEEF'),
    ('5037c655-e8fd-4677-a201-6e0203f8a9da','Chuleta de cerdo','PORK'),
    ('c854028f-60a4-43a8-9c99-ed5dc2dbc063','Pescuezos','CHICKEN'),
    ('f63f6688-0877-4b26-b8d3-6742d9716800','Zanahoria Cubitos','VEGETABLES'),
    ('48aca900-ea0a-4c19-8b70-0ddfed1674f8','Mango','VEGETABLES'),
    ('7b32f722-865a-495e-864b-32e13fdb2069','Habichuela','VEGETABLES'),
    ('da5511ab-b9b5-488e-b477-5cc5e22f8b9a','Arberja','VEGETABLES'),
    ('f6f164b0-ac6b-4564-acbc-3cbd628e9c98','Sandia','VEGETABLES'),
    ('e63fd477-7371-4d1f-ab8a-2fa029e1fd1a','Arbolitos','VEGETABLES'),
    ('49796990-cb28-4657-8fdc-e9ce9dd57e57','Apio','VEGETABLES'),
    ('c28b20b5-acb9-4582-9bc2-e4203c77389f','Patacones','VEGETABLES'),
    ('9731acb8-f816-42fe-9ae2-4e06f0692ffc','Yuca','VEGETABLES'),
    ('07cf03e1-4d5e-4b42-86fb-ec342ed13025','Mazorca','VEGETABLES'),
    ('490f1843-651e-453f-a714-13c0c9bf312a','Platano Moneditas','VEGETABLES'),
    ('4c9b01a2-b278-43c8-b0b3-3598256e8fd9','Platano Sancocho','VEGETABLES'),
    ('dd392c45-1566-44d6-b995-d35e7391f258','Papa Cubitos','VEGETABLES'),
    ('976b244d-7c39-4f93-80ee-bd44f6909e12','Papa Rellena','FRIED_FOOD'),
    ('f55ee4ae-6f43-4565-93f3-66ec95b237db','Empanadas','FRIED_FOOD'),
    ('39c923ca-eaab-45e6-8f1c-f966550d2d40','Arepas','FRIED_FOOD'),
    ('2bb47d8d-dfa5-42bc-909f-d13bdf67b603','Marranitas','FRIED_FOOD'),
    ('282ade86-4a58-4142-891f-434e4b651ae9','Cerdo Tiritas','PORK'),
    ('fd0bad31-be81-42c0-93ee-5d1402283794','Costilla sopa','PORK'),
    ('fc4157ba-9b01-432f-9b1d-ac8bdabbfe28','Chicharron Cubitos','PORK'),
    ('0560970e-b5da-4353-ada0-71c26c46869e','Chicharron Cubos','PORK'),
    ('9fc2a413-adf6-4c8e-8579-9938f0787e2f','Bananos','FRUITS'),
    ('dcbf9699-cfc6-4885-93d6-cbe0037e619e','Zanahoria Circulos','VEGETABLES'),
    ('a4455706-442f-45b1-be81-93380bbed32f','Papas Dulces','FRIED_FOOD'),
    ('45817543-882c-498f-9f95-b2cc62de3526','Papas francesas','FRIED_FOOD'),
    ('1db91a65-f450-42c4-a980-f6371914c032','Papas cascos','FRIED_FOOD'),
    ('16aad1cf-b661-43e5-93f2-9a8c3bee9428','Zapallo','VEGETABLES'),
    ('07f23a1a-0ef3-445e-ac4a-770efbe1ea82','BlueBerries','FRUITS'),
    ('9fdf4f5a-ae24-4e26-85ae-a2c3e579da8b','Churros','FRIED_FOOD'),
    ('637f930b-2002-44ef-961f-fe502e7765c6','Fresas','FRUITS')
;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from products;
delete from categories;
-- +goose StatementEnd

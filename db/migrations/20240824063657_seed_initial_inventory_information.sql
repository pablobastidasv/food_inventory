-- +goose Up
-- +goose StatementBegin
insert into inventory_items (id, product_id, ammount)
 values ('bade892a-6423-4bd7-825f-86418c80be46','e8f2d7fe-ce0b-4ebf-8380-3599106b997b',0),
        ('1141020e-1c19-4a68-a5c7-f758062281bc','3dea6634-7bb5-4ee6-be0c-537eeeb6d795',1),
        ('4611dd0d-3a89-4c48-a368-36185213e75f','28bee164-4600-4cdc-aa55-2fc467df4691',1),
        ('b683080f-9b94-4f57-8cbf-e477db800fa6','16ba7af5-3023-4657-b374-7182839b24f4',1),
        ('cdd05988-3238-410f-a25f-8986f2356181','33ba85b1-2374-48f2-85cd-8f3907d93993',5),
        ('b59530ab-5664-433f-a6a1-1d8be696ec5e','cbb68ac9-f638-436d-99d3-b103aecc73ad',1),
        ('5a465833-3336-4271-b3a6-7614da48305a','6ffb90bb-d879-4376-a537-4a1798de6b3e',0),
        ('b7520257-df00-49af-a9e6-89c411d40285','d144aecf-25b1-4e8c-ae3c-ad3fa73a718d',2),
        ('30a5f5ab-e636-481e-a41d-ca749500196f','ceb8fe52-4970-4d3c-b91d-e46e9a98785d',1),
        ('011443da-af71-42ba-833d-bf2d6be7727c','62ac76b5-718e-49bb-b452-1ada5050d9ae',0),
        ('c8855bcb-5c27-49d4-a371-6178db316e48','544cd14c-f2bf-461b-aeb0-6c99159ae07f',1),
        ('59cc811e-5385-4bbc-bea0-6c7997029342','865896b8-1d23-488c-b45a-145e32a4905e',0),
        ('5c29a633-f4bf-4b5b-87c6-af6892447718','aa4a59aa-850b-4ef8-a2be-890febc622be',3),
        ('5b18a518-7714-442b-9d85-3d9d8a75b71e','5037c655-e8fd-4677-a201-6e0203f8a9da',2),
        ('d52a86df-6e5a-4048-821a-c167cfd3ce91','c854028f-60a4-43a8-9c99-ed5dc2dbc063',7),
        ('d3d8384b-4f1d-4bf3-96a2-f07e5266558b','f63f6688-0877-4b26-b8d3-6742d9716800',1),
        ('8cd20c25-f3f5-4a02-8a61-53025ab4b099','48aca900-ea0a-4c19-8b70-0ddfed1674f8',0),
        ('bcaa3e28-3da7-444a-84d8-6a95eac04937','7b32f722-865a-495e-864b-32e13fdb2069',0),
        ('4274e527-273e-4825-bd02-a5a96270e6fa','da5511ab-b9b5-488e-b477-5cc5e22f8b9a',0),
        ('67a747a0-24e8-413c-bf76-c582b6025544','f6f164b0-ac6b-4564-acbc-3cbd628e9c98',2),
        ('838d7ab7-759b-434e-b095-7538993983d4','e63fd477-7371-4d1f-ab8a-2fa029e1fd1a',2),
        ('a661e466-0fd7-45fe-b5fc-8ac054ed1102','49796990-cb28-4657-8fdc-e9ce9dd57e57',0),
        ('60a2e376-7719-40e1-aa1d-bd9f36889d65','c28b20b5-acb9-4582-9bc2-e4203c77389f',0),
        ('6cfd780f-0987-4dde-8130-cd46c0ef39df','9731acb8-f816-42fe-9ae2-4e06f0692ffc',1),
        ('411f6a5d-8e18-459e-aa7b-640f8e16ba14','07cf03e1-4d5e-4b42-86fb-ec342ed13025',2),
        ('d71e2221-33df-4ef5-93e0-2af18349497f','490f1843-651e-453f-a714-13c0c9bf312a',0),
        ('474c12ab-fed8-44c1-a62e-4d3d9b4a5354','4c9b01a2-b278-43c8-b0b3-3598256e8fd9',0),
        ('47ccddef-e48e-40a6-9eb1-5f7954be842a','dd392c45-1566-44d6-b995-d35e7391f258',0),
        ('c5469944-322a-425a-a10b-5266ef13103e','976b244d-7c39-4f93-80ee-bd44f6909e12',0),
        ('efbdb059-45e6-4143-9398-f924793ceac1','f55ee4ae-6f43-4565-93f3-66ec95b237db',0),
        ('2ed57f61-c704-40a0-872c-aa34c9bed5f5','39c923ca-eaab-45e6-8f1c-f966550d2d40',0),
        ('5dd2bfc6-ce28-40fb-be2b-458c9c16dcc1','2bb47d8d-dfa5-42bc-909f-d13bdf67b603',0),
        ('82da206c-383f-4815-ab67-44b5b5496409','282ade86-4a58-4142-891f-434e4b651ae9',2),
        ('5e3b4c12-9d1c-4b50-8fcf-e46e0a1aa26d','fd0bad31-be81-42c0-93ee-5d1402283794',0),
        ('a1d5a842-d4d7-4ee1-ba48-28931f66cc1d','fc4157ba-9b01-432f-9b1d-ac8bdabbfe28',1),
        ('c3a34d64-eeaf-45d7-906a-d94178962d4c','0560970e-b5da-4353-ada0-71c26c46869e',1),
        ('fd17c6e2-ec66-4f85-a46c-c01ba24efbba','9fc2a413-adf6-4c8e-8579-9938f0787e2f',4),
        ('7b4f23c1-7742-40da-93d0-45547f02035f','dcbf9699-cfc6-4885-93d6-cbe0037e619e',1),
        ('24dbd801-725f-4d67-8148-825b8ce65aed','a4455706-442f-45b1-be81-93380bbed32f',0),
        ('0599c0a7-bfa0-401b-b126-0f045414d4a7','45817543-882c-498f-9f95-b2cc62de3526',0),
        ('17d479e1-b239-45b2-b38a-459db165d92d','1db91a65-f450-42c4-a980-f6371914c032',0),
        ('4437ec66-0550-4087-be34-29699bc1b6ff','16aad1cf-b661-43e5-93f2-9a8c3bee9428',0),
        ('ccabbf1d-3b6f-4ab1-86ae-12b7124e41a0','07f23a1a-0ef3-445e-ac4a-770efbe1ea82',0),
        ('6f8238c8-cb9d-4597-a2bc-6fd40e0435ca','9fdf4f5a-ae24-4e26-85ae-a2c3e579da8b',0),
        ('ff90089d-e6a4-45d4-aed3-56f692546b90','637f930b-2002-44ef-961f-fe502e7765c6',4)
;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM inventory_items;
-- +goose StatementEnd

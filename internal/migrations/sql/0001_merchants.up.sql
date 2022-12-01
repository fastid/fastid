create table merchants
(
    merchant_id      uuid                                               not null primary key,
    name             varchar(200)                                       not null,
    created_at       timestamp with time zone default CURRENT_TIMESTAMP not null,
    updated_at       timestamp with time zone default CURRENT_TIMESTAMP not null,
    site_id          varchar(10)                                        not null,
    token            uuid                                               not null,
    notification_key uuid                                               not null
);

insert into merchants (merchant_id, name, created_at, updated_at, site_id, token, notification_key)
values
    (
     '77f512bc-7c2c-435c-9ada-764ae06d60c3',
     'iperon',
     now(),
     now(),
     'fcdv21-00',
     '17ce2805-b57f-45ea-afc4-9258675b3358',
     'a95979d8-e0ab-4da9-bbba-92eb57698c12'
);

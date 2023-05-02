CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE device_lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255)
);

CREATE TABLE users_lists
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    list_id int references device_lists (id) on delete cascade not null
);

CREATE TABLE device_items
(
    id          serial  not null unique,
    title       varchar(255) not null,
    description varchar(255),
    isPowerOn boolean not null default false
);

CREATE TABLE lists_device_items
(
    id serial not null unique ,
    item_id int references device_items (id) on DELETE cascade not null,
    list_id int references device_lists (id) on DELETE cascade not null
);
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
    name       varchar(255) not null,
    place varchar(255),
    condition boolean not null default false
);

CREATE TABLE lists_device_items
(
    id serial not null unique ,
    item_id int references device_items (id) on DELETE cascade not null,
    list_id int references device_lists (id) on DELETE cascade not null
);

CREATE TABLE device_lights
(
    id          serial  not null unique,
    name       varchar(255) not null,
    place varchar(255),
    condition boolean not null default false
);

CREATE TABLE users_lights
(
    id serial not null unique ,
    user_id int references users (id) on DELETE cascade not null,
    light_id int references device_lights (id) on DELETE cascade not null
);

CREATE TABLE device_cameras
(
    id          serial  not null unique,
    name       varchar(255) not null,
    place varchar(255)
);

CREATE TABLE users_cameras
(
    id serial not null unique ,
    camera_id int references device_cameras (id) on DELETE cascade not null,
    user_id int references users (id) on DELETE cascade not null
);

CREATE TABLE device_detectors
(
    id          serial  not null unique,
    name       varchar(255) not null,
    place varchar(255),
    statement boolean not null default false
);

CREATE TABLE users_detectors
(
    id serial not null unique ,
    detector_id int references device_detectors (id) on DELETE cascade not null,
    user_id int references users (id) on DELETE cascade not null
);
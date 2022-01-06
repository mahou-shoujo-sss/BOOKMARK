CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE list
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    link        varchar(255) not null 
);

CREATE TABLE users_lists
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    list_id int references list (id) on delete cascade not null
);
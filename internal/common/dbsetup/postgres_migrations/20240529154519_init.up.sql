create table if not exists users (
    uid text not null,
    username text not null unique,
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null
);

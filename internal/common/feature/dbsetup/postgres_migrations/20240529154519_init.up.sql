create table if not exists usernames (
  uid text not null primary key,
  username text not null unique,
  created_at timestamp default current_timestamp not null,
  updated_at timestamp default current_timestamp not null
);

create type room_statuses as enum ('open', 'closed');

create table if not exists rooms (
  id uuid not null default gen_random_uuid() primary key,
  code text not null unique,
  status room_statuses not null,
  created_at timestamp default current_timestamp not null,
  updated_at timestamp default current_timestamp not null
);

create table if not exists room_members (
  id uuid not null default gen_random_uuid() primary key,
  is_owner boolean not null,
  room_id uuid not null,
  user_uid text not null,
  created_at timestamp default current_timestamp not null,
  updated_at timestamp default current_timestamp not null, 

  foreign key (room_id)
    references rooms (id),
  unique (room_id, user_uid)
);

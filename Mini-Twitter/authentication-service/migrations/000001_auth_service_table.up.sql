create type role as enum('admin', 'user','c-admin');

create table if not exists users(
    id uuid default gen_random_uuid() primary key,
    phone varchar(13)unique ,
    email varchar unique,
    password varchar,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    deleted_at bigint default 0,
    unique(email, deleted_at)
);

create table if not exists user_profile(
    user_id uuid references users(id),
    first_name varchar,
    last_name varchar,
    username varchar unique,
    nationality varchar,
    bio varchar,
    role role default 'user',
    profile_image varchar default 'no images',
    followers_count int default 0,
    following_count int default 0,
    posts_count int default 0,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    is_active bool default 'true'
);

create table if not exists follows(
    follower_id uuid references users(id),
    following_id uuid references users(id),
    created_at timestamp with time zone default now()
)
create table tweets (
    id uuid default gen_random_uuid() primary key,
    user_id uuid not null,
    hashtag varchar default 'dodi',
    title varchar,
    content text,
    image_url varchar default 'no images',
    tweet_id uuid references tweets(id),
    is_retweeted bool default 'false',
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    deleted_at bigint
);

create table comments (
    id uuid default gen_random_uuid() unique,
    user_id uuid not null,
    tweet_id uuid references tweets(id),
    context text,
    like_count int default 0,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

create table likes (
    user_id uuid not null,
    tweet_id uuid not null,
    created_at timestamp with time zone default now(),
    UNIQUE (user_id, tweet_id)
)
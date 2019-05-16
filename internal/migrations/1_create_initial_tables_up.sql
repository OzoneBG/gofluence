create table users (
    id serial primary key,
    username varchar not null,
    password varchar not null,
    email varchar not null
);

create table articles (
    id serial primary key,
    title varchar not null,
    content varchar not null,
    author_id integer references users (id)
);
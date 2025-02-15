create table common_messages
(
    id               bigint not null
        primary key,
    message_type     text   not null,
    room_id          bigint not null,
    room_display_id  text   not null,
    room_name        text   not null,
    user_name        text   not null,
    user_id          bigint not null,
    user_display_id  text   not null,
    content          text   not null,
    timestamp        bigint not null,
    favorite_user_id bigint default 0
);

alter table common_messages
    owner to postgres;

create index common_messages_room_display_id_index
    on common_messages (room_display_id);

create index common_messages_timestamp_index
    on common_messages (timestamp);

create index common_messages_user_id_index
    on common_messages (user_id);

create index common_messages_user_name_index
    on common_messages (user_name);

create table gift_messages
(
    id                 bigint not null
        primary key,
    user_id            bigint not null,
    user_name          text   not null,
    user_display_id    text   not null,
    to_user_id         bigint not null,
    to_user_name       text   not null,
    to_user_display_id text   not null,
    gift_name          text   not null,
    gift_id            bigint not null,
    room_id            bigint not null,
    room_display_id    text   not null,
    room_name          text   not null,
    message            text   not null,
    timestamp          bigint not null,
    diamond_count      bigint not null,
    image_url          text,
    repeat_end         integer,
    combo_count        text
);

alter table gift_messages
    owner to postgres;

create index gift_messages_room_display_id_index
    on gift_messages (room_display_id);

create index gift_messages_timestamp_index
    on gift_messages (timestamp);

create index gift_messages_to_user_id_index
    on gift_messages (to_user_id);

create index gift_messages_user_id_index
    on gift_messages (user_id);

create table users
(
    id         bigserial
        primary key,
    user_id    bigint not null,
    display_id text   not null,
    user_name  text   not null
);

alter table users
    owner to postgres;

create index idx_users_user_id
    on users (user_id);

create table live_confs
(
    id              bigserial
        primary key,
    room_display_id text not null
        constraint unique_display_id
            unique,
    url             text not null
        constraint unique_url
            unique,
    name            text
        constraint unique_name
            unique,
    modified_on     bigint,
    created_on      bigint,
    modified_by     text,
    crated_by       text,
    cron            text,
    enable          boolean default true
);

alter table live_confs
    owner to postgres;

create table if not exists users(
    id bigserial primary key,
    name varchar(255) not null,
    email varchar(255) not null,
    password text not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

create table if not exists products(
    id bigserial primary key,
    name varchar(255) not null,
    stock smallint not null default 0,
    price double precision not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

create table if not exists orders (
    id bigserial primary key,
    user_id bigint not null,
    product_id bigint not null,
    quantity smallint not null default 0,
    total double precision not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    CONSTRAINT fk_product
        FOREIGN KEY (product_id)
            REFERENCES products(id)
                ON DELETE CASCADE,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES users(id)
                ON DELETE CASCADE
);
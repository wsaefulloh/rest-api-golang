CREATE TABLE public.users (
    id bigserial NOT NULL,
    name varchar(50) NOT NULL,
    username varchar(20) NOT NULL,
    email varchar(20) NOT NULL,
    role varchar NOT NULL,
    password varchar NOT NULL,
    update_at timestamp NOT NULL,
    created_at timestamp NOT NULL,
    CONSTRAINT users_pk PRIMARY KEY (id)
)
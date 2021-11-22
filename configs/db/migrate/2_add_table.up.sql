CREATE TABLE public.categories (
    id bigserial NOT NULL,
    name varchar(50) NOT NULL,
    update_at timestamp NOT NULL,
    created_at timestamp NOT NULL,
    CONSTRAINT categories_pk PRIMARY KEY (id)
);
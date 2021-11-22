CREATE TABLE public.products (
    id bigserial NOT NULL,
    name varchar(50) NOT NULL,
    price int NOT NULL,
    category int NOT NULL,
    update_at timestamp NOT NULL,
    created_at timestamp NOT NULL,
    CONSTRAINT products_pk PRIMARY KEY (id),
    CONSTRAINT products_fk FOREIGN KEY (category)
    REFERENCES public.categories (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
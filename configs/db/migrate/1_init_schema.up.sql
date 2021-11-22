CREATE TABLE public.users (
    id bigserial NOT NULL,
    name varchar(50) NOT NULL,
    username varchar(20) NOT NULL,
    email varchar(20) NOT NULL,
    password varchar NOT NULL,
    update_at timestamp NOT NULL,
    created_at timestamp NOT NULL,
    CONSTRAINT users_pk PRIMARY KEY (id)
)

-- CREATE TABLE public.categories (
--     id bigserial NOT NULL,
--     name varchar(50) NOT NULL,
--     update_at timestamp NOT NULL,
--     created_at timestamp NOT NULL,
--     CONSTRAINT categories_pk PRIMARY KEY (id)
-- );

-- CREATE TABLE public.categories (
--     "id" bigserial NOT NULL,
--     "name" varchar(50) NOT NULL,
--     "update_at" timestamp NOT NULL,
--     "created_at" timestamp NOT NULL,
--     CONSTRAINT categories_pk PRIMARY KEY ("id")
-- );

-- CREATE TABLE public.products (
--     "id" bigserial NOT NULL,
--     "name" varchar(50) NOT NULL,
--     "price" int NOT NULL,
--     "category" int NOT NULL,
--     "update_at" timestamp NOT NULL,
--     "created_at" timestamp NOT NULL,
--     CONSTRAINT products_pk PRIMARY KEY ("id"),
--     CONSTRAINT products_fk FOREIGN KEY ("category")
--     REFERENCES public.categories ("id")
--     ON DELETE CASCADE
--     ON UPDATE CASCADE
-- );

-- CREATE TABLE public.histories (
--     "id" bigserial NOT NULL,
--     "invoice" varchar NOT NULL,
--     "cashier" varchar NOT NULL,
--     "date" varchar NOT NULL,
--     "order" int NOT NULL,
--     "count" int NOT NULL,
--     CONSTRAINT histories_pk PRIMARY KEY ("id"),
--     CONSTRAINT histories_fk FOREIGN KEY ("order")
--     REFERENCES public.products ("id")
--     ON DELETE CASCADE
--     ON UPDATE CASCADE
-- );
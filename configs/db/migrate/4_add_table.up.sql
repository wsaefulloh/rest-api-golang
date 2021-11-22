CREATE TABLE public.histories (
    id bigserial NOT NULL,
    invoice varchar NOT NULL,
    cashier varchar NOT NULL,
    date varchar NOT NULL,
    "order" int NOT NULL,
    "count" int NOT NULL,
    CONSTRAINT histories_pk PRIMARY KEY (id),
    CONSTRAINT histories_fk FOREIGN KEY ("order")
    REFERENCES public.products (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
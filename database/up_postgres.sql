--CREATE DATABASE IF NOT EXISTS e-commerce;
\c e_commerce;

DROP TABLE IF EXISTS e_commerce;

CREATE TABLE
  public.products (
    id serial NOT NULL,
    name character varying (255) NOT NULL,
    description character varying (255) NOT NULL,
    price numeric (10,2) NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
  );

CREATE TABLE public.users (
    id serial NOT NULL,
    email character varying(128) NOT NULL UNIQUE,
    first_name character varying(64) NOT NULL,
    last_name character varying(64) NOT NULL,
    password character varying(128) NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
);


--FK PENDING: tokens_relation_1 trevor 60 CASCADE CASCADE
CREATE TABLE public.tokens (
    id integer NOT NULL,
    user_id integer,
    email character varying(255) NOT NULL,
    token character varying(255) NOT NULL,
    token_hash bytea NOT NULL,
    expiry timestamp with time zone NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

INSERT INTO public.products (name, description, price, created_at, updated_at)
VALUES
  ('Zapatillas deportivas', 'Zapatillas cómodas para correr', 79.99, now(), now()),
  ('Bolso de cuero', 'Bolso elegante para ocasiones especiales', 129.50, now(), now()),
  ('Camiseta estampada', 'Camiseta de algodón con diseño floral', 24.95, now(), now()),
  ('Reloj de pulsera', 'Reloj analógico con correa de acero inoxidable', 149.00, now(), now()),
  ('Gafas de sol', 'Gafas de sol polarizadas con montura de acetato', 59.99, now(), now()),
  ('Pantalones vaqueros', 'Vaqueros ajustados de color azul', 69.50, now(), now()),
  ('Vestido de fiesta', 'Vestido largo con lentejuelas y escote en V', 199.99, now(), now()),
  ('Botines de cuero', 'Botines negros con tacón bajo', 89.95, now(), now()),
  ('Pendientes de plata', 'Pendientes colgantes con piedras preciosas', 45.00, now(), now()),
  ('Bufanda de lana', 'Bufanda suave y abrigada para el invierno', 34.75, now(), now());


INSERT INTO public.users (email, first_name, last_name, password, created_at, updated_at) 
VALUES ('diego@diego.com', 'diego', 'last_name_placeholder', '$2a$12$ZTTPWrzmiT5wSz9gl2FrJuiP4wwoXNKriRbSnZuWCvN/ZgxutLYjG', now(), now());

SET statement_timeout = 0;

--bun:split

CREATE TABLE "public"."board_type"
(
    id          SERIAL PRIMARY KEY,
    create_by   int         not null,
    cafe_id     int         not null,
    name        VARCHAR(50) NOT NULL,
    description VARCHAR(2000),
    created_at  timestamptz
);


create unique index bt_cafe_id_name_unique on board_type (cafe_id, name);

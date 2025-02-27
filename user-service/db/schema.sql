CREATE TABLE users (
    id          uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name        VARCHAR NOT NULL,
    email       VARCHAR NOT NULL,
    password    TEXT NOT NULL
);
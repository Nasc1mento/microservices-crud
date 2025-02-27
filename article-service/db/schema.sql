CREATE TABLE articles (
    id          uuid PRIMARY KEY,
    author_id   uuid NOT NULL,
    title       VARCHAR NOT NULL,
    content     TEXT NOT NULL
);
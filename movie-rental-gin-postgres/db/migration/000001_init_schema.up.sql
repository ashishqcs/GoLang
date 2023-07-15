CREATE TABLE "movies" (
	"id" VARCHAR NOT NULL PRIMARY KEY,
	"title" VARCHAR NOT NULL,
	"released" DATE NOT NULL,
	"genre" VARCHAR NOT NULL,
	"actors" VARCHAR NOT NULL,
    "year" INT NOT NULL,
	"price" BIGINT NOT NULL,
    "quantity" INT NOT NULL
);

CREATE INDEX idx_genre ON movies (genre);

CREATE INDEX idx_actors ON movies (actors);

CREATE INDEX idx_year ON movies (year);
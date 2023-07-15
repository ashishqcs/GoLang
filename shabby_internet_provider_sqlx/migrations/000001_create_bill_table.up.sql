CREATE TABLE IF NOT EXISTS bill (
	id       	bigint primary key,
	name     	varchar,
	address  	varchar,
	plan_name	varchar,
	date    	varchar,
	amount   	float8
);

--postgres://user:password@localhost:5432/internet_bills?sslmode=disable&search_path=public

--migrate -path migrations -database "postgresql://user:password@localhost:5432/internet_bills?sslmode=disable" up
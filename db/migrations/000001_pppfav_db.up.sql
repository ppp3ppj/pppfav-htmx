BEGIN;

-- Set timezone
SET TIME ZONE 'asia/bangkok';
-- Install uuid extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Person id -> P0000001
CREATE SEQUENCE person_id_seq START WITH 1 INCREMENT BY 1;
-- GirlGroup id -> GG0000001
CREATE SEQUENCE girlgroup_id_seq START WITH 1 INCREMENT BY 1;
-- Member id -> M0000001
CREATE SEQUENCE member_id_seq START WITH 1 INCREMENT BY 1;
-- Song id -> S0000001
CREATE SEQUENCE song_id_seq START WITH 1 INCREMENT BY 1;

CREATE TABLE "Person" (
  "id" varchar(8) PRIMARY KEY DEFAULT CONCAT('P', LPAD(nextval('person_id_seq')::text, 7, '0')),
  "name" varchar NOT NULL,
  "age" int NOT NULL,
  "birth_date" date NOT NULL,
  "image_url" varchar,
  "description" varchar
);

CREATE TABLE "GirlGroup" (
  "id" varchar(9) PRIMARY KEY DEFAULT CONCAT('GG', LPAD(nextval('girlgroup_id_seq')::text, 8, '0')),
  "name" varchar NOT NULL,
  "debut_date" date NOT NULL,
  "company" varchar NOT NULL,
  "fandom_name" varchar NOT NULL,
  "description" varchar
);

CREATE TABLE "Member" (
  "id" varchar(8) PRIMARY KEY DEFAULT CONCAT('M', LPAD(nextval('member_id_seq')::text, 7, '0')),
  "person_id" varchar,
  "girl_group_id" varchar,
  "position" varchar
);

CREATE TABLE "Song" (
  "id" varchar(8) PRIMARY KEY DEFAULT CONCAT('S', LPAD(nextval('song_id_seq')::text, 7, '0')),
  "title" varchar NOT NULL,
  "release_date" date NOT NULL,
  "album" varchar,
  "girl_group_id" varchar,
  "description" varchar,
  "image_url" varchar
);

ALTER TABLE "Member" ADD FOREIGN KEY ("person_id") REFERENCES "Person" ("id");

ALTER TABLE "Member" ADD FOREIGN KEY ("girl_group_id") REFERENCES "GirlGroup" ("id");

ALTER TABLE "Song" ADD FOREIGN KEY ("girl_group_id") REFERENCES "GirlGroup" ("id");

COMMIT;

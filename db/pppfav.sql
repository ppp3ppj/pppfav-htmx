CREATE TABLE "Person" (
  "id" uuid PRIMARY KEY DEFAULT 'uuid_generate_v4()',
  "name" string,
  "age" int,
  "birth_date" date,
  "image_url" string,
  "description" text
);

CREATE TABLE "GirlGroup" (
  "id" uuid PRIMARY KEY DEFAULT 'uuid_generate_v4()',
  "name" string,
  "debut_date" date,
  "company" string,
  "fandom_name" string,
  "description" text
);

CREATE TABLE "Member" (
  "id" uuid PRIMARY KEY DEFAULT 'uuid_generate_v4()',
  "person_id" uuid,
  "girl_group_id" uuid,
  "position" string
);

CREATE TABLE "Song" (
  "id" uuid PRIMARY KEY DEFAULT 'uuid_generate_v4()',
  "title" string,
  "release_date" date,
  "album" string,
  "girl_group_id" uuid,
  "description" text,
  "image_url" string
);

COMMENT ON COLUMN "Person"."image_url" IS 'URL to the image of the person';

COMMENT ON COLUMN "Person"."description" IS 'Description or biography of the person';

COMMENT ON COLUMN "GirlGroup"."description" IS 'Description or biography of the girl group';

COMMENT ON COLUMN "Song"."description" IS 'Description or biography of the song';

COMMENT ON COLUMN "Song"."image_url" IS 'URL to the image of the song or album cover';

ALTER TABLE "Member" ADD FOREIGN KEY ("person_id") REFERENCES "Person" ("id");

ALTER TABLE "Member" ADD FOREIGN KEY ("girl_group_id") REFERENCES "GirlGroup" ("id");

ALTER TABLE "Song" ADD FOREIGN KEY ("girl_group_id") REFERENCES "GirlGroup" ("id");

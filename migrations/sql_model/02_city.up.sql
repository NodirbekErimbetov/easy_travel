CREATE TABLE "city" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "title" VARCHAR NOT NULL,
    "country_id" UUID REFERENCES "country"("id"),
    "city_code" VARCHAR NOT NULL,
    "latitude" VARCHAR NOT NULL,
    "longitude" VARCHAR NOT NULL,
    "offset" VARCHAR NOT NULL,
    "country_name" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
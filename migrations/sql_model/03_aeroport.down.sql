DROP TABLE "aeroport" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "title" VARCHAR NOT NULL,
    "country_id" UUID REFERENCES "country"("id"),
    "city_id" UUID REFERENCES "city"("id"),
    "latitude" NUMERIC,
    "longitude" NUMERIC,
    "radius" VARCHAR DEFAULT null,
    "image" VARCHAR DEFAULT '',
    "address" VARCHAR DEFAULT '',
    "country" VARCHAR NOT NULL,
    "city" VARCHAR NOT NULL,
    "search_text" VARCHAR,
    "code" VARCHAR  ,
    "product_count" NUMERIC ,
    "gmt" VARCHAR ,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
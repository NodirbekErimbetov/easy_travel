
jq -c '.[]' country.json >> your_data.json
cat country.json | jq -cr '.[]' | sed 's/\\[tn]//g' > countryjson.json
cat aeroport.json | jq -cr '.[]' | sed 's/\\[tn]//g' > aeroportjson.json


\copy temp(data) from './mock/countryjson.json';
\copy temp(data) from './mock_data/your_data.json';
\copy aeroportjson(data) from 'aeroportjson.json';

DELETE FROM temp WHERE length(data ->> 'country_id') = 2;

INSERT INTO country (
    "title",
    "code",
    "continent"
)
SELECT 
    JSON_BUILD_OBJECT(
        countrydata ->>'title'
        countrydata ->>'code'
        countrydata ->>'continent'
    )
FROM 
    countryjson;


INSERT INTO city(
    "title",
    "country_id",
    "city_code",
    "latitude",
    "longitude",
    "offset",
    "country_name",
    "updated_at"
) 
SELECT 
    JSON_BUILD_OBJECT(
        data -> 'title',
        data -> 'country_id',
        data -> 'city_code',
        data -> 'latitude',
        data -> 'longitude',
        data -> 'offset',
        data -> 'country_name',
        NOW()
    )
FROM 
    temp;

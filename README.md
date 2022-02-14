## Go Lang Test

Test 01:

- List races that are visible:

curl -X "POST" "http://localhost:8000/v1/list-races"      -H 'Content-Type: application/json'      -d $'{ "filter": {"visible": true} }'

- List races that are NOT visible:

curl -X "POST" "http://localhost:8000/v1/list-races"      -H 'Content-Type: application/json'      -d $'{ "filter": {"visible": false} }'

- List all races regardless of their visibility:

curl -X "POST" "http://localhost:8000/v1/list-races"      -H 'Content-Type: application/json'      -d $'{}'


- Combining filters:
curl -X "POST" "http://localhost:8000/v1/list-races"      -H 'Content-Type: application/json'      -d $'{ "filter": {"meeting_ids": [5], "visible": false} }'


Test 02:

- List all races sorted by advertised_start_time
curl -X "POST" "http://localhost:8000/v1/list-races"      -H 'Content-Type: application/json'      -d $'{ "filter": {}, "sortBy": { "property_name" : "advertised_start_time" } }'

- You can sort by any property, e.g.:

curl -X "POST" "http://localhost:8000/v1/list-races"      -H 'Content-Type: application/json'      -d $'{ "filter": {}, "sortBy": { "property_name" : "meeting_id" } }'

curl -X "POST" "http://localhost:8000/v1/list-races"      -H 'Content-Type: application/json'      -d $'{ "filter": {}, "sortBy": { "property_name" : "name" } }'

Test 03:

- See status property:

curl -X "POST" "http://localhost:8000/v1/list-races"      -H 'Content-Type: application/json'      -d $'{}'

Test 04:

- Fetch a single race by its ID:

curl -X "GET" "http://localhost:8000/v1/races/race/1"

curl -X "GET" "http://localhost:8000/v1/races/race/56"

Test 05:

- Sports service

- List all events:

curl -X "POST" "http://localhost:8000/v1/list-events"      -H 'Content-Type: application/json'      -d $'{}'

- Filter events by sport type:

curl -X "POST" "http://localhost:8000/v1/list-events"      -H 'Content-Type: application/json'      -d $'{ "filter": {"type": ["RUGBY","AFL"]}} }'

- Fetch single event:

curl -X "GET" "http://localhost:8000/v1/events/event/59"

curl -X "GET" "http://localhost:8000/v1/events/event/36"






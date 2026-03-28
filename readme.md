# NWS Demo Project
This project takes an HTTP request with coordinates and provides a forecast at that location.

More details about my thought process and scope in [notes](notes.md)

## Building
Requires:
1. go toolchain 1.25.3 or newer
2. gnu make
3. curl (optional)

### Build and Run
```bash
# Build
make build

# Tests
make test

# Run
make run
10:06PM WRN ctrl+c to exit
10:06PM INF API starting port=8080
```

### Included end-to-end tests (requires curl)
Assuming you've already [built and running](#build-and-run) the binary in another console

```bash
# Make a valid call
make call
curl -d '{"latitude": 36.7158451, "longitude": -91.8739187}' http://localhost:8080/api/v1/forecast
{"forecast":"Mostly Clear","short":"cold","temperatureFormat":"F","temperature":31}

# Make a call that the server knows is bad
make call-invalid
curl -d '{"latitude": 1000.0, "longitude": -91.8739187}' http://localhost:8080/api/v1/forecast
{"error":"invalid latitude"}

# Make a call that the NWS will return an error on
make call-nws-invalid
curl -d '{"latitude": -36.7158451, "longitude": -91.8739187}' http://localhost:8080/api/v1/forecast
{"error":"nws error\nData Unavailable For Requested Point"}
```

## Endpoints
### `POST /api/v1/forecast`
Body:
```json
{
    "latitude": 36.7158451,  // Float of WGS84 latitude (X)
    "longitude": -91.8739187 // Float of WGS84 longitude (Y)
}
```

#### Expected Response on OK
Code: 200
Body:
```json
{
    "forecast":"Mostly Clear",
    "short":"cold",
    "temperatureFormat":"F",
    "temperature":31
}
```

#### Expected Response on Error
Code: 400, 404, 500, etc
```json
{
    "error": "description of error"
}
```
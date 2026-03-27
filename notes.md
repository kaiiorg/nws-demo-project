# Notes
The following notes are what I'd normally keep in a scratch text document or write on paper. They are included here to get an idea of my thinking process for this project.

This implementation is more complex than is strictly needed, but is closer to what I'd implement for a real application instead of the bare minimum.

## Requirements as Provided
Write an HTTP server that serves the forecasted weather. Your server should expose an endpoint that:
1. Accepts latitude and longitude coordinates
2. Returns the short forecast for that area for Today (“Partly Cloudy” etc)
3. Returns a characterization of whether the temperature is “hot”, “cold”, or “moderate” (use your discretion on mapping temperatures to each type) 
4. Use the [National Weather Service API Web Service](https://www.weather.gov/documentation/services-web-api) as a data source.

The purpose of this exercise is to provide a sample of your work that we can discuss together in the Technical Interview.
* We respect your time. Spend as long as you need, but we intend it to take around an hour.
* We do not expect a production-ready service, but you might want to comment on your shortcuts.
* The submitted project should build and have brief instructions so we can verify that it works.
* The Coding Project should be written in the language for the job you’re applying for. 

## Out of Scope
The following things are considered out of scope for this project, mostly due to time constraints. In a real, production ready project these would all be in scope.

1. Security on endpoints
2. Graceful shutdown
3. Extensive automated end-to-end testing
4. Extensive automated unit testing
5. Extensive documentation and automated documentation generation

## Endpoint
Path: `/api/v1/forecast`
Body:
```json
{
    "latitude": 36.7158451,  // WGS84 latitude (X)
    "longitude": -91.8739187 // WGS84 longitude (Y)
}
```

### Endpoint Algorithm
1. Parse and sanity check values for body
    1. latitude:  -90 < X > 90
    2. longitude: -180.0 < Y > 180.0
2. API call to NWS /points/{latitude},{longitude}
    1. If not 200, return error to caller
3. Parse API call results for "gridId", "gridX", "gridY", "forecast" values
4. API call to NWS using "forcase" result, which includes "gridId", "gridX", "gridY"
    1. If not 200, return error to caller
5. Parse API call results for "temperature" and "temperatureUnit"
6. Map value to configured "characterization" ranges from config file

## Config
JSON for simplicity, but I'm quite fond of HCL.

```json
{
    "api": {
        "port": 8080 // Default to 8080
    },
    "forecast": {
        // Assumes Fahrenheit as this is the default scale used in the US.
        "hot": {
            "min": 90,  // Inclusive, leave undefined define this as the bottom range value
        }, 
        "moderate": {
            "max": 90,  // Exclusive, leave undefined define this as the top range value
            "min": 60,  // Inclusive, leave undefined define this as the bottom range value
        },
        "cold": {
            "max": 60,  // Exclusive, leave undefined define this as the top range value
        }
    }
}
```

## Places for hypothetical future improvement
1. Everything in [Out of Scope](./requirements.md#out-of-scope)
2. Allow user to make request in different coordinate systems, such as a given state plane
3. Allow administrator to define forecast mappings in Celsius
4. Cache results for a configurable amount of time to limit load on NWS resources and speed up response times

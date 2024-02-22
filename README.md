# go_weatherapi

## Intro:

In this project we are going to write a simple api that will call the weather api:

```go
https://api.openweathermap.org/data/2.5/weather?lat=63.5888&lon=154.4931&appid=apiId
```

and return the following data:

*This endpoint should return what the weather condition is outside in that area (snow, rain,
etc), whether it’s hot, cold, or moderate outside (use your own discretion on what temperature equates to
each type).*

This api will take 3 inputs:

```go
lat
lon
appid
```

Where the `appid` is the SECRET_KEY variable in the .env file

once you place *your* unique apikey, the project should run as normal 

Then to get the information I created a GET http route:

```go
/api/getweather?lat={lat}&lon={lon}
```

and that should return you the information of the weather conditions & temperature

### Other important variables that you can setup in your .env are:

`PORT` 

`BINARY`

Although the binary can be also added in the Makefile, you can add your BINARY executable name in the .env too.

### To build and execute the project run the following comands:
```
make build
make run
```
alternatively you could also just run:
```
make all
```
# go_weatherapi

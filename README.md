# go_weatherapi

## Intro:

In this project we are going to write a simple api that will call the weather api:

```
https://api.openweathermap.org/data/2.5/weather?lat=63.5888&lon=154.4931&appid=apiId
```

and return the following data:

*This endpoint should return what the weather condition is outside in that area (snow, rain,
etc), whether it’s hot, cold, or moderate outside (use your own discretion on what temperature equates to
each type).*

### This is an example json response
```
{
    "feels_like": "It's Cold",
    "conditon": "clear sky",
    "temp": 55.02,
    "place_name": "US",
    "country": "Miami"
}
```

This api will take 3 inputs:

```
lat
lon
appid
```

Where the `appid` is the **`SECRET_KEY`** variable in the .env file

### Other important variables that you can setup in your .env are:

`PORT` 

.env file example
```
SECRET_KEY=yoursecretkeyprovidedbyopenweathermap
PORT=8080
```

Once you place *your* unique apikey, you can build and run the project

### To build and execute the project run the following comands:
```
make build
make run
```
alternatively you could also just run:
```
make all
```

### Then to get the information I created a GET http route:

```
http://localhost:{PORT}/api/get-weather?lat={lat}&lon={lon}
```
### You can test it with the command
```
make get_miami_info
```
This will show the information in the example response in the console 

and that should return you the information of the weather conditions & temperature


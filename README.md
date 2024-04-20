# Web application (known as GOTTH and REST+ReactJs)

This project shows how to use templ+htmx on  frontside and rest+react on adminside.
Backend part is written on go+gorm it provides session auth for frontside and basic auth for adminside.
For sessions storage filesystem used. For db sqlite used.


## Live url

https://godash.fly.dev - frontside/userside

https://godash.fly.dev/dashboard - adminside admin:admin

## How to run locally

- You need to have docker.

### Build image

`docker build --tag=goweb .`

### Run image


`docker run -v ./var:/goweb_var -e DATABASE_URL='/goweb_var/gorm.sqlite' -e SESSIONS_DIR='/goweb_var/sessions' -p 8080:8080 goweb`

### Web app

- http://localhost:8080 - frontside
- http://localhost:8080/dashboard - adminside admin:admin

## Development

All commands better run in separated console

### Run tailwind css generator watch

`cd webapp`
`npm run tailwindcss_watch`

### Run webapp watch

`cd webapp`
`npm run watch`


### Run go app recompiler

`~/go/bin/air`

Usually air path is ~/go/bin/air 
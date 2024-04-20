# Web application (known as GOTTH and REST+ReactJs)

This project shows how to use templ+htmx on  frontside and rest+react on adminside.
Backend part is written on go+gorm it provides session auth for frontside and basic auth for adminside.
For sessions storage filesystem used. For db sqlite used.


## How to run locally

- You need to have docker.

### Build image

`docker build --tag=goweb .`

### Run image


`docker run -v ./var:/goweb_var -e DATABASE_URL='/goweb_var/gorm.sqlite' -e SESSIONS_DIR='/goweb_var/sessions' -p 8080:8080 goweb`

### Web app

- http://localhost:8080 - frontside
- http://localhost:8080/dashboard - adminside admin:admin


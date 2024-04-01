# ☕ coffee.mykal.codes

A simple microblog for my coffees and thoughts on coffee.

## Functional goals 

- A webapp that shows some of my recent coffees and notes on them.
- A simple interface for uploading new photos and notes on my coffees. 
- A simple way for people to like / react to my posts.
- A log of my coffees over time, maybe some fun graphs. 

## Technical goals

- Astro-based frontend (`/services/app`)
  - Should have a performant, accessible, and homepage.
  - Should provide an admin interface for writing and editing posts.
  - Should have animations using framer motion (want to learn).
- Go-based api (`/services/api`) 
  - Should use the standard go 1.22 `net/http` package.
  - Should provide cookie-based authentication for me to login with.
  - Should use a simple `sqlite` database with `gorm` store information. 


### Note: on splitting the app and api apart

I realize that astro is more than capable to handle the API portion. I want to learn go so I'm taking the chance to split the frontend and the API out.  
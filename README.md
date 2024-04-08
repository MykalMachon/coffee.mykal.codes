# ☕ coffee.mykal.codes

A simple microblog for my coffees and thoughts on coffee.

## Functional goals 

- A webapp that shows some of my recent coffees and notes on them.
- A simple interface for uploading new photos and notes on my coffees. 
- A simple way for people to like / react to my posts.
- A log of my coffees over time, maybe some fun graphs. 

## Technical goals

- Vite-based frontend (`/services/app`)
  - Should have a performant, accessible, homepage with react.
  - Should provide an admin interface for writing and editing posts.
  - Should have animations using framer motion (want to learn).
- Go-based api (`/services/api`) 
  - Should use the standard go 1.22 `net/http` package.
  - Should provide cookie-based/bearer authentication for me to login with.
  - Should use `gorm` and postgresql for the database.
  - Should allow users to upload photos through to cloudflare images.
- Web server reverse proxy to stitch things together (`/services/server`)
  - Should act as a reverse proxy for both the frontend and api 
  - Should provide a simple way to manage scaling up services if needed.

### Note: on splitting the app and api apart

I realize that astro is more than capable to handle the API portion. I want to learn go so I'm taking the chance to split the frontend and the API out.  

## Note: using caddy instead of NGINX for the reverse proxy

I used Caddy for the reverse proxy because I was having some issues with NGINX holding on to stale IPs when upstreams would redeploy. Here's my understanding of the issue after some digging in. 

- When NGINX starts up it does a DNS lookup for the upstreams.
- By default NGINX stores the IPs and sends traffic upstream via the IPv4/IPv6 address not the DNS record. 
- NGINX respects DNS TTLs and will not revalidate on a failed request or timeout by default. 
- When you redeploy a service on hosts like railway or even render your services often receive a new IP
- With that in mind, on redeploy NGINX continues forward requests to a now dead upstream's IPv4/IPv6 address. 

You can get around that by either 
- Restarting NGINX on every upstream redeploy (that sucks)
- Paying for NGINX plus which apparenly has some dynamic DNS lookup stuff to fix this 
- Setting a custom `resolver` key (`resolver <dns-server> valid=Ns`) which tells NGINX to ignore TTL and revalidate DNS records every N seconds (can still result in failed requests)

All of these are subpar solutions in my opnion. With that in mind, I moved to Caddy Server which doesn't seem to have this issue.

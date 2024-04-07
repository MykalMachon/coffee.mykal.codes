# # custom NGINX image to support custom configuration
# FROM nginx:stable-alpine3.17
# COPY ./nginx.conf /etc/nginx/nginx.conf

FROM caddy:2.7.6-alpine

COPY Caddyfile /etc/caddy/Caddyfile
RUN caddy fmt --overwrite /etc/caddy/Caddyfile

CMD caddy run --config /etc/caddy/Caddyfile --adapter caddyfile 2>&1
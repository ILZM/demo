FROM node:16-alpine AS build_stage

WORKDIR /app
COPY package*.json .
RUN npm ci

COPY tsconfig*.json .
COPY .env .
COPY public public
COPY tools tools
COPY src src

RUN npm run build

FROM caddy:2.6.4-alpine

COPY caddy/Caddyfile /etc/caddy/Caddyfile
COPY --from=build_stage /app/dist /usr/share/caddy/html

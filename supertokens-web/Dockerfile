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

FROM nginx:stable-alpine3.17-slim

COPY server/nginx.conf /etc/nginx/nginx.conf
COPY --from=build_stage /app/dist /usr/share/nginx/html

RUN chown nginx.nginx /usr/share/nginx/html/ -R

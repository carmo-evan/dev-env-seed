FROM node:12.17.0-alpine as node

WORKDIR /root

COPY ./ /root
RUN apk update && apk add bash && apk add curl
RUN npm install -g @angular/cli
RUN npm install
RUN ng version
ENTRYPOINT ng serve --host 0.0.0.0
# Stage 2
# FROM nginx:1.13.12-alpine


# COPY --from=node /usr/src/app/dist/client /usr/share/nginx/html

# COPY ./nginx.conf /etc/nginx/conf.d/default.conf

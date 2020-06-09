FROM nginx:alpine

COPY ./ /root/client

WORKDIR /root/client

RUN ls
RUN apk update && apk add --update nodejs && apk add --update npm
RUN npm install
RUN npm install -g @angular/cli
RUN ng version
RUN ng build --output-path=/usr/share/nginx/html
CMD ["./start.sh"]

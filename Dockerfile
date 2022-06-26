FROM alpine:latest

EXPOSE 8080

ADD zeus /bin/zeus

CMD ["zeus"]
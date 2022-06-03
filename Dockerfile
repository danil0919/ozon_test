FROM golang
WORKDIR /app
COPY . . 
RUN make build

EXPOSE 8080
CMD ["/app/apiserver"]
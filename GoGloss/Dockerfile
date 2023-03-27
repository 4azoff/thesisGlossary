# Start from a base Go image
FROM golang:1.18-alpine

# Set the working directory
WORKDIR /app

# Copy the Go files into the container
COPY ./ /app/
COPY /cmd/mindMap.html /app/
COPY /cmd/terms.html /app/
COPY /cmd/static/ /app/static/

RUN go mod download
RUN go build -o ./termsGo ./cmd

ENTRYPOINT [ "./termsGo" ]

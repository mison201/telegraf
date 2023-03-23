
# Build Stage
FROM golang:1.19-alpine AS BuildStage

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build ./cmd/telegraf

# Deploy Stage
FROM alpine:latest

WORKDIR /app

COPY --from=BuildStage /app/telegraf /app/telegraf
COPY --from=BuildStage /app/sample-config/telegraf.conf /app/config/telegraf.conf
COPY --from=BuildStage /app/sample-config/input-nginx-vts.conf /app/config/input-nginx-vts.conf
COPY --from=BuildStage /app/input-nginx-vts.tmpl /app/input-nginx-vts.tmpl
COPY --from=BuildStage /app/run.sh /app/run.sh

ENTRYPOINT ["/bin/sh", "-c", "/app/run.sh"]
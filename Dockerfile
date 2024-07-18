FROM node:20.15-alpine3.20 AS build-fe

WORKDIR /app/frontend
COPY ./frontend/package.json .
COPY ./frontend/package-lock.json .
# RUN NODE_ENV=production npm ci .
RUN npm ci .
COPY ./frontend .
RUN npm run build

FROM golang:1.22.5-alpine3.20 AS build-be

WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN rm -rf frontend
COPY --from=build-fe /app/frontend/dist /go/src/app/frontend/dist
RUN GOOS=linux go build -ldflags="-s" -o /go/bin/app .

# Create minimal /etc/passwd wiht appuser
# RUN echo "appuser:x:10001:10001:App User:/:/sbin/nologin" > /etc/minimal-passwd

FROM scratch
COPY --from=build-be /go/bin/app /go/bin/app
# COPY --from=build-be /etc/minimal-passwd /etc/passwd
# USER appuser

EXPOSE 3000
ENTRYPOINT [ "/go/bin/app" ]

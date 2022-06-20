FROM node:18.4.0-bullseye AS frontend-builder
WORKDIR /frontend
COPY package.json package.json
COPY package-lock.json package-lock.json
COPY Makefile Makefile
COPY patches patches
RUN make dep-frontend
COPY . .
RUN make build-frontend

FROM golang:1.18-bullseye AS backend-builder
WORKDIR /backend
COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile
RUN make dep-backend
COPY . .
RUN make build-backend
COPY --from=frontend-builder /frontend/frontend/build ./frontend/build
RUN make build

FROM python:3.10.5-slim
RUN apt-get update \
    && apt-get install --no-install-recommends -y make \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /readability
COPY requirements.txt requirements.txt
COPY Makefile Makefile
RUN make dep-readability
WORKDIR /app
RUN rm -rf /readability
COPY --from=backend-builder /backend/bin/hackernews /app/hackernews

EXPOSE 8080
CMD ["/app/hackernews"]
# Frontend
FROM node:lts-alpine AS frontend
WORKDIR /frontend

COPY frontend/package.json ./
RUN npm install

COPY frontend/ .
RUN npm run build


# Backend
FROM golang:1.22 AS base
WORKDIR /yacc

COPY go.mod go.sum ./

RUN go mod download

COPY main.go ./
COPY backend/ ./backend/

RUN CGO_ENABLED=0 GOOS=linux go build


# Merge front and back end
COPY --from=frontend /frontend/dist /frontend/dist

ENV PORT=8000

WORKDIR /

CMD ["/yacc/yacc"]

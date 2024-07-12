FROM golang:1.22-alpine AS build-stage
WORKDIR /admin_svc
COPY ./ /admin_svc
RUN mkdir -p /admin_svc/build
RUN go mod download
RUN go build -v -o /admin_svc/build/api ./cmd


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /admin_svc/build/api /
COPY --from=build-stage /admin_svc/.env /
EXPOSE 8888
CMD [ "/api" ]
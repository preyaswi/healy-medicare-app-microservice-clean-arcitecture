FROM golang:1.22-alpine AS build-stage
WORKDIR /doctor_svc
COPY ./ /doctor_svc
RUN mkdir -p /doctor_svc/build
RUN go mod download
RUN go build -v -o /doctor_svc/build/api ./cmd


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /doctor_svc/build/api /
COPY --from=build-stage /doctor_svc/.env /
EXPOSE 8888
CMD [ "/api" ]
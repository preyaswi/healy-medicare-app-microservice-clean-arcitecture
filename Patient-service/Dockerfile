FROM golang:1.22-alpine AS build-stage
WORKDIR /patient_svc
COPY ./ /patient_svc
RUN mkdir -p /patient_svc/build
RUN go mod download
RUN go build -v -o /patient_svc/build/api ./cmd


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /patient_svc/build/api /
COPY --from=build-stage /patient_svc/.env /
EXPOSE 50051
CMD [ "/api" ]
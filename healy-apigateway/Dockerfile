FROM golang:1.22-alpine AS build-stage
WORKDIR /healy
COPY ./ /healy
RUN mkdir -p /healy/build
RUN go mod download
RUN go build -v -o /healy/build/api ./cmd/main.go


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /healy/build/api /
COPY --from=build-stage /healy/template/ /template/
COPY --from=build-stage /healy/.env /
EXPOSE 8000
CMD [ "/api" ]
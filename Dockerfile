FROM golang:1.15.3-alpine3.12 as builder

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build \
    -o /app .

FROM public.ecr.aws/lambda/go

COPY --from=builder app ${LAMBDA_TASK_ROOT}

CMD [ "app" ]
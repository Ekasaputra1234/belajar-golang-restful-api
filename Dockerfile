FROM golang:1.23
ARG CI_SERVER_HOST
ARG CI_JOB_TOKEN
ARG CI_COMMIT_TAG
ARG TAG_DELETE_TOKEN
ARG TAG

WORKDIR /app
COPY . .

RUN echo "CI_SERVER_HOST: $CI_SERVER_HOST"
RUN echo "CI_JOB_TOKEN: $CI_JOB_TOKEN"
RUN git config --global url."https://gitlab-ci-token:${CI_JOB_TOKEN}@${CI_SERVER_HOST}".insteadOf https://${CI_SERVER_HOST}
RUN export GOPRIVATE=${CI_SERVER_HOST}

RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install github.com/go-critic/go-critic/cmd/gocritic@latest
# RUN gocritic check $(go list ./... | grep -v /vendor/)
# RUN staticcheck $(go list ./... | grep -v /vendor)

RUN go fmt $(go list ./... | grep -v /vendor/)
RUN go vet $(go list ./... | grep -v /vendor/)
RUN go test -race $(go list ./... | grep -v /vendor/)
RUN go build -race -ldflags "-extldflags '-static'" -o /app/main

CMD ["/app/main"]
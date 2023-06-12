FROM fuzzers/go-fuzz:1.2.0 as builder

#Update go

RUN rm -rf /usr/local/go
RUN curl -OL https://golang.org/dl/go1.20.3.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz
RUN export PATH=$PATH:/usr/local/go/bin



RUN go install github.com/dvyukov/go-fuzz/go-fuzz-build@latest

WORKDIR /tmp/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
WORKDIR /tmp/app/fuzzing
RUN go-fuzz-build -libfuzzer -o username_fuzzer.a
RUN clang -fsanitize=fuzzer username_fuzzer.a -o username_fuzzer.libfuzzer

#entrypoint? idk docker.

FROM ubuntu:20.04
COPY --from=builder /tmp/app/fuzzing/username_fuzzer.libfuzzer /

CMD ["username_fuzzer.libfuzzer"]


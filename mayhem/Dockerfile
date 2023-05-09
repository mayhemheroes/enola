FROM fuzzers/go-fuzz:1.2.0

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
RUN go-fuzz-build -libfuzzer -o fuzzer.a
RUN clang -fsanitize=fuzzer fuzzer.a -o fuzzer.libfuzzer

#entrypoint? idk docker.

CMD ["/tmp/app/fuzzing/fuzzer.libfuzzer"]


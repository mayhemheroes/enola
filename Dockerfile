FROM fuzzers/go-fuzz:1.2.0 
WORKDIR /tmp/app
COPY go.mod .
COPY go.sum .
RUN go mod download
# COPY fuzzing .
COPY . .
WORKDIR /tmp/app/cmd/enola/fuzzing
RUN go-fuzz-build -libfuzzer -o fuzzer.a
RUN clang -fsanitize=fuzzer fuzzer.a -o fuzzer.libfuzzer
RUN mv fuzzer.libfuzzer /go

#entrypoint? idk docker.
ENTRYPOINT []
CMD ["/go/fuzzer.libfuzzer"]


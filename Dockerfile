FROM golang:1.6

# RUN mkdir /go/src/gentest2 && /go/bin 
ADD . /go/src/gentest2/

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN cd /go/src/gentest2 && go get . &&  go install .

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/gentest2 serve --config=/go/src/gentest2/config/config.json

# Document that the service listens on port 8080.
EXPOSE 3000
FROM ubuntu:latest
LABEL authors="Alex"

ENTRYPOINT ["top", "-b"]
FROM alpine:3.6

ADD tmp/_output/bin/bruslan-operator /usr/local/bin/bruslan-operator

RUN adduser -D bruslan-operator
USER bruslan-operator

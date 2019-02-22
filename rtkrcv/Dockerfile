from alpine:latest as builder

RUN apk update \
    && apk add --virtual build-dependencies \
        build-base \
        gcc

COPY gcc/ /rtkrcv/gcc/
COPY src/ /rtkrcv/src/
COPY rtk.conf rtkrcv.c vt.c vt.h /rtkrcv/

RUN cd /rtkrcv/gcc && make clean && make

# step 2

FROM alpine:latest

COPY single.conf /single.conf

# Copy our static executable
COPY --from=builder /rtkrcv/gcc/rtkrcv /rtkrcv

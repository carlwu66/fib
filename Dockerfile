FROM ubuntu
EXPOSE 8001

COPY ./fib /
CMD ["/fib"]

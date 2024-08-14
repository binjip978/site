FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates
COPY site /bin/site
ENTRYPOINT ["/bin/site"]

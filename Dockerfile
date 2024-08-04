FROM debian:bookworm-slim
COPY site /bin/site
ENTRYPOINT ["/bin/site"]

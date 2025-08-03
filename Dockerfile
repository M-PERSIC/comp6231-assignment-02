FROM docker.io/library/ubuntu:24.04
COPY comp6231-assignment-02 /comp6231-assignment-02
ENTRYPOINT ["/comp6231-assignment-02"]
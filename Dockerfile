FROM scratch
LABEL maintainer="Matt Blewitt <matt@blwt.io>"
COPY ["/cross/linux/amd64/hookshot", "/"]
CMD ["/hookshot"]

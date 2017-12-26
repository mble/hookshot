FROM scratch
LABEL maintainer="Matt Blewitt <matt@blwt.io>"
COPY ["/cross/linux/amd64/deployhook", "/"]
EXPOSE 80
CMD ["/deployhook"]

FROM centos/systemd

MAINTAINER "Akash Gautam" <akash.gautam@velotio.com>

COPY crd-blog /

ENTRYPOINT ["/crd-blog"]
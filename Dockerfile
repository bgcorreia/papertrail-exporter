FROM        quay.io/prometheus/busybox:latest
MAINTAINER  Hiroaki Murayama <hiroaki.murayama@mixi.co.jp>

COPY papertrail-exporter /bin/papertrail-exporter

EXPOSE      9098
ENTRYPOINT  [ "/bin/papertrail-exporter" ]

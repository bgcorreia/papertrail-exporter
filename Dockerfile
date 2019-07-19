FROM        quay.io/prometheus/busybox:latest
MAINTAINER  Hiroaki Murayama <hiroaki.murayama@mixi.co.jp>

COPY papertrail_exporter /bin/papertrail_exporter

EXPOSE      9098
ENTRYPOINT  [ "/bin/papertrail_exporter" ]

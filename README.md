# Papertrail Exporter

GitHub: https://github.com/oppai/papertrail-exporter

## Running

You can run via docker with:

```
docker run -d -p 9098:9098 --name papertrail-exporter \
  -web.listen-address=":9098" \
  -web.telemetry-path="/metrics" \
  -config.token="xxxxxxxx" \
  oppai/papertrail-exporter:master
```

You'll need to customize the docker image or use the binary on the host system
to install tools such as curl for certain scenarios.

## Proving

TBD

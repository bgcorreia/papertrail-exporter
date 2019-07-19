# Papertrail Exporter

GitHub: https://github.com/oppai/papertrail-exporter

## Running

You can run via docker with:

```
docker run -p 9098:9098 kodam/papertrail-exporter:1.0.0 \
    -web.listen-address=":9098" \
    -web.telemetry-path="/metrics" \
    -config.token="xxxxxxxxxx"
```

You'll need to customize the docker image or use the binary on the host system
to install tools such as curl for certain scenarios.

## Proving

TBD

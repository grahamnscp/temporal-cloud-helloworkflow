# temporal-cloud-helloworkflow
Basic Hello Workflow for use with Temporal Cloud

```
go mod tidy
```

Set environment variables:
* TEMPORAL_HOST_URL
* TEMPORAL_NAMESPACE
* TEMPORAL_TLS_CERT
* TEMPORAL_TLS_KEY


To run, start one or more workers in separate terminals:
```
go run worker/main.go -target-host $TEMPORAL_HOST_URL -namespace $TEMPORAL_NAMESPACE -client-cert $TEMPORAL_TLS_CERT -client-key $TEMPORAL_TLS_KEY
```

Start/run the workflow:
```
go run starter/main.go -target-host $TEMPORAL_HOST_URL -namespace $TEMPORAL_NAMESPACE -client-cert $TEMPORAL_TLS_CERT -client-key $TEMPORAL_TLS_KEY
```


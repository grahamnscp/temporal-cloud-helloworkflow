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
* TEMPORAL_INSECURE_SKIP_VERIFY=false
* ENCRYPT_PAYLOAD=true
* DATACONVERTER_ENCRYPTION_KEY_ID=secret

To run, start one or more workers in separate terminals:
```
go run worker/main.go
```

Start/run the workflow:
```
go run starter/main.go
```

Note: Added upserting a custom search attribute, please ensure that the attribute is present on the temporal server before running sample.  
* Attribute name: CustomStringField  
* Attribute type: Text  


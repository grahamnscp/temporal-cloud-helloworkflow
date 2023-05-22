package temporalcloudhelloworkflow

import (
  "fmt"
  "os"
  "log"
  "strconv"
  "crypto/tls"
  "crypto/x509"

  "go.temporal.io/sdk/client"
  "go.temporal.io/sdk/converter"

  dataconverter "temporalcloudhelloworkflow/dataconverter"
)


/* LoadClientOption - Return client options for Temporal Cloud */
func LoadClientOption () (client.Options, error) {

  // Read env variables
  targetHost := os.Getenv("TEMPORAL_HOST_URL")
  namespace  := os.Getenv("TEMPORAL_NAMESPACE")
  clientCert := os.Getenv("TEMPORAL_TLS_CERT")
  clientKey  := os.Getenv("TEMPORAL_TLS_KEY")

  // Optional:
  serverRootCACert := os.Getenv("TEMPORAL_SERVER_ROOT_CA_CERT")
  serverName       := os.Getenv("TEMPORAL_SERVER_NAME")

  insecureSkipVerify, _ := strconv.ParseBool(os.Getenv("TEMPORAL_INSECURE_SKIP_VERIFY"))

  encyptPayload, _ := strconv.ParseBool(os.Getenv("ENCRYPT_PAYLOAD"))

  log.Println("LoadClientOption:", targetHost, namespace, clientCert, clientKey, serverRootCACert, serverName, insecureSkipVerify, encyptPayload)

  // Load client cert
  cert, err := tls.LoadX509KeyPair(clientCert, clientKey)
  if err != nil {
    return client.Options{}, fmt.Errorf("failed loading client cert and key: %w", err)
  }

  // Load server CA if given
  var serverCAPool *x509.CertPool
  if serverRootCACert != "" {
    serverCAPool = x509.NewCertPool()
    b, err := os.ReadFile(serverRootCACert)
    if err != nil {
      return client.Options{}, fmt.Errorf("failed reading server CA: %w", err)
    } else if !serverCAPool.AppendCertsFromPEM(b) {
      return client.Options{}, fmt.Errorf("server CA PEM file invalid")
    }
  }

  // Return client options
  if encyptPayload {

    return client.Options {
      HostPort:  targetHost,
      Namespace: namespace,
      ConnectionOptions: client.ConnectionOptions {
        TLS: &tls.Config {
          Certificates:       []tls.Certificate{cert},
          RootCAs:            serverCAPool,
          ServerName:         serverName,
          InsecureSkipVerify: insecureSkipVerify,
        },
      },
      // Set DataConverter to ensure that workflow inputs and results are
      // encrypted/decrypted as required.
      DataConverter: dataconverter.NewEncryptionDataConverter(
        converter.GetDefaultDataConverter(), 
          dataconverter.DataConverterOptions{KeyID: os.Getenv("DATACONVERTER_ENCRYPTION_KEY_ID")},
      ),
    }, nil

  } else {

    return client.Options {
      HostPort:  targetHost,
      Namespace: namespace,
      ConnectionOptions: client.ConnectionOptions {
        TLS: &tls.Config {
          Certificates:       []tls.Certificate{cert},
          RootCAs:            serverCAPool,
          ServerName:         serverName,
          InsecureSkipVerify: insecureSkipVerify,
        },
      },
    }, nil
  }
}


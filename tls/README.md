To generate a self-signed certificate you'll need to find where the source code for Go standard library is installed and then run the following command, changing the path to the generate_cert.go file as needed:

```bash
  go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```
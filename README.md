# factorysign

An example Go web service that digitally signs basic device details.

## Install
Go get it:

  ```bash
	$ go get github.com/slimjim777/factorysign
  ```

Run it:
  ```bash
	$ cd factorysign
  $ go run server.go -config=/path/to/settings.yaml
  ```

## API Methods

### /v1/sign (POST)
> Clear-sign the device details.

Takes the details from the device, formats the data and clear-signs it.

#### Input message

  ```json
  {
    "serial":"M12345/LN",
    "model":"Sekret Device",
    "publickey":"abcd1234"
  }
  ```
- model: the name of the device (string)
- serial: serial number of the device (string)
- publickey: the device's public key (string)

#### Output message  

```json
{
  "success":true,
  "message":"",
  "signature":"-----BEGIN PGP SIGNED MESSAGE-----\nHash: SHA256\n\nabcd1234||Sekret Device||M12345/LN\n-----BEGIN PGP SIGNATURE-----\n\nwsFcBAEBCA...A5LT\n-----END PGP SIGNATURE-----"}
```
- success: whether the submission was successful (bool)
- message: error message from the submission (string)
- signature: the formatted, clear-signed data (string)

#### Example
```bash
curl -X POST -d '{"serial":"M12345/LN", "model":"Sekret Device", "publickey":"abcd1234"}' http://localhost:8080/v1/sign
```

# Identity Vault

A Go web service that digitally signs device assertion details.

## Install
Go get it:

  ```bash
	$ go get github.com/ubuntu-core/identity-vault
  ```

Run it:
  ```bash
	$ cd identity-vault
  $ go run server.go -config=/path/to/settings.yaml
  ```

## API Methods

### /1.0/version (GET)
> Return the version of the identity vault service.

#### Output message

```json
{
  "version":"0.1.0",
}
```
  - version: the version of the identity vault service (string)


### /1.0/sign (POST)
> Clear-sign the device identity details.

Takes the details from the device, formats the data and clear-signs it.

#### Input message

  ```json
  {
    "serial":"M12345/LN",
    "brand": "System",
    "model":"Device 1000",
    "device-key":"ssh-rsa abcd1234",
    "revision": 2
  }
  ```
- brand: the name of the manufacturer (string)
- model: the name of the device (string)
- serial: serial number of the device (string)
- device-key: the type and public key of the device (string)
- revision: the revision of the device (integer)

#### Output message

```json
{
  "success":true,
  "message":"",
  "signature":"-----BEGIN PGP SIGNED MESSAGE-----\nHash: SHA256\n\ntype: device\nbrand: Device 1000\nmodel: Device 1000\nserial: M12345/LN\ntimestamp: 2016-02-03 17:22:59.93489652 +0000 UTC\nrevision: 2\ndevice-key: ssh-rsa abcd1234\n-----BEGIN PGP SIGNATURE-----\n\nwsFcBAEBCA ... A5LT\n-----END PGP SIGNATURE-----"}
```
- success: whether the submission was successful (bool)
- message: error message from the submission (string)
- identity: the formatted, clear-signed data (string)

#### Example
```bash
curl -X POST -d '{"serial":"M12345/LN","brand":"System",  "model":"Device 1000", "revision": 2, "device-key":"rsa-ssh abcd1234"}' http://localhost:8080/1.0/sign
```

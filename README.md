# ScaleIO exporter
The ScaleIO exporter is a Prometheus exporter for PowerFlex( ex: ScaleIO ) written in Go.
For now, This script can collect few metrics, such as a capacity of storage.

# Usage
```
# for linux
$ GOOS=linux go build -o scaleio_exporter cmd/main.g

# for MacOS
$ GOOS=darwin go build -o scaleio_exporter cmd/main.go


$ ./sclaeio_exporter --username username --password password --host IPaddress
```
# test
TBD

## Options
| Name     | Flag        | Env vars  | Default | Description                                     |
|----------|-------------|-----------|---------|-------------------------------------------------|
| Username | username, u | USERNAME  | -       | Username for ScaleIO                            |
| Password | password    | PASSWORD  | -       | Password for ScaleIO                            |
| IPAddr   | ipaddr, i   | IPADDRESS | -       | Specify ScaleIO Host IP Address                 |
| Refresh  | refresh, r  | REFRESH   | 300     | Refresh time fetch GitHub billing report in sec |
| Port     | port, p     | PORT      | 10000   | Exporter port                                   |
| Insecure | insecure, k | INSECURE  | true    | Verify Certificate or not                       |

# License
Apache License 2.0, see LICENSE.

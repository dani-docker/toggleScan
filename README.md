# disableScan
This tool is used to modify DTR config stored in UCP and disable BatchScanningDataEnabled

# Usage
```
[centos@dlouca ~]$ docker run -it dlouca/disablescan:latest -a 34.205.41.253 -u admin
Password for UCP user admin: 
***** Please save this output; Original config 


 *****
{
  "Registries": [
    {
      "hostAddress": "35.171.225.23",
      "serviceID": "60e847b5-f115-45bb-9e4a-871f840e8d61",
      "caBundle": "-----BEGIN CERTIFICATE-----\nMIIB3zCCAYWgAwIBAgIQO/4zU+/lPSnPfxwT8lsxrzAKBggqhkjOPQQDAjBHMQsw\nCQYDVQQGEwJVUzEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEPMA0GA1UEChMGRG9j\na2VyMQ8wDQYDVQQLEwZEb2NrZXIwHhcNMTkwMzA4MjEzNTM2WhcNMjAwMzA3MjEz\nNTM2WjBHMQswCQYDVQQGEwJVUzEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEPMA0G\nA1UEChMGRG9ja2VyMQ8wDQYDVQQLEwZEb2NrZXIwWTATBgcqhkjOPQIBBggqhkjO\nPQMBBwNCAATkr0I6l5KXexi7JSVBtRd/L3gZbiF2X98r5YRZ0gidclg4mJaYKoUy\n3Q/PhkdEnaGiLd3QgVw1MrbzPnF/l9pYo1MwUTAOBgNVHQ8BAf8EBAMCAqQwDwYD\nVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUBh3xDlpT5Tb9Z+4QzhusWFBhTAcwDwYD\nVR0RBAgwBocEI6vhFzAKBggqhkjOPQQDAgNIADBFAiEA/vZh/SkyJlsmAIjyn0ba\nsSqwvwEXdpf/UOrwWZYd8m4CIBQo/dpWbuHbe5KN17awuRuY1CFRxgJ2wRtxhgie\n1NbD\n-----END CERTIFICATE-----\n",
      "batchScanningDataEnabled": true
    },
    {
      "hostAddress": "35.171.225.24",
      "serviceID": "60e847b5-f115-45bb-9e4a-871f840e8d61",
      "caBundle": "-----BEGIN CERTIFICATE-----\nMIIB3zCCAYWgAwIBAgIQO/4zU+/lPSnPfxwT8lsxrzAKBggqhkjOPQQDAjBHMQsw\nCQYDVQQGEwJVUzEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEPMA0GA1UEChMGRG9j\na2VyMQ8wDQYDVQQLEwZEb2NrZXIwHhcNMTkwMzA4MjEzNTM2WhcNMjAwMzA3MjEz\nNTM2WjBHMQswCQYDVQQGEwJVUzEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEPMA0G\nA1UEChMGRG9ja2VyMQ8wDQYDVQQLEwZEb2NrZXIwWTATBgcqhkjOPQIBBggqhkjO\nPQMBBwNCAATkr0I6l5KXexi7JSVBtRd/L3gZbiF2X98r5YRZ0gidclg4mJaYKoUy\n3Q/PhkdEnaGiLd3QgVw1MrbzPnF/l9pYo1MwUTAOBgNVHQ8BAf8EBAMCAqQwDwYD\nVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUBh3xDlpT5Tb9Z+4QzhusWFBhTAcwDwYD\nVR0RBAgwBocEI6vhFzAKBggqhkjOPQQDAgNIADBFAiEA/vZh/SkyJlsmAIjyn0ba\nsSqwvwEXdpf/UOrwWZYd8m4CIBQo/dpWbuHbe5KN17awuRuY1CFRxgJ2wRtxhgie\n1NbD\n-----END CERTIFICATE-----\n",
      "batchScanningDataEnabled": true
    }
  ]
}
***** Successfuly disabled scan to all DTR instances. New Config *****
{
  "Registries": [
    {
      "hostAddress": "35.171.225.23",
      "serviceID": "60e847b5-f115-45bb-9e4a-871f840e8d61",
      "caBundle": "-----BEGIN CERTIFICATE-----\nMIIB3zCCAYWgAwIBAgIQO/4zU+/lPSnPfxwT8lsxrzAKBggqhkjOPQQDAjBHMQsw\nCQYDVQQGEwJVUzEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEPMA0GA1UEChMGRG9j\na2VyMQ8wDQYDVQQLEwZEb2NrZXIwHhcNMTkwMzA4MjEzNTM2WhcNMjAwMzA3MjEz\nNTM2WjBHMQswCQYDVQQGEwJVUzEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEPMA0G\nA1UEChMGRG9ja2VyMQ8wDQYDVQQLEwZEb2NrZXIwWTATBgcqhkjOPQIBBggqhkjO\nPQMBBwNCAATkr0I6l5KXexi7JSVBtRd/L3gZbiF2X98r5YRZ0gidclg4mJaYKoUy\n3Q/PhkdEnaGiLd3QgVw1MrbzPnF/l9pYo1MwUTAOBgNVHQ8BAf8EBAMCAqQwDwYD\nVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUBh3xDlpT5Tb9Z+4QzhusWFBhTAcwDwYD\nVR0RBAgwBocEI6vhFzAKBggqhkjOPQQDAgNIADBFAiEA/vZh/SkyJlsmAIjyn0ba\nsSqwvwEXdpf/UOrwWZYd8m4CIBQo/dpWbuHbe5KN17awuRuY1CFRxgJ2wRtxhgie\n1NbD\n-----END CERTIFICATE-----\n",
      "batchScanningDataEnabled": false
    },
    {
      "hostAddress": "35.171.225.24",
      "serviceID": "60e847b5-f115-45bb-9e4a-871f840e8d61",
      "caBundle": "-----BEGIN CERTIFICATE-----\nMIIB3zCCAYWgAwIBAgIQO/4zU+/lPSnPfxwT8lsxrzAKBggqhkjOPQQDAjBHMQsw\nCQYDVQQGEwJVUzEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEPMA0GA1UEChMGRG9j\na2VyMQ8wDQYDVQQLEwZEb2NrZXIwHhcNMTkwMzA4MjEzNTM2WhcNMjAwMzA3MjEz\nNTM2WjBHMQswCQYDVQQGEwJVUzEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEPMA0G\nA1UEChMGRG9ja2VyMQ8wDQYDVQQLEwZEb2NrZXIwWTATBgcqhkjOPQIBBggqhkjO\nPQMBBwNCAATkr0I6l5KXexi7JSVBtRd/L3gZbiF2X98r5YRZ0gidclg4mJaYKoUy\n3Q/PhkdEnaGiLd3QgVw1MrbzPnF/l9pYo1MwUTAOBgNVHQ8BAf8EBAMCAqQwDwYD\nVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUBh3xDlpT5Tb9Z+4QzhusWFBhTAcwDwYD\nVR0RBAgwBocEI6vhFzAKBggqhkjOPQQDAgNIADBFAiEA/vZh/SkyJlsmAIjyn0ba\nsSqwvwEXdpf/UOrwWZYd8m4CIBQo/dpWbuHbe5KN17awuRuY1CFRxgJ2wRtxhgie\n1NbD\n-----END CERTIFICATE-----\n",
      "batchScanningDataEnabled": false
    }
  ]
}

```

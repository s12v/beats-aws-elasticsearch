
beats plugin experiment (copy-paste from https://github.com/elastic/beats/tree/master/libbeat/outputs/elasticsearch)

```
 go build -buildmode=plugin ./plugins/aws_elasticsearch
 ./filebeat -e -plugin aws_elasticsearch.so -d '*'
```

config:
```
filebeat.prospectors:
- type: log
  enabled: true
  paths:
    - ./*.log

output.aws_elasticsearch:
  hosts: ["https://search-test2-secret.eu-central-1.es.amazonaws.com:443"]
  aws:
    enabled: true
    region: eu-central-1
```

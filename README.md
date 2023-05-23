# ip2country

Download mmdb from:
1. https://ipinfo.io/data/free/country.mmdb
2. https://github.com/Dreamacro/maxmind-geoip/releases/latest/download/Country.mmdb


```
#!/bin/bash

while :
do
        echo `date && dig www.amazon.co.jp @192.168.1.1 +short | tail -1 | ./ip2country && dig www.amazon.co.jp @192.168.1.1 -p 50353 +short | tail -1 | ./ip2country && dig www.amazon.co.jp @192.168.2.1 +short | tail -1 | ./ip2country && dig www.amazon.co.jp +short | tail -1 | ./ip2country` | tr -d '\n'
        echo ""
        sleep 30
done
```

```
dig www.taobao.com +short | tail -1 | ./ip2country

dig www.amazon.co.jp +short | tail -1 | ./ip2country

dig www.amazon.co.jp @192.168.100.1 -p 50353 +tcp +short | tail -1 | ./ip2country
```

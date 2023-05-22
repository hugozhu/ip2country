# ip2country

Download mmdb from:
1. https://ipinfo.io/data/free/country.mmdb
2. https://github.com/Dreamacro/maxmind-geoip/releases/latest/download/Country.mmdb

```
dig www.taobao.com +short | tail -1 | ./ip2country

dig www.amazon.co.jp +short | tail -1 | ./ip2country

dig www.amazon.co.jp @192.168.100.1 -p 50353 +tcp +short | tail -1 | ./ip2country
```
#!/bin/bash

# Check if the correct number of arguments is provided
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <domain_name> <hosts_file>"
    exit 1
fi

domain=$1
hosts_file=$2

#ip_to_check=`dig $domain @9.9.9.9 +tcp +short | tail -1`
ip_to_check=`dig $domain +short | tail -1`

line=`curl --max-time 30 --resolve "$domain:443:$ip_to_check" -s -o /dev/null -w '%{http_code} %{remote_ip} %{size_download} %{time_total}'  -H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:66.0) Gecko/20100101 Firefox/66.0' -H 'DNT: 1' -H 'Connection: close' -H 'Upgrade-Insecure-Requests: 1'  https://$domain/`

http_status_code=$(echo "$line" | awk '{print $1}')
content_length=$(echo "$line" | awk '{print $3}')
time_cost=$(echo "$line" | awk '{print $4}')

if [[ "$http_status_code" != "200" ]]; then
    echo "$ip_to_check => $http_status_code $content_length $time_cost"
    #bad ip, ignore
    exit 0
fi

#valid ip, save to hosts file
if grep -q "^$ip_to_check" "$hosts_file"; then
    #echo "IP $ip_to_check is in the hosts file."
    exit 0
else
    echo "$ip_to_check $domain" >> $hosts_file
    lines=$(cat $hosts_file | tail -n 3)
    echo "$lines" > $hosts_file  
fi
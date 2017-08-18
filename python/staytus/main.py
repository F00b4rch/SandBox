"""This program curl api Staytus"""
#!/usr/bin/python

import time
import requests

def main():
    """Main content"""

    # Website
    urls = {'http://website1.com' : 'website1-com',
            'http://website2.com' : 'website2-com',
            'http://website3.com' : 'website3-com'}

    # Staytus
    ndd = "https://status.website.com"
    url = "/api/v1/services/set_status"
    headers = {'X-Auth-Token': 'XXXXXXXXXXXXXX',
               'X-Auth-Secret': 'XXXXXXXXXXXXXXX'}

    # -- List of statuses --
    # operational
    # degraded-performance
    # partial-outage
    # major-outage
    # maintenance
    # offline

    for k, value in urls.items():
        try:
        
            req_site = requests.head(k)
            total_sec = requests.get(k, timeout=30).elapsed.total_seconds()
            
            if req_site.status_code != 200:
                params = {'service': value, 'status':'degraded-performance'}
                requests.get(ndd + url, headers=headers, json=params)
            elif req_site.status_code == 200:
                params = {'service': value, 'status':'operational'}
                requests.get(ndd + url, headers=headers, json=params)
            else:
                pass
            
            if total_sec > 30:
                params = {'service': value, 'status':'degraded-performance'}
                requests.get(ndd + url, headers=headers, json=params)
        
        except Exception as x:
            print(x)
            
while True:
    main()
    time.sleep(5)
    

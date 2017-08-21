"""This program curl api Staytus"""
#!/usr/bin/python

import time
import requests

def main():
    """Main content"""

    # Website
    urls = {'http://my.site1.com' : 'my-site1-com',
            'http://my.site2.com' : 'my-site2-com',
            'http://my.site3.com' : 'my-site2-com'
           }

    # Staytus
    ndd = "https://status.my-site.com"
    url = "/api/v1/services/set_status"
    headers = {'X-Auth-Token': 'XXX-XXX-XXX-XXX-XXX',
               'X-Auth-Secret': 'xxxxxxxxxxxxxxxxxxxxx'}

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
            else:
                params = {'service': value, 'status':'operational'}
                requests.get(ndd + url, headers=headers, json=params)

        except requests.RequestException as err:
            print(err)

while True:
    main()
    time.sleep(5)

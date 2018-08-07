# crt_sheep

# UPDATE AUG 6, 2018
All is good with crt.sh again. If you have any issues in the future refer to the Aug 4 update and ensure the domain isn't currently down


# UPDATE AUG 4, 2018
I was wondering why my tool was erroring out so I checked out the https://crt.sh website and realized it was 404'ing with the following showing up when cURL'd . Until further notice this tool is sadly broken as it depends on crt.sh

```
*   Trying 178.255.82.12...
* TCP_NODELAY set
* Connected to crt.sh (178.255.82.12) port 443 (#0)
* ALPN, offering h2
* ALPN, offering http/1.1
* successfully set certificate verify locations:
*   CAfile: /etc/ssl/certs/ca-certificates.crt
  CApath: /etc/ssl/certs
* TLSv1.2 (OUT), TLS handshake, Client hello (1):
* TLSv1.2 (IN), TLS handshake, Server hello (2):
* TLSv1.2 (IN), TLS handshake, Certificate (11):
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
* TLSv1.2 (IN), TLS handshake, Server finished (14):
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
* TLSv1.2 (OUT), TLS change cipher, Client hello (1):
* TLSv1.2 (OUT), TLS handshake, Finished (20):
* TLSv1.2 (IN), TLS handshake, Finished (20):
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256
* ALPN, server accepted to use http/1.1
* Server certificate:
*  subject: serialNumber=04058690; jurisdictionC=GB; businessCategory=Private Organization; C=GB; postalCode=M5 3EQ; ST=Greater Manchester; L=Salford; street=Trafford Road; street=Exchange Quay; street=3rd Floor, 26 Office Village; O=COMODO CA Limited; OU=COMODO EV Multi-Domain SSL; CN=crt.sh
*  start date: Jul  1 00:00:00 2016 GMT
*  expire date: Sep 30 23:59:59 2018 GMT
*  subjectAltName: host "crt.sh" matched cert's "crt.sh"
*  issuer: C=GB; ST=Greater Manchester; L=Salford; O=COMODO CA Limited; CN=COMODO RSA Extended Validation Secure Server CA
*  SSL certificate verify ok.
> GET / HTTP/1.1
> Host: crt.sh
> User-Agent: curl/7.58.0
> Accept: */*
>
< HTTP/1.1 404 Not Found
HTTP/1.1 404 Not Found
< Date: Sat, 04 Aug 2018 13:50:08 GMT
Date: Sat, 04 Aug 2018 13:50:08 GMT
< Server: Apache
Server: Apache
< Content-Length: 255
Content-Length: 255
< Content-Type: text/html; charset=iso-8859-1
Content-Type: text/html; charset=iso-8859-1

<
<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
<html><head>
<title>404 Not Found</title>
</head><body>
<h1>Not Found</h1>
<p>The requested URL / was not found on this server.</p>
<hr>
<address>Apache Server at crt.sh Port 443</address>
</body></html>
* Connection #0 to host crt.sh left intact
```
---


Enumerates organization's domain names through SSL Subject Alternative Name entries via crt.sh 

© Blacksun Labs 2018

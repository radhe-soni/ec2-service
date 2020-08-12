# ec2-service

1. build using below command
<pre> go build -ldflags="-s -w" -o ec2Service.exe</pre>

2. provide config.yml dir while running the app
e.g.
<pre>ec2Service.exe c:/config-dir</pre>

3. The configuration filename has to be config.yml

4. Provide new-ip to be added in Ingress rules.

5. Provide rule description to update (rule-description is case-insensitive)
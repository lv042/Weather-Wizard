#!/bin/bash

# Tunnels the traffic from our nginx-docker container to the localhost.run server. Writes the output to a file.
ssh -o StrictHostKeyChecking=no -R 80:nginx:80 nokey@localhost.run -- --output json > /tmp/localhost-run &

# PHP creates a simple REST-server on port 8000 where we make the /tmp directory available.
php -S 0.0.0.0:8000 -t /tmp
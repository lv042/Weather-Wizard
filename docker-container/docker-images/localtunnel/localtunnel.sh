#!/bin/bash

# Tunnels the traffic from our nginx-docker container to the localtunnel server. Writes the output to a file.
lt --local-host nginx --port 80 --subdomain $1 > /tmp/localtunnel &

# PHP creates a simple REST-server on port 8000 where we make the /tmp directory available.
php -S 0.0.0.0:8000 -t /tmp


# Open the terminal of the docker 'iot-localtunnel'
# top > find id of lt --local-host ....
# kill process-id
# run command by hand: [line 4]
# go to url; https://abc3.loca.lt/index.php?action=config

git filter-branch --env-filter 'if [ $GIT_COMMIT = 54dc9ee8b2bd0980da868b3c8670ad676b883578 ]; then export GIT_AUTHOR_DATE="2023-03-09T00:13:00" GIT_COMMITTER_DATE="<new date>"; fi' HEAD~4..HEAD

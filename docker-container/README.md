# Individual Project Docker

## Used Stack

The stack we use in the Individual Project of IoT is as follows;

- Webserver: Nginx
- Database server: MariaDB
- Reverse proxy: localhost.run and/or localtunnel

## Login details

MySQL Database login details

|Key|Value|
|---|:--|
|Host|localhost|
|Port|3306|
|Username|root|
|Password|7YKyE8R2AhKzswfN|

## How to use

Update the configuration in .env by entering your own HvA username in the ```TUNNEL_SUBDOMAIN``` variable.

Run the Docker container by using the following command from this directory:

```docker-compose up```

Now you should see 6 different docker containers, combined as 'individual-project-docker':

- IoT-mariadb
- IoT-phpmyadmin
- IoT-php
- IoT-nginx
- IoT-localhost-run
- IoT-localtunnel

You can visit your website by going in your browser to: ```http://localhost/```.

To alter the files of your website you can look into the '```web```'-folder. You can open this folder in your favorite IDE, like Visual Studio Code. Every time you update files in the directory you only have to resfresh the page, you don't have to restart the docker containers.

## Resetting Docker container

If you only want to delete all the contents of your database you can delete the 'config/mariadb'-folder completely. After removall you have to restart the docker container. Be aware that you will lose everything inside your database!!!!

If you want to start over from the beginning, remove the old Docker container, delete this directory and start over by cloning this repository again.

When you delete the Docker Containers but keep the files, everything will be back when starting the containers again.
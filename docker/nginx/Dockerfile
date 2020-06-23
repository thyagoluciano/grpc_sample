FROM nginx
RUN mkdir /etc/nginx/cert && mkdir -p /usr/local/var/log/nginx
COPY cert/server-cert.pem cert/server-key.pem cert/ca-cert.pem /etc/nginx/cert/
COPY docker/nginx/nginx.conf /etc/nginx/nginx.conf
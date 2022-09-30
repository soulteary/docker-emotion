FROM  soulteary/emotion:ui-2022.09.30 AS WEBUI

FROM soulteary/emotion:base-2022.09.30
LABEL MAINTAINER=soulteary@gmail.com
RUN pip install gradio==3.4.0 && \
    pip cache purge
COPY docker/web/*.py ./
CMD ["python", "app.py"]

RUN apt update && apt install supervisor -y && \
    apt-get remove --purge -y && rm -rf /var/lib/apt/lists/*
SHELL ["/bin/bash", "-c"]
RUN echo $' \n\
[unix_http_server] \n\
file=/var/run/supervisor.sock \n\
chmod=0700 \n\
[inet_http_server] \n\
port=0.0.0.0:8080 \n\
[supervisord] \n\
nodaemon=true \n\
logfile=/var/log/supervisor/supervisord.log \n\
pidfile=/var/run/supervisord.pid \n\
childlogdir=/var/log/supervisor \n\
[rpcinterface:supervisor] \n\
supervisor.rpcinterface_factory = supervisor.rpcinterface:make_main_rpcinterface \n\
[supervisorctl] \n\
serverurl=unix:///var/run/supervisor.sock \n\
[program:server] \n\
directory=/app \n\
command=python app.py \n\
[program:webui] \n\
command=/app/web \n\
'> /etc/supervisor/supervisord.conf

CMD ["/usr/bin/supervisord","-c","/etc/supervisor/supervisord.conf"]
COPY --from=WEBUI /web  /app/web

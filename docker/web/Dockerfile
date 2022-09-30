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

COPY --from=WEBUI /web                                 /app/web

# # extra deps for python application
# COPY --from=Server /usr/lib/x86_64-linux-gnu/libGL.so.1.7.0               /usr/lib/x86_64-linux-gnu/
# COPY --from=Server /usr/lib/x86_64-linux-gnu/libGLX.so.0.0.0              /usr/lib/x86_64-linux-gnu/
# COPY --from=Server /usr/lib/x86_64-linux-gnu/libGLdispatch.so.0.0.0       /usr/lib/x86_64-linux-gnu/
# COPY --from=Server /usr/lib/x86_64-linux-gnu/libX11.so.6.3.0              /usr/lib/x86_64-linux-gnu/
# COPY --from=Server /usr/lib/x86_64-linux-gnu/libXext.so.6.4.0             /usr/lib/x86_64-linux-gnu/
# COPY --from=Server /usr/lib/x86_64-linux-gnu/libxcb.so.1.1.0              /usr/lib/x86_64-linux-gnu/
# COPY --from=Server /usr/lib/x86_64-linux-gnu/libXau.so.6.0.0              /usr/lib/x86_64-linux-gnu/
# COPY --from=Server /usr/lib/x86_64-linux-gnu/libXdmcp.so.6.0.0            /usr/lib/x86_64-linux-gnu/
# COPY --from=Server /usr/lib/x86_64-linux-gnu/libbsd.so.0.9.1              /usr/lib/x86_64-linux-gnu/
# COPY --from=Server /usr/lib/x86_64-linux-gnu/libgthread-2.0.so.0.5800.3   /usr/lib/x86_64-linux-gnu/
# COPY --from=Server /usr/lib/x86_64-linux-gnu/libglib-2.0.so.0.5800.3      /usr/lib/x86_64-linux-gnu/
# COPY --from=Server /lib/x86_64-linux-gnu/libdl-2.28.so                    /lib/x86_64-linux-gnu/
# RUN ln -s /usr/lib/x86_64-linux-gnu/libGL.so.1.7.0             /usr/lib/x86_64-linux-gnu/libGL.so.1           && \
#     ln -s /usr/lib/x86_64-linux-gnu/libGLX.so.0.0.0            /usr/lib/x86_64-linux-gnu/libGLX.so.0          && \
#     ln -s /usr/lib/x86_64-linux-gnu/libGLdispatch.so.0.0.0     /usr/lib/x86_64-linux-gnu/libGLdispatch.so.0   && \
#     ln -s /usr/lib/x86_64-linux-gnu/libX11.so.6.3.0            /usr/lib/x86_64-linux-gnu/libX11.so.6          && \
#     ln -s /usr/lib/x86_64-linux-gnu/libXext.so.6.4.0           /usr/lib/x86_64-linux-gnu/libXext.so.6         && \
#     ln -s /usr/lib/x86_64-linux-gnu/libxcb.so.1.1.0            /usr/lib/x86_64-linux-gnu/libxcb.so.1          && \
#     ln -s /usr/lib/x86_64-linux-gnu/libXau.so.6.0.0            /usr/lib/x86_64-linux-gnu/libXau.so.6          && \
#     ln -s /usr/lib/x86_64-linux-gnu/libXdmcp.so.6.0.0          /usr/lib/x86_64-linux-gnu/libXdmcp.so.6        && \
#     ln -s /usr/lib/x86_64-linux-gnu/libbsd.so.0.9.1            /usr/lib/x86_64-linux-gnu/libbsd.so.0          && \
#     ln -s /usr/lib/x86_64-linux-gnu/libgthread-2.0.so.0.5800.3 /usr/lib/x86_64-linux-gnu/libgthread-2.0.so.0  && \
#     ln -s /usr/lib/x86_64-linux-gnu/libglib-2.0.so.0.5800.3    /usr/lib/x86_64-linux-gnu/libglib-2.0.so.0     && \
#     ln -s /lib/x86_64-linux-gnu/libdl-2.28.so                  /lib/x86_64-linux-gnu/libdl.so.1
FROM soulteary/emotion:base-2022.09.30
LABEL MAINTAINER=soulteary@gmail.com
RUN pip install gradio==3.4.0 && \
    pip cache purge
COPY docker/web/*.py ./
CMD ["python", "app.py"]
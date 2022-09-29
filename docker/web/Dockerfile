FROM base
LABEL MAINTAINER=soulteary@gmail.com
RUN pip install gradio==3.4.0
COPY docker/web/app.py ./
CMD ["python", "app.py"]
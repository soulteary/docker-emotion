FROM python:3.9-slim-buster
LABEL MAINTAINER=soulteary@gmail.com
RUN sed -i -E "s/\w+.debian.org/mirrors.tuna.tsinghua.edu.cn/g" /etc/apt/sources.list
RUN pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple && \
    pip3 install --upgrade pip && \
    pip install transformers==4.22.2 && \
    pip install torch --extra-index-url https://download.pytorch.org/whl/cpu && \
    pip cache purge
RUN python -c "from transformers import pipeline; pipeline('text-classification',model='bhadresh-savani/bert-base-uncased-emotion', top_k=1)" && \
    python -c "import transformers; transformers.utils.move_cache()"
WORKDIR /app/
COPY ./docker/bert-base-uncased-emotion/app.py ./
CMD ["python", "app.py"]
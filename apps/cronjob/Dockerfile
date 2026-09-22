FROM python:3-slim

WORKDIR /app

RUN pip install requests

COPY main.py ./

ENTRYPOINT [ "python", "./main.py" ]
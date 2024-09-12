FROM python:3.9.19

WORKDIR /telemetry-analysis

RUN pip install --upgrade pip

# RUN python -m pip install pip-tools

COPY ./requirements.txt /telemetry-analysis/requirements.txt
RUN pip install --no-cache-dir --upgrade -r /telemetry-analysis/requirements.txt

COPY . /telemetry-analysis/

CMD ["python", "main.py"]

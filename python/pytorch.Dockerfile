FROM pytorch/pytorch:latest

COPY . .
RUN pip install --upgrade pip && pip install -e ./kfserving
RUN pip install -e ./sklearnserver
ENTRYPOINT ["python", "-m", "pytorchserver"]

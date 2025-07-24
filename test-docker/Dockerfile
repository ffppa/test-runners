FROM ubuntu:latest

LABEL version="1.3"
LABEL description="Dockerfile di test per pipeline CI/CD"

RUN apt-get update && apt-get install -y \
  curl \
  wget \
  vim \
  && apt-get clean

WORKDIR /app

COPY test_script.sh /app/test_script.sh

RUN chmod +x /app/test_script.sh

ENTRYPOINT ["/app/test_script.sh"]
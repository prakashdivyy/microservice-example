# set base image (host OS)
FROM python:3.8-alpine

# set the working directory in the container
WORKDIR /code

# copy the dependencies file to the working directory
COPY requirements.txt .

# install dependencies
RUN pip install -r requirements.txt

# copy receive.py to the working directory
COPY receive.py .

# command to run on container start
CMD [ "python", "./receive.py" ]
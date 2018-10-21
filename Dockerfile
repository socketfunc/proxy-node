FROM arm32v7/node:10.12

MAINTAINER yu-icchi

RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get dist-upgrade -y
RUN apt-get install -y libstdc++6

RUN mkdir app
WORKDIR app

COPY package.json .
COPY package-lock.json .

RUN mkdir lib
RUN mkdir proto
RUN mkdir bin

COPY lib lib
COPY proto proto
COPY bin bin

RUN npm install --production
RUN npm rebuild --build-from-source grpc

EXPOSE 5000

CMD ["node", "./bin/proxy-node", "--port", "5000"]

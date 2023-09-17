FROM node:alpine

WORKDIR /gateway

COPY ./package.json ./
RUN yarn install

COPY . .

CMD ["yarn", "dev"]
FROM golang:latest

ARG GITHUB_ACCESS_TOKEN
ARG MANAGEMENT_SERVER_ADDR
ARG API_VERSION
ARG OWNER_ROLE
ARG REFRESH_SECRET_KEY
ARG ACCESS_SECRET_KEY
ARG EXP_ACCESS
ARG EXP_REFRESH
ARG X_HEADER

COPY . .

RUN mkdir ~/.ssh
RUN touch ~/.ssh/known_hosts
RUN ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts

# allow private repo pull
RUN git config --global url."https://$GITHUB_ACCESS_TOKEN:x-oauth-basic@github.com/".insteadOf "https://github.com/"

# go mod downlod now has access to private modules.
RUN go mod download

RUN go build

CMD [ "./skultem-gateway" ]
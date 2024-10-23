FROM golang:latest

ARG GITHUB_ACCESS_TOKEN

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
FROM golang:latest

#ARG GITHUB_TOKEN
#RUN git config --global url."https://aicactus:${GITHUB_TOKEN}@github.com".insteadOf "https://github.com"

#Set the current working directory insite the container
WORKDIR /app

COPY go.mod go.sum ./

#Download all dependencies
RUN go mod download

#Copy the source from the current directory to the Working Directory insite the Container
COPY . .

RUN go build -o cob-as-prt

EXPOSE 8080

ENTRYPOINT ["./cob-as-prt"]

#Cammand to run the executable
CMD ["server"]
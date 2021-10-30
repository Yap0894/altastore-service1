# FROM sonarsource/sonar-scanner-cli:latest AS sonarqube_scan
# WORKDIR /app
# COPY . .
# RUN ls -list    
# # sonar.projectName property used for providing human-friendly project name in addition 
# # for projectKey
# RUN sonar-scanner \
#     -Dsonar.host.url="http://localhost:9090" \
#     -Dsonar.projectKey="AltaStore" \
#     -Dsonar.sources="." \
#     -Dsonar.go.coverage.reportPaths="coverage.out" \
#     -Dsonar.login="ff9f3a19daf80a937e25559f51f3d6049b0525a6" \ 
#     -Dsonar.exclusions="business/**/service_test.go, business/errors.go, api/**, app/**, config/**, modules/**, util/**"\
#     -Dsonar.test.exclusions="business/**/service_test.go"\  

# stage I - khusus build dengan envinroment yang sama
FROM golang:1.16-alpine AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download -x

RUN go build -o main 
# EXPOSE 8080
# CMD ["/app/main"]

# stage 2
WORKDIR /app/webservice

COPY --from=builder /app/.env ./.env
COPY --from=builder /app/main .

EXPOSE 8000
CMD ["./main"]
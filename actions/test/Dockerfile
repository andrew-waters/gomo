FROM golang:1.10

LABEL "name"="go test"
LABEL "version"="1.0.0"

LABEL "com.github.actions.icon"="code"
LABEL "com.github.actions.color"="green-dark"
LABEL "com.github.actions.name"="gotest"
LABEL "com.github.actions.description"="This is an action to run go test."

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

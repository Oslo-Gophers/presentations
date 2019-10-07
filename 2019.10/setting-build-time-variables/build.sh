
# Ask the user for the version we are building
# (this script can't figure it out on its own)
read -p 'Version: ' VERSION


# Set the variables we can figure out automatically 
# (here to clean up the `go build` command a little)
TIME=$(date -u +%d.%m.%Y_%H:%M:%S)
GIT_COMMIT=$(cd serious-server/ &&  git log -1 --pretty=format:"%H")  # Of course this assumes evrything is commited
HOSTNAME=$(hostname)
USERNAME=$(whoami)

# Compile
echo "Running go build on version ${VERSION}.."

go build -o server -ldflags "               \
    -X main.version=${VERSION}              \
    -X main.buildTime=${TIME}-UTC           \
    -X main.buildGitCommit=${GIT_COMMIT}    \
    -X main.buildHost=${HOSTNAME}           \
    -X main.buildUser=${USERNAME}"          \
    ./serious-server


echo "Done!"
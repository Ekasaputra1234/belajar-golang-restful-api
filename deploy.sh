TAG=$(echo "$CI_COMMIT_TAG" | sed "s/+/-/" | sed "s/v//")
git reset --hard && \
git fetch $CI_REPOSITORY_URL +refs/heads/*:refs/remotes/origin/* && \
git checkout $CI_COMMIT_TAG && \
docker build \
  -t "${CI_SERVER_HOST}/voltunes/api-master-project:$TAG" \
  --build-arg CI_SERVER_HOST="${CI_SERVER_HOST}" \
  --build-arg CI_JOB_TOKEN="${CI_JOB_TOKEN}" \
  --build-arg CI_COMMIT_TAG="${CI_COMMIT_TAG}" \
  --build-arg TAG_DELETE_TOKEN="${TAG_DELETE_TOKEN}" \
  --build-arg TAG="${TAG}" \
  --no-cache \
  . && \
docker-compose down && \
TAG=$TAG docker-compose up -d

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
  echo "*** Calling ${BASH_SOURCE[0]} directly has no effect. It should be sourced."
  exit -1
fi

# standard checks
set -ex
set -o errexit
set -o nounset
set -o pipefail

function usage() {
  echo "$0 \
    -h,-hub <docker image repository> \
    -t,-tag <comma separated list of docker image TAGS> \
    -b,-build-only <docker image repository>"
  exit 1
}

function docker_push() {
  local im="${1}"
  if [[ "${im}" =~ ^gcr\.io ]]; then
    gcloud docker -- push ${im}
  else
    docker push ${im}
  fi
}

# Tag and push
function tag_and_push() {
  local IMAGES="$1"

  for IMAGE in ${IMAGES[@]}; do
    for TAG in ${TAGS[@]}; do
      for HUB in ${HUBS[@]}; do
        docker tag "${IMAGE}" "${HUB}/${IMAGE}:${TAG}"
        if [ "${build_only}" != "true" ]; then
          docker_push "${HUB}/${IMAGE}:${TAG}"
        fi
      done
    done
  done
}

HUBS="gcr.io/istio-testing"
local_tag=$(whoami)_$(date +%y%m%d_%H%M%S)
TAGS="${local_tag}"
BUILD_ONLY="false"

while [[ $# -gt 0 ]]; do
    case "$1" in
        -tag) TAGS="$2"; shift ;;
        -t) TAGS="$2"; shift ;;
        -hub) HUBS="$2"; shift ;;
        -h) HUBS="$2"; shift ;;
        -build-only) BUILD_ONLY="true";;
        -b) BUILD_ONLY="true";;
        *) ;;
    esac
    shift
done


IFS=',' read -ra TAGS <<< "${TAGS}"
IFS=',' read -ra HUBS <<< "${HUBS}"



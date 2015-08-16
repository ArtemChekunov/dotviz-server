#!/usr/bin/env bash

echo "${DOTVIZ_PORT}" "${DOTVIZ_HOST}"


if [ -n "${DOTVIZ_PORT}" ] &&  [ -n "${DOTVIZ_HOST}" ] ; then
  /go/bin/dotviz-server -host ${DOTVIZ_HOST} -port ${DOTVIZ_PORT}

elif [ -n "${DOTVIZ_HOST}" ] ; then
  /go/bin/dotviz-server -host ${DOTVIZ_HOST}

elif [ -n "${DOTVIZ_PORT}" ] ; then
  /go/bin/dotviz-server -port ${DOTVIZ_PORT}

else
  /go/bin/dotviz-server
fi

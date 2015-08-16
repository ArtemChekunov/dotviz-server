#!/usr/bin/env bash

echo "${DOTVIZ_PORT}" "${DOTVIZ_HOST}"


if [ -n "${DOTVIZ_PORT}" ] &&  [ -n "${DOTVIZ_HOST}" ] ; then
  dotviz-server -host ${DOTVIZ_HOST} -port ${DOTVIZ_PORT}

elif [ -n "${DOTVIZ_HOST}" ] ; then
  dotviz-server -host ${DOTVIZ_HOST}

elif [ -n "${DOTVIZ_PORT}" ] ; then
  dotviz-server -port ${DOTVIZ_PORT}

else
  dotviz-server
fi

#!/bin/bash -i
#set -x
COMP_LINE="$1"
COMP_POINT=${#1}
COMP_WORDS=($COMP_LINE'')
COMP_CWORD=$((${#COMP_WORDS[@]}-1))
COMP_WORDBREAKS=" "

f=$(complete -p "${COMP_WORDS[0]}" | awk '{print $3}')
function compopt { return; }

$f ${COMP_WORDS[0]}
( IFS=$'\n'; echo "${COMPREPLY[*]}" )

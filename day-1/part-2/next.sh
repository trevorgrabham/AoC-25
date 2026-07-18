#!/usr/bin/env bash 

if [[ $PWD =~ /day-([0-9]+)/part-2$ ]]; then 
  day=${BASH_REMATCH[1]}
  ((day++))

  cd ../..
  . scaffold.sh "$day"
else 
  echo "invalid working dir $PWD"
fi

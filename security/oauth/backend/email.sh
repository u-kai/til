#!/bin/zsh
curl -L \
  -X GET \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer OAUTH_TOKEN"\
  -H "X-GitHub-Api-Version: 2022-11-28" \
   https://api.github.com/user/repos
##  https://api.github.com/user/emails \

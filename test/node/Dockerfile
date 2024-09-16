#--------------------------------------
# Image: base
#--------------------------------------
FROM ghcr.io/jamiemagee/dependabot-tools AS base

RUN touch /.dummy

#--------------------------------------
# git
#--------------------------------------
FROM base AS testa

RUN dependabot-tools install node 20.17.0

RUN node --version

#--------------------------------------
# final
#--------------------------------------
FROM base

COPY --from=testa /.dummy /.dummy


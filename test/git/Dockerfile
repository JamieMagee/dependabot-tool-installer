#--------------------------------------
# Image: base
#--------------------------------------
FROM ghcr.io/jamiemagee/dependabot-tools AS base

RUN touch /.dummy

#--------------------------------------
# git
#--------------------------------------
FROM base AS testa

RUN dependabot-tools install git

RUN git --version

#--------------------------------------
# final
#--------------------------------------
FROM base

COPY --from=testa /.dummy /.dummy


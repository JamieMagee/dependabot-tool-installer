#--------------------------------------
# Image: base
#--------------------------------------
FROM dependabot-tools AS base

RUN touch /.dummy

#--------------------------------------
# .NET Core 3.1
#--------------------------------------
FROM base AS testa

RUN dependabot-tools install dotnet 3.1.101

RUN dotnet --version

#--------------------------------------
# .NET 6.0
#--------------------------------------
FROM base AS testb

RUN dependabot-tools install dotnet 6.0.424

RUN dotnet --version

#--------------------------------------
# .NET 9.0
#--------------------------------------
FROM base AS testc

RUN dependabot-tools install dotnet 9.0.100-preview.6.24328.19

RUN dotnet --version

#--------------------------------------
# final
#--------------------------------------
FROM base

COPY --from=testa /.dummy /.dummy
COPY --from=testb /.dummy /.dummy
COPY --from=testc /.dummy /.dummy

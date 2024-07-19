# Dependabot Tools Installer (DTI)

Dependabot Tools Installer (DTI) is a CLI to install compilers, SDKs, and other tools required by Dependabot to run against a repository.
This project is not intended to be used directly by Dependabot users.

## Problem Statement

Currently, Dependabot uses a fixed version of all the tools it needs to run.
If your repository requires a different version of a tool, Dependabot will likely fail with [`ToolVersionNotSupported`][1], [`SubprocessFailed `][2], [`HelperSubprocessFailed`][3], or similar errors.

Some ecosystems are more resilient to this problem than others.
However, some ecosystems, like Python, are more prone to this problem.
For example, Dependabot installs 5 different versions of Python[^1] to accommodate different repositories.
This approach is not scalable or sustainable in the long run, and it doesn't solve the problem for other ecosystems.

## Solution

Instead of packaging all the tools with Dependabot, DTI will install the required tools on the fly.
This way, Dependabot can run with the tools required by the repository, and the repository owner can control the versions of the tools.

Each tool needs to have its own installer implementation to install it.
For example, to install .NET, DTI will:

1. Use `apt-get` to install .NET's dependencies.
2. Download and install the appropriate .NET SDK for the current architecture and Linux distribution.
3. Configure the environment variables required by .NET.

Roughly, an installer implementation is expected to implement the following interface:

```go
type Installer interface {
  InstallPreRequisites(distro Distro, arch Arch) error
  Install(distro Distro, arch Arch) error
  PostInstall() error
}
```

## Other Approaches Considered

### GitHub Actions

One approach to solve this problem is to use GitHub Actions to install the tools.
GitHub Actions such as [`actions/setup-dotnet`][4], [`actions/setup-python`][5], and [`swift-actions/setup-swift`][6] are already available.

However, the main problem with this approach is that GitHub Actions are designed to run within a GitHub Actions workflow. For example, environment variables are set by writing to standard output.
We would need to re-implement the logic to set environment variables in the shell.

Additionally, GitHub Actions are tailored to run in the specific container images from [`actions/runner-images`][7].

Finally, it would require us to ship a Node.js environment to run the GitHub Actions.
And take on the responsibility of maintaining the GitHub Actions and their dependencies.

### Distribution Packages

Another approach is to use the distribution's package manager to install the tools.
Unfortunately, Ubuntu only provides the latest version of the tools, and it's not possible to install a specific version.

While some projects like .NET provide a package repository[^3], it's not a common practice.

### Containerbase

[Containerbase][8] is a similar project that installs tools required by Renovate.
It uses shell scripts to install the tools, which is a similar approach to DTI.
However, Containerbase is not designed to be used by Dependabot.
If Dependabot were to use Containerbase, we would either need to fork the project or add a lot of complexity to support Dependabot.

[1]: https://github.sentry.io/issues/?project=1451818&query=error.type%3ADependabot%3A%3AToolVersionNotSupported
[2]: https://github.sentry.io/issues/?project=1451818&query=error.type%3ADependabot%3A%3AUpdater%3A%3ASubprocessFailed
[3]: https://github.sentry.io/issues/?project=1451818&query=error.type%3ADependabot%3A%3ASharedHelpers%3A%3AHelperSubprocessFailed
[4]: https://github.com/actions/setup-dotnet
[5]: https://github.com/actions/setup-python
[6]: https://github.com/actions/setup-python
[7]: https://github.com/actions/runner-images
[8]: https://github.com/containerbase/base

[^1]: https://github.com/dependabot/dependabot-core/blob/main/python/Dockerfile
[^3]: https://packages.microsoft.com/ubuntu/22.04/prod/pool/main/d/dotnet-sdk-8.0/
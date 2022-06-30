# mono-repo

A test mono-repo to practice building out some tools that would help for mono repos in future projects.

## CI

I am using the new CircleCi dynamic config feature to build up one config from multiple nested configs. The config is included within the main CI job if there are any changes within the service, this allows us to only build and test services that have been modified within the pull request.

## Versioning

Each service has a `VERSION` file stored in its directory. If the service has been modified we check to ensure the new version number is a valid bump.

## TODO
- Add orb for version checking
- Build image and deploy to wherever we want to store our images
- Run tests within CI
- Integration tests of multiple services
- Run linting on services
- Add services with various languages to ensure they all build, test & deploy successfully, not just Go.
- Ensure everything still runs if the commit is made directly to master
- Put the code that checks if a service has changed into a CLI or orb

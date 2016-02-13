# go-github

Golang helper to query github repositories.

## Motivation

I needed to get the latest release URL of a github project to import it into my docker images.

## Usage

Simply provide the owner/organisation and the repository name and off you go.

```
➜  ./go-github rLatestUrl --ghorg ashwanthkumar --ghrepo gocd-slack-build-notifier
https://github.com/ashwanthkumar/gocd-slack-build-notifier/releases/download/v1.1.2/gocd-slack-notifier-1.1.2.jar
https://github.com/ashwanthkumar/gocd-slack-build-notifier/releases/download/v1.1.1/gocd-slack-notifier-1.1.1.jar
https://github.com/ashwanthkumar/gocd-slack-build-notifier/releases/download/1.0/gocd-slack-notifier-1.0.jar
➜ ./go-github rLatestUrl --ghorg ashwanthkumar --ghrepo gocd-slack-build-notifier --limit 1
https://github.com/ashwanthkumar/gocd-slack-build-notifier/releases/download/v1.1.2/gocd-slack-notifier-1.1.2.jar
```

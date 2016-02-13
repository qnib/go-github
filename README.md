# go-github

Golang helper to query github repositories.

## Motivation

I needed to get the latest release URL of a github project to import it into my docker images.

## Usage

Simply provide the owner/organisation and the repository name and off you go.

```
➜ ./go-github_0.2.2_Darwin rLatestUrl --ghrepo go-github --regex ".*inux.*" --limit 0
https://github.com/qnib/go-github/releases/download/0.2.1/go-github_0.2.1_Linux
https://github.com/qnib/go-github/releases/download/0.1/go-github-0.1_linux_amd64
➜  ./go-github_0.2.2_Darwin rLatestUrl --ghrepo go-github --regex ".*inux.*" --limit 1
https://github.com/qnib/go-github/releases/download/0.2.1/go-github_0.2.1_Linux
➜ 
```

```
➜  ./go-github rLatestUrl --ghorg ashwanthkumar --ghrepo gocd-slack-build-notifier
https://github.com/ashwanthkumar/gocd-slack-build-notifier/releases/download/v1.1.2/gocd-slack-notifier-1.1.2.jar
https://github.com/ashwanthkumar/gocd-slack-build-notifier/releases/download/v1.1.1/gocd-slack-notifier-1.1.1.jar
https://github.com/ashwanthkumar/gocd-slack-build-notifier/releases/download/1.0/gocd-slack-notifier-1.0.jar
➜ ./go-github rLatestUrl --ghorg ashwanthkumar --ghrepo gocd-slack-build-notifier --limit 1
https://github.com/ashwanthkumar/gocd-slack-build-notifier/releases/download/v1.1.2/gocd-slack-notifier-1.1.2.jar
```

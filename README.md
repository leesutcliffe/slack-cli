# slack-cli

Basic CLI tool that currently only sets Slack profile status and presence (away or active)
Writen maily because clicking 'Set yourself as away' or 'out to lunch' is too much effort. 

# Slack API Permisions

To access the Slack API endpoints, you'll need to create access to the following Slack API scopes

- users.profile:write
- users:write

# config file

A config file needs to be saved in ~/.slack/config
This will contain the token required to access the Slack API for each workspace and and preset status profiles
A default workspace can be specified so that future versions of the code can select a particular workspace and potentially loop through all workspaces

``` yaml
workspaces:
  - name: work
    token: xoxp-123-456-789-10
  - name: personal
    token: xoxp-123-456-789-10
status:
  - name: lunch
    message: "Out to lunch"
    emoji: ":bagel:"
  - name: holiday
    message: "Holiday"
    emoji: ":beach_with_umbrella:"
default: work
```

# install 

`go build'

# commands

```
# reset status to blank and 'Active'
slack set status

# set status to 'away'
slack set status -a

# set status to 'lunch' profile and away
slack set status --profile=lunch --away
slack set status -p=lunch -a
```


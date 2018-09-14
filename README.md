# gist-diff

Status: Blocked - at this current time there is no API available to create gists.

Takes the inputs of a git diff command and publishes it as a gist in Github.

`git diff | gist_diff`

# Creating a personal access token

https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/

Note* you only need to give the PAT `gist` permissions.

Once you have a token as per the instructions above, add it to your `config.json`.

```json
{
  "github_username": "willis7",
  "github_password": "secretfagbsajhvdakhdsecret"
}

```
# PRs Management Tool

This tool is designed to help you manage your Azure DevOps pull requests (PRs) efficiently. It allows you to complete or abandon multiple PRs by providing their links through standard input (STDIN). This can be particularly useful when you have a large number of PRs to handle and want to automate the process.

## Features

- Complete multiple PRs with a single command
- Abandon multiple PRs with a single command
- Automatically squash and delete the source branch when completing PRs

## Installation

To install this tool, you need to have Go installed on your machine. Follow these steps:

1. Clone the repository:

    ```sh
    git clone git@github.com/cmdctl/prs.git
    cd prs
    ```

2. Build the executable:

    ```sh
    go build -o prs
    ```

3. Move the executable to a directory in your PATH:

    ```sh
    mv prs /usr/local/bin/
    ```

## Usage

The tool provides two main commands: `complete` and `abandon`. Each command requires a list of PR links provided through STDIN.

### Complete PRs

To complete PRs, use the `complete` command. This command will:

- Mark the PR as completed
- Squash the commits [OPTIONAL] (use `--squash` flag) default is `false`
- Delete source branch [OPTIONAL] (use `--delete-source-branch` flag) default is `false`

#### Example

1. Create a file `prs.txt` with the list of PR links, each on a new line:

    ```
    https://dev.azure.com/yourorganization/yourproject/_git/yourrepo/pullrequest/123
    https://dev.azure.com/yourorganization/yourproject/_git/yourrepo/pullrequest/456
    ```

2. Run the command:

    ```sh
    cat prs.txt | prs complete
    ```

    or

    ```sh
    prs complete < prs.txt
    ```

### Abandon PRs

To abandon PRs, use the `abandon` command. This command will mark the PR as abandoned.

#### Example

1. Create a file `prs.txt` with the list of PR links, each on a new line:

    ```
    https://dev.azure.com/yourorganization/yourproject/_git/yourrepo/pullrequest/123
    https://dev.azure.com/yourorganization/yourproject/_git/yourrepo/pullrequest/456
    ```

2. Run the command:

    ```sh
    cat prs.txt | prs abandon
    ```

    or

    ```sh
    prs abandon < prs.txt
    ```

### Help

To display the help message, use the `--help` or `-h` flag:

```sh
prs --help

```
Outputs:
```
Usage:
  prs [command] [options]

Commands:
  complete    Complete the specified pull requests
  abandon     Abandon the specified pull requests

Options:
  --help, -h  Show this help message and exit

Complete Command Options:
  --delete-source-branch  Delete the source branch after completing the pull request
  --squash                Squash the commits when completing the pull request

Abandon Command Options:
  --delete-source-branch  Delete the source branch after abandoning the pull request

```

## Input Structure

The input should be a list of PR links, each on a new line. The tool will read these links from STDIN and process them accordingly. Make sure the links are in the following format:

```
https://dev.azure.com/yourorganization/yourproject/_git/yourrepo/pullrequest/123
https://dev.azure.com/yourorganization/yourproject/_git/yourrepo/pullrequest/456
```

### Using Comments in Input

You can include comments in your input file by prefixing the comment lines with `--`. These lines will be ignored by the tool.

#### Example

Create a file `prs.txt` with PR links and comments:

```
-- This is a comment and will be ignored
https://dev.azure.com/yourorganization/yourproject/_git/yourrepo/pullrequest/123
-- Another comment
https://dev.azure.com/yourorganization/yourproject/_git/yourrepo/pullrequest/456
```

## Error Handling

If the tool encounters an error while processing a PR, it will print an error message and continue to the next PR. Make sure to check the output for any errors that may occur during the execution.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.




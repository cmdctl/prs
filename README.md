# PRS cli tool
A handy tool to complete or abandon multiple PRs in azure devops at once.
Uses azure cli under the hood.

Collect your PR links in a file and pass that file to the tool

INSTALLATION
```bash
git clone git@github.com/cmdctl/prs.git
cd prs
go install
```

USAGE
```bash
$ prs --help
```

OUTPUT:
```bash
---------------------------
USAGE: prs complete | abandon
Make sure to provide a list of PRs to STDIN - each link should be on a new line
EXAMPPLE 1: cat prs.txt | prs complete
EXAMPPLE 2: prs complete < prs.txt
---------------------------
```

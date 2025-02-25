Feature: display all executed Git commands

  Background:
    Given the current branch is a feature branch "old"
    And the commits
      | BRANCH | LOCATION      | MESSAGE    |
      | old    | local, origin | old commit |

  # TODO: eliminate redundant "git rev-parse --short HEAD"
  Scenario: result
    When I run "git-town prepend parent --verbose"
    Then it runs the commands
      | BRANCH | TYPE     | COMMAND                                       |
      |        | backend  | git version                                   |
      |        | backend  | git config -lz --global                       |
      |        | backend  | git config -lz --local                        |
      |        | backend  | git rev-parse --show-toplevel                 |
      |        | backend  | git stash list                                |
      |        | backend  | git branch -vva                               |
      |        | backend  | git remote                                    |
      | old    | frontend | git fetch --prune --tags                      |
      |        | backend  | git branch -vva                               |
      |        | backend  | git rev-parse --verify --abbrev-ref @{-1}     |
      |        | backend  | git status --long --ignore-submodules         |
      | old    | frontend | git checkout main                             |
      | main   | frontend | git rebase origin/main                        |
      |        | backend  | git rev-list --left-right main...origin/main  |
      | main   | frontend | git checkout old                              |
      | old    | frontend | git merge --no-edit origin/old                |
      |        | frontend | git merge --no-edit main                      |
      |        | backend  | git rev-list --left-right old...origin/old    |
      |        | backend  | git show-ref --verify --quiet refs/heads/main |
      | old    | frontend | git branch parent main                        |
      |        | backend  | git show-ref --verify --quiet refs/heads/main |
      |        | backend  | git config git-town-branch.parent.parent main |
      |        | backend  | git show-ref --verify --quiet refs/heads/old  |
      |        | backend  | git config git-town-branch.old.parent parent  |
      | old    | frontend | git checkout parent                           |
      |        | backend  | git show-ref --quiet refs/heads/old           |
      |        | backend  | git rev-parse --verify --abbrev-ref @{-1}     |
      |        | backend  | git config -lz --global                       |
      |        | backend  | git config -lz --local                        |
      |        | backend  | git branch -vva                               |
      |        | backend  | git stash list                                |
    And it prints:
      """
      Ran 31 shell commands.
      """
    And the current branch is now "parent"

  Scenario: undo
    Given I ran "git-town prepend parent"
    When I run "git-town undo --verbose"
    Then it runs the commands
      | BRANCH | TYPE     | COMMAND                                          |
      |        | backend  | git version                                      |
      |        | backend  | git config -lz --global                          |
      |        | backend  | git config -lz --local                           |
      |        | backend  | git rev-parse --show-toplevel                    |
      |        | backend  | git stash list                                   |
      |        | backend  | git branch -vva                                  |
      |        | backend  | git rev-parse --verify --abbrev-ref @{-1}        |
      |        | backend  | git status --long --ignore-submodules            |
      |        | backend  | git remote get-url origin                        |
      |        | backend  | git config --unset git-town-branch.parent.parent |
      |        | backend  | git config git-town-branch.old.parent main       |
      | parent | frontend | git checkout old                                 |
      | old    | frontend | git branch -D parent                             |
      |        | backend  | git show-ref --quiet refs/heads/old              |
      |        | backend  | git show-ref --quiet refs/heads/parent           |
      |        | backend  | git rev-parse --verify --abbrev-ref @{-1}        |
      |        | backend  | git checkout old                                 |
      |        | backend  | git checkout old                                 |
      |        | backend  | git config -lz --global                          |
      |        | backend  | git config -lz --local                           |
      |        | backend  | git branch -vva                                  |
      |        | backend  | git stash list                                   |
    And it prints:
      """
      Ran 22 shell commands.
      """
    And the current branch is now "old"

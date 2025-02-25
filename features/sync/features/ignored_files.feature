Feature: ignore files

  Background:
    Given the current branch is a feature branch "feature"
    And the commits
      | BRANCH  | LOCATION | MESSAGE   | FILE NAME  | FILE CONTENT |
      | feature | local    | my commit | .gitignore | ignored      |
    And an uncommitted file with name "test/ignored/important" and content "changed ignored file"
    When I run "git-town sync"

  Scenario: result
    Then file "test/ignored/important" still has content "changed ignored file"

  Scenario: undo
    When I run "git-town undo"
    Then it runs the commands
      | BRANCH  | COMMAND                                                               |
      | feature | git push --force-with-lease origin {{ sha 'Initial commit' }}:feature |
    And the current branch is still "feature"
    And now the initial commits exist
    And the initial branches and hierarchy exist

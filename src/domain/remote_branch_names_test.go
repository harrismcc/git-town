package domain_test

import (
	"testing"

	"github.com/git-town/git-town/v10/src/domain"
	"github.com/shoenig/test/must"
)

func TestRemoteBranchNames(t *testing.T) {
	t.Parallel()

	t.Run("Sort", func(t *testing.T) {
		t.Parallel()
		have := domain.RemoteBranchNames{
			domain.NewRemoteBranchName("origin/branch-3"),
			domain.NewRemoteBranchName("origin/branch-2"),
			domain.NewRemoteBranchName("origin/branch-1"),
		}
		have.Sort()
		want := domain.RemoteBranchNames{
			domain.NewRemoteBranchName("origin/branch-1"),
			domain.NewRemoteBranchName("origin/branch-2"),
			domain.NewRemoteBranchName("origin/branch-3"),
		}
		must.Eq(t, want, have)
	})

	t.Run("Strings", func(t *testing.T) {
		t.Parallel()
		give := domain.RemoteBranchNames{
			domain.NewRemoteBranchName("origin/branch-1"),
			domain.NewRemoteBranchName("origin/branch-2"),
			domain.NewRemoteBranchName("origin/branch-3"),
		}
		have := give.Strings()
		want := []string{
			"origin/branch-1",
			"origin/branch-2",
			"origin/branch-3",
		}
		must.Eq(t, want, have)
	})
}

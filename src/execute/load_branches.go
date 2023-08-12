package execute

import (
	"github.com/git-town/git-town/v9/src/git"
	"github.com/git-town/git-town/v9/src/validate"
)

// LoadBranches loads the typically used information about Git branches using a single Git command.
func LoadBranches(pr *git.ProdRunner, args LoadBranchesArgs) (git.Branches, error) {
	allBranches, initialBranch, err := pr.Backend.BranchesSyncStatus()
	if err != nil {
		return git.EmptyBranches(), err
	}
	branchDurations := pr.Config.BranchDurations()
	if args.ValidateIsConfigured {
		branchDurations, err = validate.IsConfigured(&pr.Backend, allBranches, branchDurations)
	}
	return git.Branches{
		All:       allBranches,
		Durations: branchDurations,
		Initial:   initialBranch,
	}, err
}

type LoadBranchesArgs struct {
	ValidateIsConfigured bool
}
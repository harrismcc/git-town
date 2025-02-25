package undo

import "github.com/git-town/git-town/v10/src/domain"

// BranchSpans describes how a Git Town command has modified the branches in a Git repository.
type BranchSpans []BranchSpan

func NewBranchSpans(beforeSnapshot, afterSnapshot domain.BranchesSnapshot) BranchSpans {
	result := BranchSpans{}
	for _, before := range beforeSnapshot.Branches {
		after := afterSnapshot.Branches.FindMatchingRecord(before)
		result = append(result, BranchSpan{
			Before: before,
			After:  after,
		})
	}
	for _, after := range afterSnapshot.Branches {
		if beforeSnapshot.Branches.FindMatchingRecord(after).IsEmpty() {
			result = append(result, BranchSpan{
				Before: domain.EmptyBranchInfo(),
				After:  after,
			})
		}
	}
	return result
}

// Changes describes the specific changes made in this BranchSpans.
func (self BranchSpans) Changes() BranchChanges {
	result := EmptyBranchChanges()
	for _, branchSpan := range self {
		if branchSpan.NoChanges() {
			continue
		}
		if branchSpan.IsOmniRemove() {
			result.OmniRemoved[branchSpan.Before.LocalName] = branchSpan.Before.LocalSHA
			continue
		}
		if branchSpan.IsOmniChange() {
			result.OmniChanged[branchSpan.Before.LocalName] = domain.Change[domain.SHA]{
				Before: branchSpan.Before.LocalSHA,
				After:  branchSpan.After.LocalSHA,
			}
			continue
		}
		if branchSpan.IsInconsistentChange() {
			result.InconsistentlyChanged = append(result.InconsistentlyChanged, domain.InconsistentChange{
				Before: branchSpan.Before,
				After:  branchSpan.After,
			})
			continue
		}
		switch {
		case branchSpan.LocalAdded():
			result.LocalAdded = append(result.LocalAdded, branchSpan.After.LocalName)
		case branchSpan.LocalRemoved():
			result.LocalRemoved[branchSpan.Before.LocalName] = branchSpan.Before.LocalSHA
		case branchSpan.LocalChanged():
			result.LocalChanged[branchSpan.Before.LocalName] = domain.Change[domain.SHA]{
				Before: branchSpan.Before.LocalSHA,
				After:  branchSpan.After.LocalSHA,
			}
		}
		switch {
		case branchSpan.RemoteAdded():
			result.RemoteAdded = append(result.RemoteAdded, branchSpan.After.RemoteName)
		case branchSpan.RemoteRemoved():
			result.RemoteRemoved[branchSpan.Before.RemoteName] = branchSpan.Before.RemoteSHA
		case branchSpan.RemoteChanged():
			result.RemoteChanged[branchSpan.Before.RemoteName] = domain.Change[domain.SHA]{
				Before: branchSpan.Before.RemoteSHA,
				After:  branchSpan.After.RemoteSHA,
			}
		}
	}
	return result
}

package main

func findChangedFiles(previousCommitMap, currentCommitMap map[string]string) []Change {

	changedFiles := []Change{}

	for fileInPreviousCommit, fileHashInPreviousCommit := range previousCommitMap {
		if currentHashForFile, exists := currentCommitMap[fileInPreviousCommit]; exists {
			if currentHashForFile != fileHashInPreviousCommit {
				changedFiles = append(changedFiles, Change{
					Filename: fileInPreviousCommit,
					Type:     ChangeTypeUpdated,
				})
			}
			continue
		}

		changedFiles = append(changedFiles, Change{
			Filename: fileInPreviousCommit,
			Type:     ChangeTypeDeleted,
		})
	}

	for fileInCurrentCommit := range currentCommitMap {
		if _, exists := previousCommitMap[fileInCurrentCommit]; exists {
			continue
		}

		changedFiles = append(changedFiles, Change{
			Filename: fileInCurrentCommit,
			Type:     ChangeTypeAdded,
		})
	}

	return changedFiles
}

type Change struct {
	Filename string
	Type     ChangeType
}

type ChangeType string

const (
	ChangeTypeAdded   ChangeType = "Added"
	ChangeTypeDeleted ChangeType = "Deleted"
	ChangeTypeUpdated ChangeType = "Updated"
)

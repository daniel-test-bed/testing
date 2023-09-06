package main

import (
	"fmt"
	"github.com/terraform-ibm-modules/ibmcloud-terratest-wrapper/common"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Git Root Path: ")
	path, err := common.GitRootPath(".")
	fmt.Println(path)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Wrapper PR Repo and Branch: ")
	defaultRepo, defaultBranch, err := common.GetCurrentPrRepoAndBranch()
	fmt.Println(defaultRepo, defaultBranch)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Wrapper Base Repo and Branch: ")
	baseRepo, baseBranch := common.GetBaseRepoAndBranch("", "")
	fmt.Println(baseRepo, baseBranch)

	fmt.Println("PR Repo and Branch: ")
	defaultRepo, defaultBranch, err = GetCurrentPRRepoAndBranch()
	fmt.Println(defaultRepo, defaultBranch)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Base Repo and Branch: ")
	baseRepo, baseBranch, err = GetBaseRepoAndBranch("", "")
	fmt.Println(baseRepo, baseBranch)
	if err != nil {
		fmt.Println(err)
	}

}

// GetDefaultRepoAndBranch returns the default repo and branch for the given repo path.
func GetCurrentPRRepoAndBranch() (string, string, error) {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	repoURL, err := cmd.Output()
	if err != nil {
		return "", "", err
	}

	cmd = exec.Command("git", "symbolic-ref", "--short", "HEAD")
	branch, err := cmd.Output()
	if err != nil {
		return "", "", err
	}

	return strings.TrimSpace(string(repoURL)), strings.TrimSpace(string(branch)), nil
}

// GetBaseRepoAndBranch returns the base repo and branch for the given repo and branch.
func GetBaseRepoAndBranch(repo string, branch string) (string, string, error) {
	// If both repo and branch are provided, return them directly
	if repo != "" && branch != "" {
		return repo, branch, nil
	}

	// Determine the URL of the upstream remote (usually "origin")
	if repo == "" {
		cmd := exec.Command("git", "remote", "get-url", "upstream")
		output, err := cmd.Output()
		if err != nil {
			// If there's no "upstream" remote, fall back to "origin"
			cmd := exec.Command("git", "remote", "get-url", "origin")
			output, err = cmd.Output()
			if err != nil {
				return "", "", err
			}
			repo = strings.TrimSpace(string(output))
		}
	}

	// Determine the branch name of the upstream remote (usually "origin/HEAD")
	if branch == "" {
		cmd := exec.Command("git", "remote", "show", "origin")
		output, err := cmd.Output()
		if err != nil {
			return "", "", err
		}

		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if strings.Contains(line, "HEAD branch:") {
				parts := strings.SplitN(line, ":", 2)
				if len(parts) == 2 {
					branch = strings.TrimSpace(parts[1])
					break
				}
			}
		}
	}

	// If branch information is not found, leave it as an empty string
	return repo, branch, nil
}

//// GetBaseRepoAndBranch returns the base repo and branch for the current repo.
//func GetBaseRepoAndBranch(repoURL string, branch string) (string, string, error) {
//	if repoURL == "" {
//		cmd := exec.Command("git", "remote", "get-url", "origin")
//		output, err := cmd.Output()
//		if err != nil {
//			return "", "", err
//		}
//		repoURL = strings.TrimSpace(string(output))
//	}
//
//	if branch == "" {
//		cmd := exec.Command("git", "symbolic-ref", "--short", "refs/remotes/origin/HEAD")
//		output, err := cmd.Output()
//		if err != nil {
//			return "", "", err
//		}
//		branch = strings.TrimPrefix(string(output), "origin/")
//	}
//
//	return repoURL, branch, nil
//}

//// GetDefaultRepoAndBranch returns the default repo and branch for the given repo path
//func GetDefaultRepoAndBranch(repoPath string) (string, string, error) {
//	repo, err := git.PlainOpen(repoPath)
//	if err != nil {
//		return "", "", err
//	}
//
//	// Get the current HEAD reference (branch)
//	headRef, err := repo.Head()
//	if err != nil {
//		return "", "", err
//	}
//
//	return repoPath, headRef.Name().Short(), nil
//}
//
//// GetBaseRepoAndBranch returns the base repo and branch for the current repo. If repoURL or branch are provided, they are used instead of the current repo.
//// The repoURL should be the origin repo URL and the branch should be the origin branch HEAD.
//func GetBaseRepoAndBranch(repoURL string, branch string) (string, string, error) {
//	if repoURL == "" {
//		repoURL = getOriginURL()
//	}
//
//	if branch == "" {
//		branch = getOriginBranch()
//	}
//
//	return repoURL, branch, nil
//}
//
//func getOriginURL() string {
//	repo, err := git.PlainOpen(".")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	remotes, err := repo.Remotes()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, remote := range remotes {
//		if remote.Config().Name == "origin" {
//			return remote.Config().URLs[0]
//		}
//	}
//
//	return ""
//}
//
//func getOriginBranch() string {
//	repo, err := git.PlainOpen(".")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Open the reference to origin/HEAD
//	ref, err := repo.Reference("refs/remotes/origin/HEAD", true)
//	if err != nil {
//		return ""
//	}
//
//	// Extract the branch name from the full reference name
//	branchName := strings.TrimPrefix(ref.Name().Short(), "origin/")
//	return branchName
//}

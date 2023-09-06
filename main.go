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

	fmt.Println("Wrapper Default Repo and Branch: ")
	defaultRepo, defaultBranch, err := common.GetDefaultRepoAndBranch(path)
	fmt.Println(defaultRepo, defaultBranch)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Wrapper Base Repo and Branch: ")
	baseRepo, baseBranch := common.GetBaseRepoAndBranch("", "")
	fmt.Println(baseRepo, baseBranch)

	fmt.Println("Default Repo and Branch: ")
	defaultRepo, defaultBranch, err = GetDefaultRepoAndBranch(path)
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
func GetDefaultRepoAndBranch(repoPath string) (string, string, error) {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	cmd.Dir = repoPath
	repoURL, err := cmd.Output()
	if err != nil {
		return "", "", err
	}

	cmd = exec.Command("git", "symbolic-ref", "--short", "HEAD")
	cmd.Dir = repoPath
	branch, err := cmd.Output()
	if err != nil {
		return "", "", err
	}

	return strings.TrimSpace(string(repoURL)), strings.TrimSpace(string(branch)), nil
}

// GetBaseRepoAndBranch returns the base repo and branch for the given repo and branch.
func GetBaseRepoAndBranch(repoURL string, branch string) (string, string, error) {

	// Determine the URL of the upstream remote (usually "origin")
	cmd := exec.Command("git", "remote", "get-url", "upstream")
	output, err := cmd.Output()
	if err != nil {
		// If there's no "upstream" remote, fall back to "origin"
		cmd := exec.Command("git", "remote", "get-url", "origin")
		output, err = cmd.Output()
		if err != nil {
			return "", "", err
		}
	}
	// Check if repoURL and branch are empty; if so, use default values
	if repoURL == "" {
		repoURL = strings.TrimSpace(string(output))
	}

	// Determine the branch name of the upstream remote (usually "origin")
	cmd = exec.Command("git", "rev-parse", "--abbrev-ref", "--symbolic-full-name", "@{u}")
	output, err = cmd.Output()
	if err != nil {
		return "", "", err
	}

	if branch == "" {
		branch = strings.TrimPrefix(string(output), "origin/")
	}

	return repoURL, branch, nil
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

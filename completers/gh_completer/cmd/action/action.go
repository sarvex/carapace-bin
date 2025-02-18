package action

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action/config"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action/ghrepo"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action/git"
	"github.com/spf13/cobra"
)

func ghExecutable() string {
	return "gh"
}

func ActionConfigHosts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if config, err := config.ParseDefaultConfig(); err != nil {
			return carapace.ActionMessage("failed to parse DefaultConfig: " + err.Error())
		} else {
			if hosts, err := config.Hosts(); err != nil {
				return carapace.ActionMessage("failed ot loadd hosts: " + err.Error())
			} else {
				return carapace.ActionValues(hosts...)
			}
		}
	})
}

func ApiV3Action(cmd *cobra.Command, query string, v interface{}, transform func() carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if repo, err := repoOverride(cmd); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return carapace.ActionExecCommand("gh", "api", "--hostname", repo.RepoHost(), "--preview", "mercy", query)(func(output []byte) carapace.Action {
				if err := json.Unmarshal(output, &v); err != nil {
					return carapace.ActionMessage("failed to unmarshall response: " + err.Error())
				}
				return transform()
			})
		}
	})
}

func GraphQlAction(cmd *cobra.Command, query string, v interface{}, transform func() carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		params := make([]string, 0)
		if strings.Contains(query, "$owner") {
			params = append(params, "$owner: String!")
		}
		if strings.Contains(query, "$repo") {
			params = append(params, "$repo: String!")
		}
		queryParams := strings.Join(params, ",")
		if queryParams != "" {
			queryParams = "(" + queryParams + ")"
		}

		if repo, err := repoOverride(cmd); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return carapace.ActionExecCommand(ghExecutable(), "api", "--hostname", repo.RepoHost(), "--header", "Accept: application/vnd.github.merge-info-preview+json", "graphql", "-F", "owner="+repo.RepoOwner(), "-F", "repo="+repo.RepoName(), "-f", fmt.Sprintf("query=query%v {%v}", queryParams, query))(func(output []byte) carapace.Action {
				if err := json.Unmarshal(output, &v); err != nil {
					return carapace.ActionMessage("failed to unmarshall response: " + err.Error())
				}
				return transform()
			})
		}
	})
}

func repoOverride(cmd *cobra.Command) (ghrepo.Interface, error) {
	repoOverride := ""
	if flag := cmd.Flag("repo"); flag != nil {
		repoOverride = flag.Value.String()
	}
	if repoFromEnv := os.Getenv("GH_REPO"); repoOverride == "" && repoFromEnv != "" {
		repoOverride = repoFromEnv
	}
	if repoOverride != "" {
		splitted := strings.Split(repoOverride, "/")
		switch len(splitted) {
		case 1:
			return ghrepo.New(splitted[0], ""), nil // assume owner
		case 2:
			if strings.Contains(splitted[0], ".") {
				return ghrepo.NewWithHost(splitted[1], "", splitted[0]), nil
			} else {
				return ghrepo.New(splitted[0], splitted[1]), nil
			}
		default:
			return ghrepo.NewWithHost(splitted[1], splitted[2], splitted[0]), nil
		}
	} else {
		if remotes, err := git.Remotes(); err == nil {
			for _, remote := range remotes {
				if remote.Resolved == "base" {
					return ghrepo.FromURL(remote.FetchURL) // TODO ssh translator
				}
			}
		}
	}
	return ghrepo.New("", ""), nil
}

func ActionAuthScopes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionValuesDescribed(
			"admin:gpg_key", "Fully manage GPG keys.",
			"admin:org", "Fully manage the organization and its teams, projects, and memberships.",
			"admin:org_hook", "Grants read, write, ping, and delete access to organization hooks.",
			"admin:public_key", "Fully manage public keys.",
			"admin:repo_hook", "Grants read, write, ping, and delete access to repository hooks in public and private repositories.",
			"codespace", "Full control of codespaces",
			"delete:packages", "Grants access to delete packages from GitHub Packages.",
			"delete_repo", "Grants access to delete adminable repositories.",
			"gist", "Grants write access to gists.",
			"notifications", "Grants read access to a user's notifications",
			"public_repo", "Limits access to public repositories.",
			"read:discussion", "Allows read access for team discussions.",
			"read:gpg_key", "List and view details for GPG keys.",
			"read:org", "Read-only access to organization membership, organization projects, and team membership.",
			"read:packages", "Grants access to download or install packages from GitHub Packages.",
			"read:public_key", "List and view details for public keys.",
			"read:repo_hook", "Grants read and ping access to hooks in public or private repositories.",
			"read:user", "Grants access to read a user's profile data.",
			"repo", "Grants full access to private and public repositories.",
			"repo:invite", "Grants accept/decline abilities for invitations to collaborate on a repository.",
			"repo:status", "Grants read/write access to public and private repository commit statuses.",
			"repo_deployment", "Grants access to deployment statuses for public and private repositories.",
			"security_events", "Grants read and write access to security events in the code scanning API.",
			"user", "Grants read/write access to profile info only.",
			"user:email", "Grants read access to a user's email addresses.",
			"user:follow", "Grants access to follow or unfollow other users.",
			"workflow", "Grants the ability to add and update GitHub Actions workflow files.",
			"write:discussion", "Allows read and write access for team discussions.",
			"write:gpg_key", "Create, list, and view details for GPG keys.",
			"write:org", "Read and write access to organization membership, organization projects, and team membership.",
			"write:packages", "Grants access to upload or publish a package in GitHub Packages.",
			"write:public_key", "Create, list, and view details for public keys.",
			"write:repo_hook", "Grants read, write, and ping access to hooks in public or private repositories.",
		).Invoke(c).ToMultiPartsA(":")
	})
}

// https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/
func byteCountSI(b int) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func defaultUser() (user string, err error) {
	var cfg config.Config
	if cfg, err = config.ParseDefaultConfig(); err == nil {
		var hosts []string
		if hosts, err = cfg.Hosts(); err == nil {
			if len(hosts) < 1 {
				err = errors.New("could not retrieve default user")
			} else {
				user, err = cfg.Get(hosts[0], "user")
			}
		}
	}
	return
}

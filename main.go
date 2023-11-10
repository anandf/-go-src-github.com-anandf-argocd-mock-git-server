package main

import (
	"io"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8341", &githubSCMMockHandler{})
	if err != nil {
		panic(err)
	}
	err = http.ListenAndServe(":8342", &githubPullMockHandler{})
	if err != nil {
		panic(err)
	}
}

type githubSCMMockHandler struct{}
type githubPullMockHandler struct{}

func (h githubSCMMockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.RequestURI {
	case "/api/v3/orgs/argoproj/repos?per_page=100":
		_, err := io.WriteString(w, `[
			{
			  "id": 1296269,
			  "node_id": "MDEwOlJlcG9zaXRvcnkxMjk2MjY5",
			  "name": "argo-cd",
			  "full_name": "argoproj/argo-cd",
			  "owner": {
				"login": "argoproj",
				"id": 1,
				"node_id": "MDQ6VXNlcjE=",
				"avatar_url": "https://github.com/images/error/argoproj_happy.gif",
				"gravatar_id": "",
				"url": "https://api.github.com/users/argoproj",
				"html_url": "https://github.com/argoproj",
				"followers_url": "https://api.github.com/users/argoproj/followers",
				"following_url": "https://api.github.com/users/argoproj/following{/other_user}",
				"gists_url": "https://api.github.com/users/argoproj/gists{/gist_id}",
				"starred_url": "https://api.github.com/users/argoproj/starred{/owner}{/repo}",
				"subscriptions_url": "https://api.github.com/users/argoproj/subscriptions",
				"organizations_url": "https://api.github.com/users/argoproj/orgs",
				"repos_url": "https://api.github.com/users/argoproj/repos",
				"events_url": "https://api.github.com/users/argoproj/events{/privacy}",
				"received_events_url": "https://api.github.com/users/argoproj/received_events",
				"type": "User",
				"site_admin": false
			  },
			  "private": false,
			  "html_url": "https://github.com/argoproj/argo-cd",
			  "description": "This your first repo!",
			  "fork": false,
			  "url": "https://api.github.com/repos/argoproj/argo-cd",
			  "archive_url": "https://api.github.com/repos/argoproj/argo-cd/{archive_format}{/ref}",
			  "assignees_url": "https://api.github.com/repos/argoproj/argo-cd/assignees{/user}",
			  "blobs_url": "https://api.github.com/repos/argoproj/argo-cd/git/blobs{/sha}",
			  "branches_url": "https://api.github.com/repos/argoproj/argo-cd/branches{/branch}",
			  "collaborators_url": "https://api.github.com/repos/argoproj/argo-cd/collaborators{/collaborator}",
			  "comments_url": "https://api.github.com/repos/argoproj/argo-cd/comments{/number}",
			  "commits_url": "https://api.github.com/repos/argoproj/argo-cd/commits{/sha}",
			  "compare_url": "https://api.github.com/repos/argoproj/argo-cd/compare/{base}...{head}",
			  "contents_url": "https://api.github.com/repos/argoproj/argo-cd/contents/{path}",
			  "contributors_url": "https://api.github.com/repos/argoproj/argo-cd/contributors",
			  "deployments_url": "https://api.github.com/repos/argoproj/argo-cd/deployments",
			  "downloads_url": "https://api.github.com/repos/argoproj/argo-cd/downloads",
			  "events_url": "https://api.github.com/repos/argoproj/argo-cd/events",
			  "forks_url": "https://api.github.com/repos/argoproj/argo-cd/forks",
			  "git_commits_url": "https://api.github.com/repos/argoproj/argo-cd/git/commits{/sha}",
			  "git_refs_url": "https://api.github.com/repos/argoproj/argo-cd/git/refs{/sha}",
			  "git_tags_url": "https://api.github.com/repos/argoproj/argo-cd/git/tags{/sha}",
			  "git_url": "git:github.com/argoproj/argo-cd.git",
			  "issue_comment_url": "https://api.github.com/repos/argoproj/argo-cd/issues/comments{/number}",
			  "issue_events_url": "https://api.github.com/repos/argoproj/argo-cd/issues/events{/number}",
			  "issues_url": "https://api.github.com/repos/argoproj/argo-cd/issues{/number}",
			  "keys_url": "https://api.github.com/repos/argoproj/argo-cd/keys{/key_id}",
			  "labels_url": "https://api.github.com/repos/argoproj/argo-cd/labels{/name}",
			  "languages_url": "https://api.github.com/repos/argoproj/argo-cd/languages",
			  "merges_url": "https://api.github.com/repos/argoproj/argo-cd/merges",
			  "milestones_url": "https://api.github.com/repos/argoproj/argo-cd/milestones{/number}",
			  "notifications_url": "https://api.github.com/repos/argoproj/argo-cd/notifications{?since,all,participating}",
			  "pulls_url": "https://api.github.com/repos/argoproj/argo-cd/pulls{/number}",
			  "releases_url": "https://api.github.com/repos/argoproj/argo-cd/releases{/id}",
			  "ssh_url": "git@github.com:argoproj/argo-cd.git",
			  "stargazers_url": "https://api.github.com/repos/argoproj/argo-cd/stargazers",
			  "statuses_url": "https://api.github.com/repos/argoproj/argo-cd/statuses/{sha}",
			  "subscribers_url": "https://api.github.com/repos/argoproj/argo-cd/subscribers",
			  "subscription_url": "https://api.github.com/repos/argoproj/argo-cd/subscription",
			  "tags_url": "https://api.github.com/repos/argoproj/argo-cd/tags",
			  "teams_url": "https://api.github.com/repos/argoproj/argo-cd/teams",
			  "trees_url": "https://api.github.com/repos/argoproj/argo-cd/git/trees{/sha}",
			  "clone_url": "https://github.com/argoproj/argo-cd.git",
			  "mirror_url": "git:git.example.com/argoproj/argo-cd",
			  "hooks_url": "https://api.github.com/repos/argoproj/argo-cd/hooks",
			  "svn_url": "https://svn.github.com/argoproj/argo-cd",
			  "homepage": "https://github.com",
			  "language": null,
			  "forks_count": 9,
			  "stargazers_count": 80,
			  "watchers_count": 80,
			  "size": 108,
			  "default_branch": "master",
			  "open_issues_count": 0,
			  "is_template": false,
			  "topics": [
				"argoproj",
				"atom",
				"electron",
				"api"
			  ],
			  "has_issues": true,
			  "has_projects": true,
			  "has_wiki": true,
			  "has_pages": false,
			  "has_downloads": true,
			  "archived": false,
			  "disabled": false,
			  "visibility": "public",
			  "pushed_at": "2011-01-26T19:06:43Z",
			  "created_at": "2011-01-26T19:01:12Z",
			  "updated_at": "2011-01-26T19:14:43Z",
			  "permissions": {
				"admin": false,
				"push": false,
				"pull": true
			  },
			  "template_repository": null
			}
		  ]`)
		if err != nil {
			panic(err)
		}
	case "/api/v3/repos/argoproj/argo-cd/branches?per_page=100":
		_, err := io.WriteString(w, `[
			{
			  "name": "master",
			  "commit": {
				"sha": "c5b97d5ae6c19d5c5df71a34c7fbeeda2479ccbc",
				"url": "https://api.github.com/repos/argoproj/argo-cd/commits/c5b97d5ae6c19d5c5df71a34c7fbeeda2479ccbc"
			  },
			  "protected": true,
			  "protection": {
				"required_status_checks": {
				  "enforcement_level": "non_admins",
				  "contexts": [
					"ci-test",
					"linter"
				  ]
				}
			  },
			  "protection_url": "https://api.github.com/repos/argoproj/hello-world/branches/master/protection"
			}
		  ]
		`)
		if err != nil {
			panic(err)
		}
	case "/api/v3/repos/argoproj/argo-cd/contents/pkg?ref=master":
		_, err := io.WriteString(w, `{
			"type": "file",
			"encoding": "base64",
			"size": 5362,
			"name": "pkg/",
			"path": "pkg/",
			"content": "encoded content ...",
			"sha": "3d21ec53a331a6f037a91c368710b99387d012c1",
			"url": "https://api.github.com/repos/octokit/octokit.rb/contents/README.md",
			"git_url": "https://api.github.com/repos/octokit/octokit.rb/git/blobs/3d21ec53a331a6f037a91c368710b99387d012c1",
			"html_url": "https://github.com/octokit/octokit.rb/blob/master/README.md",
			"download_url": "https://raw.githubusercontent.com/octokit/octokit.rb/master/README.md",
			"_links": {
			  "git": "https://api.github.com/repos/octokit/octokit.rb/git/blobs/3d21ec53a331a6f037a91c368710b99387d012c1",
			  "self": "https://api.github.com/repos/octokit/octokit.rb/contents/README.md",
			  "html": "https://github.com/octokit/octokit.rb/blob/master/README.md"
			}
		  }`)
		if err != nil {
			panic(err)
		}
	case "/api/v3/repos/argoproj/argo-cd/branches/master":
		_, err := io.WriteString(w, `{
			"name": "master",
			"commit": {
			  "sha": "c5b97d5ae6c19d5c5df71a34c7fbeeda2479ccbc",
			  "url": "https://api.github.com/repos/octocat/Hello-World/commits/c5b97d5ae6c19d5c5df71a34c7fbeeda2479ccbc"
			},
			"protected": true,
			"protection": {
			  "required_status_checks": {
				"enforcement_level": "non_admins",
				"contexts": [
				  "ci-test",
				  "linter"
				]
			  }
			},
			"protection_url": "https://api.github.com/repos/octocat/hello-world/branches/master/protection"
		  }`)
		if err != nil {
			panic(err)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func (h githubPullMockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.RequestURI {
	case "/api/v3/repos/applicationset-test-org/argocd-example-apps/pulls?per_page=100":
		_, err := io.WriteString(w, `[
  {
    "number": 1,
    "labels": [
      {
        "name": "preview"
      }
    ],
	"base": {
		"ref": "master",
		"sha": "7a4a5c987fdfb2b0629e9dbf5f31636c69ba4775"
	},
    "head": {
      "ref": "pull-request",
      "sha": "824a5c987fdfb2b0629e9dbf5f31636c69ba4772"
    }
  }
]`)
		if err != nil {
			panic(err)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

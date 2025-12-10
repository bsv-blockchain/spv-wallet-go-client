<div align="center">

# 🧰&nbsp;&nbsp;spv-wallet-go-client

**Lightweight Go client library for interacting with the SPV Wallet API.**

<br/>

<a href="https://github.com/bsv-blockchain/spv-wallet-go-client/releases"><img src="https://img.shields.io/github/release-pre/bsv-blockchain/spv-wallet-go-client?include_prereleases&style=flat-square&logo=github&color=black" alt="Release"></a>
<a href="https://golang.org/"><img src="https://img.shields.io/github/go-mod/go-version/bsv-blockchain/spv-wallet-go-client?style=flat-square&logo=go&color=00ADD8" alt="Go Version"></a>
<a href="https://github.com/bsv-blockchain/spv-wallet-go-client/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-OpenBSV-blue?style=flat-square" alt="License"></a>

<br/>

<table align="center" border="0">
  <tr>
    <td align="right">
       <code>CI / CD</code> &nbsp;&nbsp;
    </td>
    <td align="left">
       <a href="https://github.com/bsv-blockchain/spv-wallet-go-client/actions"><img src="https://img.shields.io/github/actions/workflow/status/bsv-blockchain/spv-wallet-go-client/fortress.yml?branch=main&label=build&logo=github&style=flat-square" alt="Build"></a>
       <a href="https://github.com/bsv-blockchain/spv-wallet-go-client/actions"><img src="https://img.shields.io/github/last-commit/bsv-blockchain/spv-wallet-go-client?style=flat-square&logo=git&logoColor=white&label=last%20update" alt="Last Commit"></a>
    </td>
    <td align="right">
       &nbsp;&nbsp;&nbsp;&nbsp; <code>Quality</code> &nbsp;&nbsp;
    </td>
    <td align="left">
       <a href="https://goreportcard.com/report/github.com/bsv-blockchain/spv-wallet-go-client"><img src="https://goreportcard.com/badge/github.com/bsv-blockchain/spv-wallet-go-client?style=flat-square" alt="Go Report"></a>
       <a href="https://codecov.io/gh/bsv-blockchain/spv-wallet-go-client"><img src="https://codecov.io/gh/bsv-blockchain/spv-wallet-go-client/branch/main/graph/badge.svg?style=flat-square" alt="Coverage"></a>
    </td>
  </tr>

  <tr>
    <td align="right">
       <code>Security</code> &nbsp;&nbsp;
    </td>
    <td align="left">
       <a href="https://scorecard.dev/viewer/?uri=github.com/bsv-blockchain/spv-wallet-go-client"><img src="https://api.scorecard.dev/projects/github.com/bsv-blockchain/spv-wallet-go-client/badge?style=flat-square" alt="Scorecard"></a>
       <a href=".github/SECURITY.md"><img src="https://img.shields.io/badge/policy-active-success?style=flat-square&logo=security&logoColor=white" alt="Security"></a>
    </td>
    <td align="right">
       &nbsp;&nbsp;&nbsp;&nbsp; <code>Community</code> &nbsp;&nbsp;
    </td>
    <td align="left">
       <a href="https://github.com/bsv-blockchain/spv-wallet-go-client/graphs/contributors"><img src="https://img.shields.io/github/contributors/bsv-blockchain/spv-wallet-go-client?style=flat-square&color=orange" alt="Contributors"></a>
       <a href="https://github.com/sponsors/bsv-blockchain"><img src="https://img.shields.io/badge/sponsor-BSV-181717.svg?logo=github&style=flat-square" alt="Sponsor"></a>
    </td>
  </tr>
</table>

</div>

<br/>
<br/>

<div align="center">

### <code>Project Navigation</code>

</div>

<table align="center">
  <tr>
    <td align="center" width="33%">
       📦&nbsp;<a href="#-installation"><code>Installation</code></a>
    </td>
    <td align="center" width="33%">
       🧪&nbsp;<a href="#-examples--tests"><code>Examples&nbsp;&&nbsp;Tests</code></a>
    </td>
    <td align="center" width="33%">
       📚&nbsp;<a href="#-documentation"><code>Documentation</code></a>
    </td>
  </tr>
  <tr>
    <td align="center">
       🤝&nbsp;<a href="#-contributing"><code>Contributing</code></a>
    </td>
    <td align="center">
       🛠️&nbsp;<a href="#-code-standards"><code>Code&nbsp;Standards</code></a>
    </td>
    <td align="center">
       ⚡&nbsp;<a href="#-benchmarks"><code>Benchmarks</code></a>
    </td>
  </tr>
  <tr>
    <td align="center">
       🤖&nbsp;<a href="#-ai-compliance"><code>AI&nbsp;Compliance</code></a>
    </td>
    <td align="center">
       📝&nbsp;<a href="#-license"><code>License</code></a>
    </td>
    <td align="center">
       👥&nbsp;<a href="#-maintainers"><code>Maintainers</code></a>
    </td>
  </tr>
</table>
<br/>

## 📦 Installation

**spv-wallet-go-client** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get -u github.com/bsv-blockchain/spv-wallet-go-client
```

<br/>

## 📚 Documentation

- **API Reference** – Dive into the godocs at [pkg.go.dev/github.com/bsv-blockchain/spv-wallet-go-client](https://pkg.go.dev/github.com/bsv-blockchain/spv-wallet-go-client)
- **Usage Examples** – Browse practical patterns either the [examples directory](examples) or the example tests
- **Test Suite** – Review both the unit tests and fuzz tests(powered by [`testify`](https://github.com/stretchr/testify))
- **SPV Wallet Docs** - please refer to the [SPV Wallet Documentation](https://docs.bsvblockchain.org/network-topology/spv-wallet)

<br/>

<details>
<summary><strong><code>Quick Start Guide</code></strong></summary>

## Quick Start Guide

The implementation enforces separation of concerns by isolating admin and non-admin APIs, requiring separate initialization for their respective clients. This ensures clarity and modularity when utilizing the exposed functionality.

### `UserAPI` Initialization Methods:

### 1. [`NewUserAPIWithAccessKey`](/user_api.go#L468)
- **Description:** Initializes a `UserAPI` instance using an access key for authentication.
- **Note:** Requests made with this instance will be securely signed, ensuring integrity and authenticity.

### 2. [`NewUserAPIWithXPriv`](/user_api.go#L449)
- **Description:** Initializes a `UserAPI` instance using an extended private key (xPriv) for authentication.
- **Note:** Requests made with this instance will also be securely signed.
- **Recommendation:** This option offers a high level of security, making it a preferred choice alongside the access key option.

### 3. [`NewUserAPIWithXPub`](/user_api.go#L435)
- **Description:** Initializes a `UserAPI` instance using an extended public key (xPub).
- **Note:** Requests made with this instance will not be signed.
- **Security Advisory:** For enhanced security, it is strongly recommended to use either `NewUserAPIWithAccessKey` or `NewUserAPIWithXPriv` instead, as unsigned requests may be less secure.


### `AdminAPI` Initialization Methods:

### 1. [`NewAdminAPIWithXPriv`](/admin_api.go#L375)
- **Description:** Initializes a `AdminAPI` instance using an extended private key (xPriv) for authentication.
- **Note:** Requests made with this instance will be securely signed, ensuring integrity and authenticity.

### 2. [`NewAdminAPIWithXPub`](/admin_api.go#L390)
- **Description:** Initializes a `AdminAPI` instance using an extended public key (xPub).
- **Note:** Requests made with this instance will not be signed.
- **Security Advisory:** For enhanced security, it is strongly recommended to use either `NewAdminAPIWithXPriv`instead, as unsigned requests may be less secure.

**Code snippets:**
- [AdminAPI example](/examples/admin_add_user/admin_add_user.go)
- [UserAPI example](/examples/list_transactions/list_transactions.go)

</details>

<details>
<summary><strong><code>Compatibility & Support</code></strong></summary>

### Compatibility and Support

#### Deprecation Notice
The client **does not support** the following:
- **Admin and non-admin old endpoints** of the SPV Wallet API based on the `/v1/` prefix.
- Deprecated methods for building query parameters for HTTP requests.

#### Current Compatibility
The client is designed for full compatibility with the newer `/api/v1/` endpoints exposed by the SPV Wallet API. It focuses on aligning with the latest standards and structure provided by the API.

</details>

<details>
<summary><strong><code>API Endpoints Compatibility</code></strong></summary>

### API Admin Endpoints Compatibility

#### Access Keys API
| HTTP Method | Endpoint                 | Action             | Support Status | API Code                                                        | Pagination |
|-------------|--------------------------|--------------------|----------------|-----------------------------------------------------------------|------------|
| GET         | /api/v1/admin/users/keys | Search access keys | ✅              | [API](/internal/api/v1/admin/accesskeys/access_keys_api.go#L25) | ✅          |

#### Contacts API
| HTTP Method | Endpoint                             | Action          | Support Status | API Code                                                   | Pagination |
|-------------|--------------------------------------|-----------------|----------------|------------------------------------------------------------|------------|
| GET         | /api/v1/admin/contacts               | Search contacts | ✅              | [API](/internal/api/v1/admin/contacts/contacts_api.go#L42) | ✅          |
| POST        | /api/v1/admin/contacts/confirmations | Confirm contact | ✅              | [API](/internal/api/v1/admin/contacts/contacts_api.go#L83) | ❌          |
| PUT         | /api/v1/admin/contacts/{id}          | Update contact  | ✅              | [API](/internal/api/v1/admin/contacts/contacts_api.go#L68) | ❌          |
| DELETE      | /api/v1/admin/contacts/{id}          | Delete contact  | ✅              | [API](/internal/api/v1/admin/contacts/contacts_api.go#L95) | ❌          |
| POST        | /api/v1/admin/contacts/{paymail}     | Create contact  | ✅              | [API](/internal/api/v1/admin/contacts/contacts_api.go#L27) | ❌          |

#### Invitations API
| HTTP Method | Endpoint                       | Action            | Support Status | API Code                                                         | Pagination |
|-------------|--------------------------------|-------------------|----------------|------------------------------------------------------------------|------------|
| POST        | /api/v1/admin/invitations/{id} | Accept invitation | ✅              | [API](/internal/api/v1/admin/invitations/invitations_api.go#L22) | ❌          |
| DELETE      | /api/v1/admin/invitations/{id} | Reject invitation | ✅              | [API](/internal/api/v1/admin/invitations/invitations_api.go#L35) | ❌          |


#### Paymails API
| HTTP Method | Endpoint                    | Action           | Support Status | API Code                                                   | Pagination |
|-------------|-----------------------------|------------------|----------------|------------------------------------------------------------|------------|
| GET         | /api/v1/admin/paymails      | Search paymails  | ✅              | [API](/internal/api/v1/admin/paymails/paymails_api.go#L73) | ✅          |
| POST        | /api/v1/admin/paymails      | Create paymail   | ✅              | [API](/internal/api/v1/admin/paymails/paymails_api.go#L44) | ❌          |
| GET         | /api/v1/admin/paymails/{id} | Retrieve paymail | ✅              | [API](/internal/api/v1/admin/paymails/paymails_api.go#L59) | ❌          |
| DELETE      | /api/v1/admin/paymails/{id} | Delete paymail   | ✅              | [API](/internal/api/v1/admin/paymails/paymails_api.go#L27) | ❌          |

#### Stats API
| HTTP Method | Endpoint            | Action         | Support Status | API Code                                             | Pagination |
|-------------|---------------------|----------------|----------------|------------------------------------------------------|------------|
| GET         | /api/v1/admin/stats | Retrieve stats | ✅              | [API](/internal/api/v1/admin/stats/stats_api.go#L23) | ✅          |

#### Status API
| HTTP Method | Endpoint             | Action          | Support Status | API Code                                               | Pagination |
|-------------|----------------------|-----------------|----------------|--------------------------------------------------------|------------|
| GET         | /api/v1/admin/status | Retrieve status | ✅              | [API](/internal/api/v1/admin/status/status_api.go#L23) | ❌          |

#### Transactions API
| HTTP Method | Endpoint                        | Action               | Support Status | API Code                                                           | Pagination |
|-------------|---------------------------------|----------------------|----------------|--------------------------------------------------------------------|------------|
| GET         | /api/v1/admin/transactions      | Search transactions  | ✅              | [API](/internal/api/v1/admin/transactions/transactions_api.go#L39) | ✅          |
| GET         | /api/v1/admin/transactions/{id} | Retrieve transaction | ✅              | [API](/internal/api/v1/admin/transactions/transactions_api.go#L26) | ❌          |

#### UTXOs API
| HTTP Method | Endpoint            | Action       | Support Status | API Code                                             | Pagination |
|-------------|---------------------|--------------|----------------|------------------------------------------------------|------------|
| GET         | /api/v1/admin/utxos | Search UTXOs | ✅              | [API](/internal/api/v1/admin/utxos/utxos_api.go#L25) | ✅          |

#### Webhooks API
| HTTP Method | Endpoint                             | Action               | Support Status | API Code                                                   | Pagination |
|-------------|--------------------------------------|----------------------|----------------|------------------------------------------------------------|------------|
| GET         | /api/v1/admin/webhooks/subscriptions | Subscribe to webhook | ✅              | [API](/internal/api/v1/admin/webhooks/webhooks_api.go#L23) | ❌          |
| DELETE      | /api/v1/admin/webhooks/subscriptions | Unsubscribe webhook  | ✅              | [API](/internal/api/v1/admin/webhooks/webhooks_api.go#L36) | ❌          |

#### XPubs API
| HTTP Method | Endpoint            | Action       | Support Status | API Code                                             | Pagination |
|-------------|---------------------|--------------|----------------|------------------------------------------------------|------------|
| GET         | /api/v1/admin/users | Search XPubs | ✅              | [API](/internal/api/v1/admin/xpubs/xpubs_api.go#L41) | ✅          |
| POST        | /api/v1/admin/users | Create XPub  | ✅              | [API](/internal/api/v1/admin/xpubs/xpubs_api.go#L27) | ❌          |

</details>

<details>
<summary><strong><code>API Non-Admin Endpoints Compatibility</code></strong></summary>

### API Non-Admin Endpoints Compatibility

#### Access Keys API
| HTTP Method | Endpoint                        | Action              | Support Status | API Code                                                      | Pagination |
|-------------|---------------------------------|---------------------|----------------|---------------------------------------------------------------|------------|
| GET         | /api/v1/users/current/keys      | Search access keys  | ✅              | [API](/internal/api/v1/user/accesskeys/access_key_api.go#L56) | ✅          |
| POST        | /api/v1/users/current/keys      | Create access key   | ✅              | [API](/internal/api/v1/user/accesskeys/access_key_api.go#L27) | ❌          |
| GET         | /api/v1/users/current/keys/{id} | Retrieve access key | ✅              | [API](/internal/api/v1/user/accesskeys/access_key_api.go#L42) | ❌          |
| DELETE      | /api/v1/users/current/keys/{id} | Revoke access key   | ✅              | [API](/internal/api/v1/user/accesskeys/access_key_api.go#L82) | ❌          |

#### Contacts API
| HTTP Method | Endpoint                   | Action            | Support Status | API Code                                                   | Pagination |
|-------------|----------------------------|-------------------|----------------|------------------------------------------------------------|------------|
| GET         | /api/v1/contacts           | Search contacts   | ✅              | [API](/internal/api/v1/user/contacts/contacts_api.go#L27)  | ✅          |
| GET         | /api/v1/contacts/{paymail} | Retrieve contact  | ✅              | [API](/internal/api/v1/user/contacts/contacts_api.go#L53)  | ❌          |
| PUT         | /api/v1/contacts/{paymail} | Upsert contact    | ✅              | [API](/internal/api/v1/user/contacts/contacts_api.go#L67)  | ❌          |
| DELETE      | /api/v1/contacts/{paymail} | Remove contact    | ✅              | [API](/internal/api/v1/user/contacts/contacts_api.go#L89)  | ❌          |
| POST        | /api/v1/contacts/{paymail} | Confirm contact   | ✅              | [API](/internal/api/v1/user/contacts/contacts_api.go#L101) | ❌          |
| DELETE      | /api/v1/contacts/{paymail} | Unconfirm contact | ✅              | [API](/internal/api/v1/user/contacts/contacts_api.go#L113) | ❌          |

#### Invitations API
| HTTP Method | Endpoint                               | Action            | Support Status | API Code                                                        | Pagination |
|-------------|----------------------------------------|-------------------|----------------|-----------------------------------------------------------------|------------|
| POST        | /api/v1/invitations/{paymail}/contacts | Accept invitation | ✅              | [API](/internal/api/v1/user/invitations/invitations_api.go#L22) | ❌          |
| DELETE      | /api/v1/invitations/{paymail}          | Reject invitation | ✅              | [API](/internal/api/v1/user/invitations/invitations_api.go#L34) | ❌          |

#### Merkle Roots API
| HTTP Method | Endpoint            | Action              | Support Status | API Code                                                        | Pagination |
|-------------|---------------------|---------------------|----------------|-----------------------------------------------------------------|------------|
| GET         | /api/v1/merkleroots | Search Merkle roots | ✅              | [API](/internal/api/v1/user/merkleroots/merkleroots_api.go#L36) | ❌          |

#### Paymails API
| HTTP Method | Endpoint         | Action          | Support Status | API Code                                                  | Pagination |
|-------------|------------------|-----------------|----------------|-----------------------------------------------------------|------------|
| GET         | /api/v1/paymails | Search paymails | ✅              | [API](/internal/api/v1/user/paymails/paymails_api.go#L25) | ✅          |

#### Transactions API
| HTTP Method | Endpoint                    | Action               | Support Status | API Code                                                           | Pagination |
|-------------|-----------------------------|----------------------|----------------|--------------------------------------------------------------------|------------|
| GET         | /api/v1/transactions        | Search transactions  | ✅              | [API](/internal/api/v1/user/transactions/transactions_api.go#L137) | ✅          |
| POST        | /api/v1/transactions        | Record transaction   | ✅              | [API](/internal/api/v1/user/transactions/transactions_api.go#L93)  | ❌          |
| POST        | /api/v1/transactions/drafts | Draft transaction    | ✅              | [API](/internal/api/v1/user/transactions/transactions_api.go#L78)  | ❌          |
| GET         | /api/v1/transactions/{id}   | Retrieve transaction | ✅              | [API](/internal/api/v1/user/transactions/transactions_api.go#L123) | ❌          |
| PATCH       | /api/v1/transactions/{id}   | Update transaction   | ✅              | [API](/internal/api/v1/user/transactions/transactions_api.go#L108) | ❌          |

#### UTXOs API
| HTTP Method | Endpoint      | Action       | Support Status | API Code                                            | Pagination |
|-------------|---------------|--------------|----------------|-----------------------------------------------------|------------|
| GET         | /api/v1/utxos | Search UTXOs | ✅              | [API](/internal/api/v1/user/utxos/utxos_api.go#L25) | ❌          |

#### XPubs API
| HTTP Method | Endpoint              | Action                     | Support Status | API Code                                           | Pagination |
|-------------|-----------------------|----------------------------|----------------|----------------------------------------------------|------------|
| GET         | /api/v1/users/current | Retrieve current user info | ✅              | [API](/internal/api/v1/user/xpubs/xpub_api.go#L24) | ❌          |
| PATCH       | /api/v1/users/current | Update current user info   | ✅              | [API](/internal/api/v1/user/xpubs/xpub_api.go#L24) | ❌          |


</details>


<details>
<summary><strong><code>Development Build Commands</code></strong></summary>
<br/>

Get the [MAGE-X](https://github.com/mrz1836/mage-x) build tool for development:
```shell script
go install github.com/mrz1836/mage-x/cmd/magex@latest
```

View all build commands

```bash script
magex help
```

</details>

<details>
<summary><strong><code>Repository Features</code></strong></summary>
<br/>

* **Continuous Integration on Autopilot** with [GitHub Actions](https://github.com/features/actions) – every push is built, tested, and reported in minutes.
* **Pull‑Request Flow That Merges Itself** thanks to [auto‑merge](.github/workflows/auto-merge-on-approval.yml) and hands‑free [Dependabot auto‑merge](.github/workflows/dependabot-auto-merge.yml).
* **One‑Command Builds** powered by battle‑tested [MAGE-X](https://github.com/mrz1836/mage-x) targets for linting, testing, releases, and more.
* **First‑Class Dependency Management** using native [Go Modules](https://github.com/golang/go/wiki/Modules).
* **Uniform Code Style** via [gofumpt](https://github.com/mvdan/gofumpt) plus zero‑noise linting with [golangci‑lint](https://github.com/golangci/golangci-lint).
* **Confidence‑Boosting Tests** with [testify](https://github.com/stretchr/testify), the Go [race detector](https://blog.golang.org/race-detector), crystal‑clear [HTML coverage](https://blog.golang.org/cover) snapshots, and automatic uploads to [Codecov](https://codecov.io/).
* **Hands‑Free Releases** delivered by [GoReleaser](https://github.com/goreleaser/goreleaser) whenever you create a [new Tag](https://git-scm.com/book/en/v2/Git-Basics-Tagging).
* **Relentless Dependency & Vulnerability Scans** via [Dependabot](https://dependabot.com), [Nancy](https://github.com/sonatype-nexus-community/nancy) and [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck).
* **Security Posture by Default** with [CodeQL](https://docs.github.com/en/github/finding-security-vulnerabilities-and-errors-in-your-code/about-code-scanning), [OpenSSF Scorecard](https://openssf.org) and secret‑leak detection via [gitleaks](https://github.com/gitleaks/gitleaks).
* **Automatic Syndication** to [pkg.go.dev](https://pkg.go.dev/) on every release for instant godoc visibility.
* **Polished Community Experience** using rich templates for [Issues & PRs](https://docs.github.com/en/communities/using-templates-to-encourage-useful-issues-and-pull-requests/configuring-issue-templates-for-your-repository).
* **All the Right Meta Files** (`LICENSE`, `CONTRIBUTING.md`, `CODE_OF_CONDUCT.md`, `SUPPORT.md`, `SECURITY.md`) pre‑filled and ready.
* **Code Ownership** clarified through a [CODEOWNERS](.github/CODEOWNERS) file, keeping reviews fast and focused.
* **Zero‑Noise Dev Environments** with tuned editor settings (`.editorconfig`) plus curated *ignore* files for [VS Code](.editorconfig), [Docker](.dockerignore), and [Git](.gitignore).
* **Label Sync Magic**: your repo labels stay in lock‑step with [.github/labels.yml](.github/labels.yml).
* **Friendly First PR Workflow** – newcomers get a warm welcome thanks to a dedicated [workflow](.github/workflows/pull-request-management.yml).
* **Standards‑Compliant Docs** adhering to the [standard‑readme](https://github.com/RichardLitt/standard-readme/blob/master/spec.md) spec.
* **Instant Cloud Workspaces** via [Gitpod](https://gitpod.io/) – spin up a fully configured dev environment with automatic linting and tests.
* **Out‑of‑the‑Box VS Code Happiness** with a preconfigured [Go](https://code.visualstudio.com/docs/languages/go) workspace and [`.vscode`](.vscode) folder with all the right settings.
* **Optional Release Broadcasts** to your community via [Slack](https://slack.com), [Discord](https://discord.com), or [Twitter](https://twitter.com) – plug in your webhook.
* **AI Compliance Playbook** – machine‑readable guidelines ([AGENTS.md](.github/AGENTS.md), [CLAUDE.md](.github/CLAUDE.md), [.cursorrules](.cursorrules), [sweep.yaml](.github/sweep.yaml)) keep ChatGPT, Claude, Cursor & Sweep aligned with your repo's rules.
* **Go-Pre-commit System** - [High-performance Go-native pre-commit hooks](https://github.com/mrz1836/go-pre-commit) with 17x faster execution—run the same formatting, linting, and tests before every commit, just like CI.
* **Zero Python Dependencies** - Pure Go implementation with environment-based configuration via [.env.base](.github/.env.base).
* **DevContainers for Instant Onboarding** – Launch a ready-to-code environment in seconds with [VS Code DevContainers](https://containers.dev/) and the included [.devcontainer.json](.devcontainer.json) config.

</details>

<details>
<summary><strong><code>Library Deployment</code></strong></summary>
<br/>

This project uses [goreleaser](https://github.com/goreleaser/goreleaser) for streamlined binary and library deployment to GitHub. To get started, install it via:

```bash
brew install goreleaser
```

The release process is defined in the [.goreleaser.yml](.goreleaser.yml) configuration file.


Then create and push a new Git tag using:

```bash
magex version:bump push=true bump=patch branch=main
```

This process ensures consistent, repeatable releases with properly versioned artifacts and citation metadata.

</details>

<details>
<summary><strong><code>Pre-commit Hooks</code></strong></summary>
<br/>

Set up the Go-Pre-commit System to run the same formatting, linting, and tests defined in [AGENTS.md](.github/AGENTS.md) before every commit:

```bash
go install github.com/mrz1836/go-pre-commit/cmd/go-pre-commit@latest
go-pre-commit install
```

The system is configured via [.env.base](.github/.env.base) and can be customized using also using [.env.custom](.github/.env.custom) and provides 17x faster execution than traditional Python-based pre-commit hooks. See the [complete documentation](http://github.com/mrz1836/go-pre-commit) for details.

</details>

<details>
<summary><strong><code>GitHub Workflows</code></strong></summary>
<br/>

### 🎛️ The Workflow Control Center

All GitHub Actions workflows in this repository are powered by a single configuration files – your one-stop shop for tweaking CI/CD behavior without touching a single YAML file! 🎯

**Configuration Files:**
- **[.env.base](.github/.env.base)** – Default configuration that works for most Go projects
- **[.env.custom](.github/.env.custom)** – Optional project-specific overrides

This magical file controls everything from:
- **⚙️ Go version matrix** (test on multiple versions or just one)
- **🏃 Runner selection** (Ubuntu or macOS, your wallet decides)
- **🔬 Feature toggles** (coverage, fuzzing, linting, race detection, benchmarks)
- **🛡️ Security tool versions** (gitleaks, nancy, govulncheck)
- **🤖 Auto-merge behaviors** (how aggressive should the bots be?)
- **🏷️ PR management rules** (size labels, auto-assignment, welcome messages)

<br/>

| Workflow Name                                                                      | Description                                                                                                            |
|------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------|
| [auto-merge-on-approval.yml](.github/workflows/auto-merge-on-approval.yml)         | Automatically merges PRs after approval and all required checks, following strict rules.                               |
| [codeql-analysis.yml](.github/workflows/codeql-analysis.yml)                       | Analyzes code for security vulnerabilities using [GitHub CodeQL](https://codeql.github.com/).                          |
| [dependabot-auto-merge.yml](.github/workflows/dependabot-auto-merge.yml)           | Automatically merges [Dependabot](https://github.com/dependabot) PRs that meet all requirements.                       |
| [fortress.yml](.github/workflows/fortress.yml)                                     | Runs the GoFortress security and testing workflow, including linting, testing, releasing, and vulnerability checks.    |
| [pull-request-management.yml](.github/workflows/pull-request-management.yml)       | Labels PRs by branch prefix, assigns a default user if none is assigned, and welcomes new contributors with a comment. |
| [scorecard.yml](.github/workflows/scorecard.yml)                                   | Runs [OpenSSF](https://openssf.org/) Scorecard to assess supply chain security.                                        |
| [stale.yml](.github/workflows/stale-check.yml)                                     | Warns about (and optionally closes) inactive issues and PRs on a schedule or manual trigger.                           |
| [sync-labels.yml](.github/workflows/sync-labels.yml)                               | Keeps GitHub labels in sync with the declarative manifest at [`.github/labels.yml`](./.github/labels.yml).             |

</details>

<details>
<summary><strong><code>Updating Dependencies</code></strong></summary>
<br/>

To update all dependencies (Go modules, linters, and related tools), run:

```bash
magex deps:update
```

This command ensures all dependencies are brought up to date in a single step, including Go modules and any tools managed by [MAGE-X](https://github.com/mrz1836/mage-x). It is the recommended way to keep your development environment and CI in sync with the latest versions.

</details>

<br/>

## 🧪 Examples & Tests

All unit tests and [examples](examples) run via [GitHub Actions](https://github.com/bsv-blockchain/spv-wallet-go-client/actions) and use [Go version 1.24.x](https://go.dev/doc/go1.24). View the [configuration file](.github/workflows/fortress.yml).

Run all tests (fast):

```bash script
magex test
```

Run all tests with race detector (slower):
```bash script
magex test:race
```

<br/>

## ⚡ Benchmarks

Run the Go benchmarks:

```bash script
magex bench
```

<br/>

## 🛠️ Code Standards
Read more about this Go project's [code standards](.github/CODE_STANDARDS.md).

<br/>

## 🤖 AI Compliance
This project documents expectations for AI assistants using a few dedicated files:

- [AGENTS.md](.github/AGENTS.md) — canonical rules for coding style, workflows, and pull requests used by [Codex](https://chatgpt.com/codex).
- [CLAUDE.md](.github/CLAUDE.md) — quick checklist for the [Claude](https://www.anthropic.com/product) agent.
- [.cursorrules](.cursorrules) — machine-readable subset of the policies for [Cursor](https://www.cursor.so/) and similar tools.
- [sweep.yaml](.github/sweep.yaml) — rules for [Sweep](https://github.com/sweepai/sweep), a tool for code review and pull request management.

Edit `AGENTS.md` first when adjusting these policies, and keep the other files in sync within the same pull request.

<br/>

## 👥 Maintainers
| [<img src="https://github.com/mrz1836.png" height="50" width="50" alt="MrZ" />](https://github.com/mrz1836) | [<img src="https://github.com/icellan.png" height="50" alt="Siggi" />](https://github.com/icellan) |
|:-----------------------------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------------------:|
|                                      [MrZ](https://github.com/mrz1836)                                      |                                [Siggi](https://github.com/icellan)                                 |

<br/>

## 🤝 Contributing
View the [contributing guidelines](.github/CONTRIBUTING.md) and please follow the [code of conduct](.github/CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:!
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:.

[![Stars](https://img.shields.io/github/stars/bsv-blockchain/spv-wallet-go-client?label=Please%20like%20us&style=social&v=1)](https://github.com/bsv-blockchain/spv-wallet-go-client/stargazers)

<br/>

## 📝 License

[![License](https://img.shields.io/badge/license-OpenBSV-blue?style=flat&logo=springsecurity&logoColor=white)](LICENSE)

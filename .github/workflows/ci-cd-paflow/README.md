### Reusable? Workflows for CI/CD

This repository contains a set of reusable GitHub Actions workflows to standardize and simplify CI/CD processes. These workflows can be used across multiple repositories for building and promoting Docker images.

### Available Workflows

1. **Build PR Docker Image**
Builds and pushes Docker images tagged for individual Pull Requests (e.g., `pr-<number>`).

2. **Build MAIN Docker Image**
Builds and pushes a stable Docker image for the `main` branch.

3. **Promote to Environment**
Promotes a stable Docker image (`stable`) to specific environments (e.g., `uat` or `prod`).

4. **Orchestrator Workflow** 
Coordinates the above workflows dynamically based on triggers: Pull Requests, `main` branch pushes, or manual actions.

### Example Usage
Here’s how to invoke the **Orchestrator Workflow**:


```
name: Call Workflow

on:
  pull_request:
    branches:
      - "*"
  push:
    branches:
      - main
  workflow_dispatch:
    inputs:
      action:
        description: "Specify promotion target (promote-uat or promote-prod)"
        required: true
        type: choice
        options:
          - promote-uat
          - promote-prod
      stable_tag:
        description: "The stable Docker tag to promote"
        required: false
        default: "stabilisssssimo"

jobs:
  orchestrator:
    name: Invoke OctoPussy Workflow
    uses: pagopa/eng-github-actions-iac-template/.github/workflows/orchestrator.yml@v1.0.0
    with:
      github_token: ${{ secrets.GITHUB_TOKEN }}
      action: ${{ inputs.action }}
      stable_tag: ${{ inputs.stable_tag }}
```
### Inputs for the Orchestrator

| Input | Description | Default |
| --- | --- | --- |
| `action` | Target promotion action (, ). `promote-uat``promote-prod` | None |
| `stable_tag` | The Docker tag to promote. | `stable` |

### How to Trigger Workflows
1. On **Pull Request**: Automatically builds and tags a Docker image for the PR.
2. On **Push to Main**: Automatically creates and pushes a stable Docker image.
3. **Manually**: Navigate to the repository’s Actions tab and trigger the workflow with custom inputs.

This repository offers **centralized, reusable workflows** for scalable and maintainable CI/CD pipelines. Use these workflows to ensure consistency across repositories! 🚀

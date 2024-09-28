# Contributing to the pgEdge Terraform Provider

We welcome and appreciate contributions from the community! This document outlines the process for contributing to the pgEdge Terraform Provider project.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Setting Up Your Development Environment](#setting-up-your-development-environment)
3. [Making Changes](#making-changes)
4. [Running Tests](#running-tests)
5. [Submitting a Pull Request](#submitting-a-pull-request)
6. [Code Style and Guidelines](#code-style-and-guidelines)

## Prerequisites

Before you begin, ensure you have the following tools installed:

- [Go](https://golang.org/doc/install) (version 1.20 or later)
- [Terraform CLI](https://developer.hashicorp.com/terraform/downloads)
- [Python 3](https://www.python.org/downloads/)
- [Swagger Go](https://github.com/go-swagger/go-swagger)
- [golangci-lint](https://golangci-lint.run/usage/install/)

## Setting Up Your Development Environment

1. Fork the repository on GitHub.

2. Clone your fork locally:
   ```
   git clone https://github.com/pgEdge/terraform-provider-pgedge.git
   cd terraform-provider-pgedge
   ```

3. Install Dependencies:
   ```
   make setup
   ```

4. Create a `.terraformrc` file in your home directory and copy the contents from `example.terraformrc`:
   ```
   cp example.terraformrc ~/.terraformrc
   ```

5. Set the required environment variables:
   ```
   export PGEDGE_BASE_URL="https://api.pgedge.com"
   export PGEDGE_CLIENT_ID="your-client-id"
   export PGEDGE_CLIENT_SECRET="your-client-secret"
   export PGEDGE_ROLE_ARN="your-role-arn" # optional. For running Tests only.
   ```

## Making Changes

1. Create a new branch for your changes:
   ```
   git checkout -b feature/your-feature-name
   ```

2. Make your changes to the code.

3. After making changes, it's crucial to build the provider for the changes to take effect:
   ```
   go install .
   ```
   Note: Remember to run this command every time you make changes to ensure they are reflected in the provider.

4. Run the linter to ensure code quality:
   ```
   make lint
   ```

5. Update or add tests as necessary.

6. Update documentation if you've made changes to the provider's functionality.

Remember, the `go install .` step is essential for your changes to be reflected in the provider. Always run this command after making modifications to see the effects of your changes.

## Running Tests

To run the acceptance tests:

```
make test
```

Note: Running acceptance tests will create real resources in your pgEdge account.

## Submitting a Pull Request

1. Push your changes to your fork:
   ```
   git push origin feature/your-feature-name
   ```

2. Go to the original repository on GitHub and create a new pull request.

3. Ensure your PR description clearly describes the problem and solution. Include the relevant issue number if applicable.

4. Wait for the maintainers to review your PR. Make any requested changes.

## Code Style and Guidelines

- We use [terraform-plugin-framework](https://developer.hashicorp.com/terraform/plugin/framework) for this provider. Familiarize yourself with its conventions and best practices.
- Follow Go best practices and conventions.
- Write clear, concise commit messages.
- Keep your changes focused. If you're addressing multiple issues, consider creating separate PRs.
- Add or update tests to cover your changes.
- Update documentation, including the README if necessary.

Thank you for contributing to the pgEdge Terraform Provider!
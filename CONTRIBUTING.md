# Contributing to Micromachine

Thank you for your interest in contributing to Micromachine! We appreciate your
time and effort in making this project better. This document outlines the
guidelines for contributing to this repository. Following these guidelines
helps us maintain a clean and collaborative environment.

## How to Contribute

There are many ways to contribute to Micromachine, including:

* **Reporting Bugs:** Help us identify and fix issues by reporting bugs.
* **Suggesting Enhancements:** Share your ideas for improving Micromachine.
* **Submitting Pull Requests:** Contribute code changes, whether bug fixes or
new features.
* **Improving Documentation:** Help us make the documentation clearer and more
comprehensive.

## Reporting Bugs

When reporting a bug, please include the following information:

* **Version:** Specify the version of Micromachine you are using.
* **Steps to Reproduce:** Provide clear and concise steps to reproduce the
issue.
* **Expected Behavior:** Describe what you expected to happen.
* **Actual Behavior:** Describe what actually happened.
* **Relevant Code Snippets:** If possible, include code snippets that
demonstrate the issue.
* **Environment:** Specify your operating system, Go version (`go version`),
and any relevant dependencies.

## Suggesting Enhancements

When suggesting an enhancement, please include the following information:

* **Description:** Clearly describe the proposed enhancement.
* **Motivation:** Explain why this enhancement would be beneficial.
* **Example:** If possible, provide an example of how the enhancement would be
used.

## Submitting Pull Requests

Before submitting a pull request, please make sure you have:

1. **Forked the repository:** Create a copy of the repository in your GitHub
   account.
2. **Created a branch:** Create a new branch for your changes. Use a
   descriptive name (e.g., `fix-issue-123`, `add-new-feature`).
3. **Made your changes:** Implement your changes, ensuring they adhere to the
   project's coding style (see below).
4. **Tested your changes:** Thoroughly test your changes to ensure they work as
   expected and don't introduce new issues.  Run the tests (`go test ./...`).
If applicable, add new tests to cover your changes.
5. **Updated documentation:** If your changes affect the documentation, update
   it accordingly.
6. **Committed your changes:** Use clear and concise commit messages.  Follow
   the conventional commits format (e.g., `feat: Add new feature`, `fix: Fix
bug #123`).
7. **Pushed your changes:** Push your branch to your forked repository.
8. **Created a pull request:** Submit a pull request to the main branch of the
   original repository.

## Coding Style

Please adhere to the following coding style guidelines:

* **`go fmt`:**  Run `go fmt` on your code to ensure consistent formatting.
* **Effective Go:** Follow the guidelines in the [Effective
Go](https://go.dev/doc/effective_go) document.
* **Meaningful Names:** Use descriptive and meaningful names for variables,
functions, and types.
* **Comments:** Write clear and concise comments to explain complex logic or
non-obvious code.  GoDoc style comments are encouraged.

## Testing

Ensure that your changes are covered by unit tests. Write new tests if
necessary. Run the tests using `go test ./...` before submitting a pull
request.

## Documentation

Keep the documentation up-to-date. If you make changes that affect the
documentation, update it accordingly.  Consider using GoDoc style comments for
package and function documentation.

## Code of Conduct

This project adheres to the [Contributor Covenant Code of
Conduct](https://www.contributor-covenant.org/version/2/0/code_of_conduct/). By
participating in this project, you are expected to uphold this code.

## License

By contributing to this project, you agree that your contributions will be
licensed under the [MIT License](./LICENSE).

## Questions?

If you have any questions or need help, please don't hesitate to open an issue
or reach out to the maintainers.

Thank you for contributing to Micromachine!

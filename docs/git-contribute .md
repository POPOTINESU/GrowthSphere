# Git Contribute Document

## Commit Message Convention

This project follows the **Conventional Commits** standard.  
Writing consistent commit messages helps us:

- Understand changes more easily.
- Generate release notes automatically.
- Identify and group changes by type.

---

## Commit Message Prefix

When committing, **always use one of the prefixes below**:

| Prefix   | Purpose                        | Example                            |
|----------|--------------------------------|------------------------------------|
| `feat`   | Add new features               | `feat: implement user login`      |
| `fix`    | Bug fixes                      | `fix: resolve user auth issue`    |
| `docs`   | Documentation updates          | `docs: update README usage guide` |
| `style`  | Code formatting (no logic)     | `style: reformat CSS`             |
| `perf`   | Performance improvements       | `perf: optimize image loading`    |
| `chore`  | Build system, CI/CD tasks      | `chore: update npm dependencies`  |
| `refactor` | Code refactoring (no behavior change) | `refactor: simplify auth logic` |
| `test`   | Add or update tests            | `test: add unit tests for login`  |

---

## Commit Message Format

Use the following structure for your commit messages:

```plaintext
<type>(<scope>): <description>
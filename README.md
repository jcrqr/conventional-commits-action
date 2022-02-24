# conventional-commits-action

This action validates a Pull Request title and commit messages against
[Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) guidelines

A valid title or commit message must follow this structure:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

## Usage

```yaml
- name: Conventional Commits
  uses: crqra/conventional-commits-action@main
```

## Example

```yaml
name: Test

on:
  pull_request:
    branches: [main]
    types: [opened, reopened, synchronize, edited, ready_for_review]

jobs:
  conventional-commits:
    runs-on: ubuntu-latest
    steps:
      - name: Conventional Commits
        uses: crqra/conventional-commits-action@main
```

## Contributing

Thanks for considering contributing.

- [Open an issue](https://github.com/crqra/conventional-commits-action/issues) if you have a problem or found a bug
- [Open a Pull Request](https://github.com/crqra/conventional-commits-action/pulls) if you have a suggestion, improvement or bug fix

## License

This project is released under the [MIT License](LICENSE).
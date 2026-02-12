# Contributing to Todo CLI

Thank you for your interest in contributing to Todo CLI! This document provides guidelines and instructions for contributing.

## Code of Conduct

- Be respectful and inclusive
- Welcome newcomers and help them get started
- Focus on constructive feedback
- Prioritize the community and project health

## Getting Started

1. **Fork the repository**
   ```bash
   git clone https://github.com/GourangaDasSamrat/todo-cli-go.git
   cd todo-cli
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Create a branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

## Development Guidelines

### Code Style

- Follow Go best practices and idioms
- Run `go fmt` before committing
- Use meaningful variable and function names
- Add comments for complex logic
- Keep functions focused and small

### Project Structure

```
todo-cli/
├── cmd/                # CLI command implementations
├── internal/           # Private application code
│   ├── models/        # Data structures
│   ├── storage/       # Persistence layer
│   ├── ui/            # User interface utilities
│   └── utils/         # Helper functions
└── pkg/               # Public packages
    ├── filter/        # Filtering logic
    └── sort/          # Sorting algorithms
```

### Commit Messages

Use clear, descriptive commit messages:

```
feat: add search by multiple tags
fix: correct due date parsing for different formats
docs: update README with new examples
refactor: simplify task filtering logic
test: add tests for priority sorting
```

Prefix types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `refactor`: Code refactoring
- `test`: Test additions/changes
- `chore`: Maintenance tasks

### Testing

- Write tests for new features
- Ensure existing tests pass
- Run tests with: `go test ./...`
- Aim for good test coverage

### Pull Request Process

1. **Update documentation**
   - Update README.md if needed
   - Add comments to complex code

2. **Test your changes**
   ```bash
   go test ./...
   go build
   ./build/todo --help
   ```

3. **Create Pull Request**
   - Provide clear description
   - Reference related issues
   - Include screenshots if UI changes
   - Request review from maintainers

4. **Address feedback**
   - Respond to review comments
   - Make requested changes
   - Push updates to your branch

## Feature Requests

Have an idea? Open an issue with:
- Clear description of the feature
- Use cases and examples
- Why it would be valuable
- Potential implementation approach

## Bug Reports

Found a bug? Open an issue with:
- Clear title and description
- Steps to reproduce
- Expected vs actual behavior
- Environment details (OS, Go version)
- Error messages or logs

## Development Tips

### Building

```bash
# Standard build
make build

# Install locally
make install

# Build for all platforms
make build-all

# Clean build artifacts
make clean
```

### Testing Features

```bash
# Test adding tasks
./build/todo add -t "Test task" -p high

# Test listing
./build/todo list

# Test interactive mode
./build/todo interactive
```

### Debugging

- Use `fmt.Printf` for debugging during development
- Check `~/.todo-cli/` for data files
- Review error messages carefully

## Areas for Contribution

We welcome contributions in:

- **Features**: New commands, filters, or sorting options
- **UI/UX**: Better colors, formatting, or interactive elements
- **Performance**: Optimization of filters, sorting, or storage
- **Documentation**: Examples, guides, or API docs
- **Tests**: Unit tests, integration tests, or test utilities
- **Bug fixes**: Resolve open issues

## Questions?

- Open an issue for questions
- Check existing issues and PRs
- Review the README and documentation

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

Thank you for contributing! 🎉
# Basemark

**Basemark** is a CLI benchmarking tool designed for use in Continuous Integration (CI) environments. It allows you to measure performance metrics over multiple iterations and verify them against predefined acceptance criteria.

[![SLSA Go releaser](https://github.com/shabinesh/basemark/actions/workflows/go-ossf-slsa3-publish.yml/badge.svg)](https://github.com/shabinesh/basemark/actions/workflows/go-ossf-slsa3-publish.yml)

## Features
- **Iteration-based Testing**: Run a specific number of iterations to assess performance consistency.
- **Acceptance Criteria**: Validate performance against a target acceptance time.
- **Variance Threshold**: Define acceptable variance levels for performance results.

## Installation
Download the prebuilt binary for your platform or install using Go:

```bash
go install github.com/shabinesh/basemark@latest
```

## Usage

```bash
basemark [global options] command [command options]
```

### Global Options

| Option                         | Alias | Description                                   | Default |
|--------------------------------|-------|-----------------------------------------------|---------|
| `--iterations value`           | `-n`  | Number of iterations to run                  | `0`     |
| `--acceptance-time value`      | `-t`  | Acceptance time in milliseconds              | `0`     |
| `--variance value`             | `-v`  | Acceptable variance from acceptance time (%) | `0`     |
| `--help`                       | `-h`  | Show help                                    |         |

### Commands

| Command    | Alias | Description                           |
|------------|-------|---------------------------------------|
| `help`     | `h`   | Shows a list of commands or help for one command |

### Example
Run a benchmark with 10 iterations, an acceptance time of 500ms, and a 5% variance threshold:

```bash
basemark -n 10 -t 500 -v 5 sleep 2
```

### Output
The tool will provide results for each iteration and highlight any deviations from the defined thresholds.

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests to improve the tool.

## License
This project is licensed under the [MIT License](LICENSE).

---

Feel free to adapt the sections to match the exact functionality of your CLI tool or its specific repository details.
# Introduction

This is repo contains codes for CLI app that is built to help ease workload of cyblerand developers. It is built using
go-lang with Cobra framework.

## Features

- GitHub Actions YAML file generation

# Installation

- Download the latest release from GitHub releases. Right side of this page.
- Save it to a directory of your choice
- (**Only for Linux/Unix**) Run the following command to make the file executable

```bash
chmod +x cyberland-cli
```

# Usage

## GitHub Actions YAML file generation

### Command

Run the following command to generate a GitHub Actions YAML file pre-configured for the project

```bash
./cyberland-cli generate --use-case 1 --project-name bigheart
```

- Two files will be generated in the directory where the command is run.
- Place the two files in the .github/workflows directory of the project.

### Parameters

| Parameter      | Description                                                                                                     | Type    | Accepted Values |
|----------------|-----------------------------------------------------------------------------------------------------------------|---------|-----------------|
| --use-case     | Business use-case number                                                                                        | Integer | `1, 2, 3`       |
| --project-name | Name of project to generate GitHub actions YAML file for. <br/> **Project name must be same on terraform side** | String  | `anystring`     |

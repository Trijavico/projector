# Projector

## What is it?
**Projector** is a command line tool designed to manage configurations and data hierarchically in a directory structure. It allows you to store, retrieve and delete specific configurations for different directories, facilitating the management of development, deployment and collaboration environments.

## Features
- **Configuration Hierarchy**: Stores configurations hierarchically as key pair values, allowing configurations with same keys at lower levels to overwrite those at higher levels.

- **Modularity**: Facilitates modularity and reuse of configurations in different environments and directories.

## Installation
To install and use **Projector**, you need to have Go installed on your system.

1. Clone this repository:
 ```bash
 git clone https://github.com/Trijavico/projector.git
 ```

2. Navigate to the project directory:
 ```bash
 cd projector
 ```

3. Build the executable:
 ```bash
 go build -o projector cmd/main.go
 ```

4. Add the executable to your `$PATH` (optional):
 ```bash
 export PATH=$PATH:/path/to/projector
 ```

## Usage
### Print All Settings
Prints all accumulated configurations from the current directory to the root of the file system.

```bash
projector
```

### Get the Value of a Key
Gets the value of a specific key in the current directory.

```bash
projector <key>
```

### Add a Configuration
Adds a key and its value to the current directory's configuration file.

```bash
projector add <key> <value>
```

### Delete a Configuration
Removes a key from the current directory's configuration file.

```bash
projector rm <key>
```

## Example

### Add a Configuration

```bash
projector add "api_key" "123456"
```

### Print All Configurations 

```bash
projector 

output: {"projector": {"directory/path": {"api_key": "123456"}}}
```

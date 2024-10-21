# futil - File Utility

`futil` is a simple command-line utility written in Go for performing various file operations such as counting lines and calculating checksums.

## Features

- Line count of files
- Checksum calculation (MD5, SHA1, SHA256)

## Design

The application is built using the Cobra library for creating powerful modern CLI applications. It follows a command-based structure with the following main commands:

- `linecount`: Counts the number of lines in a file
- `checksum`: Calculates the checksum of a file using specified algorithms
- `version`: Displays the version information
- `help`: Provides help information for the application and its commands

## Dependencies

- [github.com/spf13/cobra](https://github.com/spf13/cobra): Used for creating the CLI application structure and handling command-line arguments.

## Building and Running

### Run from latest version

1. Download and save file as `futil`: https://aqua-tropical-skink-958.mypinata.cloud/ipfs/QmRk1X5bnr5jWR8buwZa2PhNgxfQQwcnFd1kijQX3G9bLL
2. Navigate to the directory containing the file
3. Run the application:
   ```
   ./futil [command] [flags]
   ```

### Build and run local
1. Ensure you have Go installed on your system.
2. Clone this repository:
   ```
   git clone https://github.com/Sotatek-HuyDao/golang-practice-101.git
   ```
3. Navigate to the project directory:
   ```
   cd golang-practice-101
   ```
4. Build the application:
   ```
   make run-build
   ```
5. Run the application:
   ```
   ./futil [command] [flags]
   
   // or using go run command
   go run ./cmd/app/main.go [command] [flags]
   ```

### Note
1. Use `futil` command instead of `./futil`: Copy futil file to system's $PATH. Example in Linux
   ```
   sudo cp futil /usr/bin
   ```
2. Default version is `0.0.1`, you can change version when build the application using below command:
   ```
   make run-build VERSION={your version}
   ```

## Usage Examples

1. Count lines in a file:
   ```
   ./futil linecount -f myfile.txt
   ```

2. Calculate MD5 checksum of a file:
   ```
   ./futil checksum -f myfile.txt --md5
   ```

3. Read from stdin:
   ```
   cat myfile.txt | ./futil linecount -f -
   ```

4. Show version information:
   ```
   ./futil version
   ```

5. Display help:
   ```
   ./futil help
   ```

## Features Not Yet Implemented

- Concurrent processing for large files

## Known Issues

- The application may have performance issues with extremely large files.

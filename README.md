# Ollama CLI Chat

![Platforms](https://img.shields.io/badge/platforms-linux%20%7C%20macOS%20%7C%20windows-blue?style=flat-square)
>  ![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go)
> ![Status](https://img.shields.io/badge/status-active-brightgreen?style=flat-square)
> ![Build](https://img.shields.io/badge/build-passing-success?style=flat-square)

> ![Search Enabled](https://img.shields.io/badge/Search-Enabled-green?style=flat-square&logo=google)


> A clean, minimal **CLI application** to interact with local Ollama models — featuring session management, chat persistence, and resume functionality.


## Features

- **Start New Chats** with any supported local Ollama model.
- **Save Chat Sessions** automatically with timestamp and model.
- **Resume Previous Chats** seamlessly — continue from where you left off.
- **Delete Chats** with interactive selection.
- **View Past Chat Messages** before resuming sessions.
- Fully persistent and local – **no cloud or external API** dependency.
- Built with `cobra` for a user-friendly CLI experience.
- **Internet-augmented responses** using SerpAPI for real-time search within chats
- Supports **streamed response generation** via `Ollama` API.

## Installation
> ** Option 1: Download Prebuilt Binaries (Recommended)**

Available for: Windows, macOS (Intel & ARM), Linux

Go to the Releases directory and download the binary for your OS:


> OS Files
- Windows	     `ollama-cli_windows_amd64.exe`
- macOS	`ollama-cli_darwin_amd64 or ollama-cli_darwin_arm64`
- Linux	`ollama-cli_linux_amd64`

#### For macOS 
- ollama-cli_darwin_amd64 (Intel)
- ollama-cli_darwin_arm64 (Apple Silicon: M1, M2, M3)

After downloading:

Before running the cli make sure the ollama is running.
```
# (macOS/Linux)
chmod +x ollama-cli_darwin_amd64 or ollama_cli_darwin_arm64 
or ollama-cli_linux_amd64

./ollama-cli_darwin_amd64 or ./ollama_cli_darwin_arm64 
or ./ollama_cli_linux_amd64
```

```
# (Windows)
Open cmd and navigate to Directory where the .exe is.
.\ollama-cli_windows_amd64.exe
```
> ** Option 2: Build from Source** 

> **Requirements**: Go 1.22+, [Ollama](https://ollama.com) installed and running locally.

**Important**: Ollama **must be running locally on the default URL** (`http://localhost:11434`) for this CLI to function correctly.

Start Ollama in a separate terminal before using the CLI:

```bash
ollama serve
```
Then build and run the CLI:
```bash
git clone https://github.com/ashankgupta/ollama-cli.git
cd ollama-cli
go build -o ollama-cli
./ollama-cli

```
Or run directly:
```bash
go run .
```
## Usage
Start the CLI:
```bash
./ollama-cli chat
```
You'll see an interactive menu:
```bash
What would you like to do?
[1] Start a new chat
[2] Continue a previous chat
[3] Delete a chat
[4] Exit
```
List saved chats:
```bash
./ollama-cli list
```
## Chat Storage
- Chats are saved locally in the .chats directory.
- Filenames include session name and timestamp.
- All history is preserved in JSON format.

## Usage

To search within a chat, type:

```
search: What is the latest update on NVIDIA GPUs?
```
The CLI will:
- Fetch top search results from SerpAPI
- Display the result in your chat
- Automatically append it into the context

#### Set Your API Key
You can set your SerpAPI key in this way:
```
./ollama-cli set-key keyhere
```
#### How to get API Key?

-  Go to https://serpapi.com and Register or Sign in
-  Go to Api Key Section
-  Copy the Private API Key and Done.

## Commands Overview
Command	Description
``` bash
chat	Start, resume, or delete chat sessions
list	View all saved chat sessions
set-key Set your SerpAPI key
```
## Configuration

You can choose the model interactively when starting a chat, or predefine it using:

```
./ollama-cli chat --model llama3
```
## Project Structure
```
cmd/
├── api.go                # Api Logic
├── chat.go               # Chat command & interaction
├── chat_handler.go       # Chat loop logic
├── config.go             # Save the SERP Api key
├── list_chats.go         # List command
├── listmodels.go         # List All Local Models
├── model_picker.go       # Model selection logic
├── root.go               # CLI root setup
├── setkey.go             # Setting Key Logic Here
├── storage.go            # Chat persistence (save/load/delete)
├── types.go              # Message and Session struct definitions
├── websearch.go          # Web Search Logic
```             

## FAQ

#### Q: Can I use this without internet?
```
Yes, all chats run locally if your model is downloaded.
```
#### Q: Does this support OpenAI or only Ollama?
```
It is made for Ollama Only.
```

##  Under the Hood (Tech Stack)
```
- Go 1.22
- spf13/cobra for CLI structure
- Ollama API for model interaction
- JSON-based local chat persistence
```

## License
Licensed under the MIT License.

## Acknowledgments
- Powered by Ollama
- CLI framework: spf13/cobra

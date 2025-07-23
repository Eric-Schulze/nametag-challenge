<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#project-outline">Project Outline</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

This project is a submission for the Nametag coding challenge as part of the interview process. It is in response to the following prompt:
>Write a program that updates itself. Imagine that you have a program you’ve deployed to clients and you will periodically produce new versions. When a new version is produced, we want the deployed programs to be seamlessly replaced by a new version.
Please write what you consider to be production quality code, whatever that means to you. You may choose any common programming language you like (Go, Rust, or C++ would be good choices). Your program should reasonably be expected to work across common desktop operating systems (Windows, Mac, Linux). If your scheme requires non-trivial server components, please write those as well.

The entire project is written in Go, with YAML config files and a Makefile for helper commands.

### Project Outline

The project is broken up into two main components, world-pop (the CLI) and updater-server (the remote server).

#### world-pop

The CLI portion of the project is a very simple program that displays population data for a given country. See [Usages](#usage) for details about the specific commands available.

The version for world-pop is generated from 'ldflags' during the build process.

#### updater-server

The backend server of the project is a simple REST API meant to serve metadata and binaries for the CLI. 

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/eric-schulze/nametag-challenge.git
   ```
2. From the root, start the updater-server
   ```sh
   make start-server
   ```
   This will run the API server in the current shell. 
3. Open a new shell to the root and create version 1.0.0 of the CLI
   ```sh
   make reset-versions
   ```
4. Install the CLI locally
   ```sh
   make install-local-world-pop
   ```

Everything should be installed and ready to run.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->
## Usage

### world-pop CLI

The world pop CLI has 5 main commands: 
- `world-pop -v` to show the current version of the module
- `world-pop country {name or code}` display population data of the specified country
  - Example
  ```sh
  world-pop country Egypt
  ```
- `world-pop latest` prints out the latest version of the CLI that is available on the server
- `world-pop check-update` compares the current version of the CLI to the latest version available on the server
- `world-pop update` updates the currently installed CLI module to the latest available version, if it is not up to date

### Makefile directives

- `make start-server` builds the updater-server and runs it in the current shell
- `make reset-versions` resets both the local and the server modules back to version 1.0.0
- `make increase-remote-version` installs a new version of the CLI app on the updater-server with the patch version increased by 1 from the currently installed version
- `make toggle-auto-update` toggles the auto update setting in the world-pop/internal/init/config.yaml; when true, the cli will perform an update check and self update every time a command is issued
- `make toggle-logging` toggles development logs setting in the world-pop/internal/init/config.yaml

### Example testing commands

1. Verify the app is working correctly
  ```sh
  world-pop country Egypt
  ```
2. Check the currently installed version
  ```sh
  world-pop -v
  ```
3. Check the latest available version on the server
  ```sh
  world-pop latest
  ```
4. Verify same version again with app
  ```sh
  world-pop check-update
  ```
5. Increase version on the server
  ```sh
  make increase-remote-version
  ```
6. Compare versions again
  ```sh
  world-pop check-update
  ```
  If at this point, the app auto updates, it means that the auto update setting is turned on. To turn it off, run `make toggle-auto-update`, then increase the version on the server one more time.
5. Manually update the local CLI module
  ```sh
  world-pop update
  ```
5. Check that new version has been installed
  ```sh
  world-pop -v
  ```
  or
  ```sh
  world-pop check-update
  ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Eric Schulze - eric.schulze001@gmail.com

Project Link: [https://github.com/Eric-Schulze/nametag-challenge](https://github.com/Eric-Schulze/nametag-challenge)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

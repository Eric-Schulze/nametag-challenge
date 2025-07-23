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

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* npm
  ```sh
  npm install npm@latest -g
  ```

### Installation

1. Get a free API Key at [https://example.com](https://example.com)
2. Clone the repo
   ```sh
   git clone https://github.com/eric-schulze/nametag-challenge.git
   ```
3. Install NPM packages
   ```sh
   npm install
   ```
4. Enter your API in `config.js`
   ```js
   const API_KEY = 'ENTER YOUR API';
   ```
5. Change git remote url to avoid accidental pushes to base project
   ```sh
   git remote set-url origin eric-schulze/nametag-challenge
   git remote -v # confirm the changes
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->
## Usage

Use this space to show useful examples of how a project can be used. Additional screenshots, code examples and demos work well in this space. You may also link to more resources.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Eric Schulze - eric.schulze001@gmail.com

Project Link: [https://github.com/Eric-Schulze/nametag-challenge](https://github.com/Eric-Schulze/nametag-challenge)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

# SSH Chess

This project implements a chess game server accessible via SSH, with a chess bot that is trained to mimic my chess playing style using data from my `chess.com` account.

WIP: This project is still in development.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Local Development Setup](#local-development-setup)
3. [Deployment Process](#deployment-process)
4. [Project Structure](#project-structure)
5. [Environment Variables](#environment-variables)
6. [Troubleshooting](#troubleshooting)

## Prerequisites

To work with this project, you'll need:

- Docker and Docker Compose
- Git

To deploy this project, you'll need:

- An Ubuntu server with 'ssh' access credentials and AWS CLI, Docker and Docker Compose
- An AWS account with an ECR registry and repository created for this project

## Local Development Setup

1. Clone the repository:

   ```
   git clone https://github.com/your-username/ssh-chess.git
   cd ssh-chess
   ```

2. Build and run the project locally:

   ```
   make build start
   ```

   This will start the SSH server, chess engine, and SQLite database services.

3. Access the chess server via SSH on localhost:8080 (or the port specified in your docker-compose.yml):

   ```
   ssh 127.0.0.1 -p 8080
   ```

## Usage

See the `Makefile` for other commands to manage the local development environment, including commands to clean, build, start, start in detached mode, stop, and hot-reload individual services.

## Deployment Process

The project uses GitHub Actions for CI/CD. Here's an overview of the deployment process:

1. Push changes to the `main` branch to trigger the workflow.

2. The GitHub Action will:

   - Build Docker images for ssh-server, chess-engine, and sqlite-db
   - Push these images to Amazon ECR
   - Deploy the updated services to an Ubuntu server

3. Ensure the following secrets are set in your GitHub repository:

   - `AWS_ACCESS_KEY_ID` (credentials of an AWS user with ECR access)
   - `AWS_SECRET_ACCESS_KEY`
   - `SSH_SERVER_HOST_KEY` (The SSH host key content to be used by the deployed SSH server)
   - `USER` (Ubuntu instance username)
   - `IP` (Ubuntu instance IP)
   - `KEY` (SSH key for Ubuntu instance)

4. The deployment script will update the Docker Compose file on the Ubuntu instance and restart the services.

## Project Structure

- `ssh-server/`: Contains the Dockerised SSH server implementation
- `chess-engine/`: Contains the Dockerised chess engine implementation for evalutaion / move generation
- `sqlite-db/`: Manages the SQLite database for game storage
- `.github/workflows/`: Contains the GitHub Actions workflow file
- `docker-compose.dev.yml`: Defines the local development environment
- `docker-compose.prod.yml`: Defines the production environment setup

## Environment Variables

The following environment variables are used in the project:

- `ECR_REPOSITORY`: The name of your ECR repository
- `STAGE`: The deployment stage (e.g., 'prod', 'dev')
- `AWS_REGION`: The AWS region for your ECR

These are set in the GitHub Actions workflow `.deploy.yml` file.

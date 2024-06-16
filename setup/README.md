# Baby Name Guesser Setup Guide

This guide will walk you through the steps to set up the Baby Name Guesser application on a server.

## Prerequisites

Before you begin, make sure you have the following prerequisites installed:

- Ansible
```bash
sudo apt update
sudo apt install software-properties-common
sudo add-apt-repository --yes --update ppa:ansible/ansible
sudo apt install ansible
```

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/baby-name-guesser.git
    ```

2. Navigate to the project directory:

    ```bash
    cd baby-name-guesser
    ```

3. Run the ansible playbook:

    ```bash
    npm install
    ```

4. Configure the environment variables:

    - In /etc/banague/environment.conf 
    - Adapt the environemnt variables to your needs:

5. Start the application:

    ```bash
    systemctl start banague.service
    ```

## Usage

Once the application is set up and running, you can access it by navigating to `http://localhost:8080/game` in your web browser.


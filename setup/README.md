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
    - Add the following environment variables and provide the necessary values:

      ```plaintext
      PORT=3000
      MONGODB_URI=mongodb://localhost:27017/baby-name-guesser
      ```

5. Start the application:

    ```bash
    npm start
    ```

## Usage

Once the application is set up and running, you can access it by navigating to `http://localhost:3000` in your web browser.

## Contributing

If you would like to contribute to the Baby Name Guesser application, please follow the guidelines outlined in the [CONTRIBUTING.md](link-to-contributing-file) file.

## License

This project is licensed under the [MIT License](link-to-license-file).
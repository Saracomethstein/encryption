
# Encryption Web Application

This project is a web application that allows users to encrypt and decrypt text using various SHA algorithms. The application is built with a Go backend and a simple HTML/CSS frontend.

## Features

- Encrypt text using SHA-256, SHA-384, and SHA-512 algorithms
- Decrypt text by looking up the original text from a list of previously encrypted values
- Maintain a history of encrypted/decrypted values
- Simple and clean user interface

## Requirements

- Go 1.14+
- PostgreSQL 16
- A web browser

## Installation

1. Clone the repository:

```bash
git clone https://github.com/Saracomethstein/encryption.git
cd encryption
```

2. Run the Go application:

```bash
make run
```

The application will start on `http://localhost:8000`.

## Usage

### Encrypting Text

1. Open the application in your web browser.
2. Select the encryption function (`Encrypt`) by clicking the corresponding button.
3. Select the SHA algorithm (SHA-256, SHA-384, or SHA-512) by clicking the corresponding button.
4. Enter the text you want to encrypt in the `input` field.
5. Click the `START` button.
6. The encrypted text will appear in the `output` field and will be added to the history table.

### Decrypting Text

1. Open the application in your web browser.
2. Select the decryption function (`Decrypt`) by clicking the corresponding button.
3. Enter the SHA value you want to decrypt in the `input` field.
4. Click the `START` button.
5. If the SHA value exists in the history, the original text will appear in the `output` field and will be added to the history table. If the SHA value does not exist, a message indicating the absence of the key will be shown.

## Project Structure

```
encryption/
├── app/                    # Main Go application
│   └── server.go
├── database/               # Database connection and commands
│   └── dbConnection.go
├── encrypt/                # Encryption logic
│   └── encrypt.go
├── materials/              # Database dump
│   └── dbScrypt.sql
└── ui/                     # Frontend files
    ├── index.html
    ├── style.css
    └── script.js

```

## Contributing

1. Fork the repository.
2. Create a new branch: `git checkout -b your-feature-branch`.
3. Make your changes and commit them: `git commit -m 'Add some feature'`.
4. Push to the branch: `git push origin your-feature-branch`.
5. Submit a pull request.

## Contact

For any questions or suggestions, please open an issue or contact [Saracomethstein](https://github.com/Saracomethstein).
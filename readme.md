# Image Server with Golang and Ngrok

This project demonstrates how to create a simple image server using Golang and expose it to the internet using Ngrok. The server serves images from a local directory and allows access via Ngrok-generated URLs.

## Prerequisites

Before getting started, make sure you have the following installed:

- [Golang](https://golang.org/dl/): The Go programming language.
- [Ngrok](https://ngrok.com/download): A tool to expose your local server to the internet.

## Getting Started

- Clone the repository

    ```
    git clone https://github.com/Sahil-4555/go-ngrok.git
    ```

- Place images

    Place your image files in the "images" directory. Ensure your images have extensions like .jpg, .jpeg, or .png.

- Run the server

    ```
    go run main.go
    ```
    This will start the Golang server, and it will be accessible at http://localhost:8080.

- Expose with Ngrok
    In a separate terminal, run Ngrok to expose your local server:
    ```
    ngrok http 8080
    ```
    Ngrok will generate public URLs (e.g., https://abcd1234.ngrok.io) that you can use to access your server from anywhere.

- Access Images

    You can access your images by using the Ngrok-generated URL with the following structure:
    ```
    https://abcd1234.ngrok.io/images/your-image.jpg
    ```

## Support Images Types

    The server supports the following image types by default:

    - JPEG (.jpg, .jpeg)
    - PNG (.png)

    You can add more image types by extending the getContentType function in the code.

## Customization

    You can customize the server code to fit your needs. For example, you can change the port, modify the supported image types, or enhance security.


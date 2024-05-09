# Description

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of our last project, ascii-art.

# Authors

- Sayed Hussain Mahfoodh
- Ali Alfardan

# Usage

To run this project, follow these steps:

1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. Install any necessary dependencies.
4. Run the server using the appropriate command `go run .`
5. Open your web browser and go to the specified URL to access the web GUI.

# Implementation Details

1. The program starts by importing necessary modules from Go's standard library, such as "fmt" for formatting, "net/http" for handling HTTP requests, and "text/template" for generating HTML.

2. It defines a struct called `AsciiArt` to hold ASCII art data.

3. The `main` function simply calls another function called `launchServer`.

4. Inside `launchServer` function:
   - It sets up a file server to serve files from the current directory.
   - Registers a handler for the `/ascii-art` endpoint, which handles requests related to generating ASCII art.
   - Starts the server on the localhost (your own computer) on port 8080.

5. The `handler` function is where most of the action happens:
   - It checks if the request method is POST, which typically means a form submission.
   - If it's a POST request, it reads form data submitted by the user, specifically `ascii-input` (the text to be converted into ASCII art) and `art-type` (the type of ASCII art to generate).
   - It then processes the input data, generates the ASCII art, and prepares it to be displayed.
   - Next, it loads an HTML template from the file system (`templates/ascii-art.html`), and fills it with the generated ASCII art data.
   - Finally, it sends the filled-out HTML template as a response to the user's browser.

6. If the request method is not POST, meaning it's not a form submission, the server responds with an error saying the method is not allowed.

In essence, this code sets up a simple web server that serves files from the current directory and provides a way for users to input text and get back ASCII art.

# Docker
1- create an image

      docker build -t ascii-art-web .

2- run a container

      docker run -p 8080:8080 ascii-art-web
    
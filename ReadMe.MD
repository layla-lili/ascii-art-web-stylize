
# Ascii Art Generator
This is an ASCII art generator implemented using Go for the server-side logic and HTML, CSS, and JavaScript for the client-side interface. It allows users to input text and select a banner style to generate corresponding ASCII art.


## Authors

- [@labdulla](www.https://learn.reboot01.com/git/labdulla)
- [@malkhoza](www.https://learn.reboot01.com/git/@malkhoza)
- [@famohamed]((www.https://learn.reboot01.com/git/@famohamed))




## Usage
To run the ASCII art generator, follow these steps:

#### 1. Clone the repository:
``` bash
git clone https://learn.reboot01.com/git/labdulla/ascii-art-web.git

```
#### 2. Navigate to the project directory:
``` bash
cd ascii-art-web

```
#### 3. Navigate to the main folder:
``` bash
cd main

```
#### 4. Build and run the Go server:
``` bash
go run .
```
#### 5. Open your web browser and visit http://localhost:8081 to access the ASCII art generator.

## Implementation Details
The ASCII art generator is implemented using the following components and algorithms:

* **Server-side:** The server-side logic is built with Go. It utilizes the net/http package to handle HTTP requests and responses. The server renders HTML templates using the html/template package, which allows dynamic generation of HTML content with data from the server.

* **Client-side:** The client-side interface is built with HTML, CSS, and JavaScript. The HTML structure defines the input fields, select options, and buttons. CSS is used for styling the user interface, providing an appealing visual experience. JavaScript is used to handle user input and make asynchronous requests to the server to generate ASCII art dynamically.

## Algorithm
The ASCII art generation algorithm follows these steps: 

    1. Receive user input, including the desired text and banner style, from the client-side interface.
    2. Validate the user input for any special characters or language limitations.
    3. Generate ASCII art based on the selected banner style and user-provided text.
    4. Return the generated ASCII art to the client-side interface for display.

## Error Pages
The application includes the following error pages:

* **400.html:** This page is shown when the server receives an invalid request or the user input contains non-ASCII characters. It informs the user about the error and provides instructions on how to correct it.

* **404.html:** This page is displayed when the user tries to access a non-existent URL. It indicates that the requested page cannot be found and provides suggestions or a navigation menu to help the user find their way.

* **500.html:** This page is shown when the server encounters an internal error during the execution of the application. It notifies the user about the error and may include a small hangman game 🎮😆 to keep the user engaged while the issue is being resolved.




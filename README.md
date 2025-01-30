# BashBrain 🤖💻

**BashBrain** is a terminal command executor written in Go, powered by **Gemini** (a generative AI model). The tool allows you to execute terminal commands using natural language prompts and get automated results. It makes interacting with the terminal more intuitive by leveraging the Gemini API, offering a seamless experience to perform common tasks with ease.

## Features 🌟

- **Natural Language Command Execution**: Use natural language prompts to execute terminal commands. 📝➡️💻
- **Powered by Gemini API**: Currently uses the Gemini API to generate command responses. ⚡
- **Customizable & Extensible**: Easily add new commands or modify existing ones for your use case. 🔧
- **Environment Configuration**: Requires a Gemini API key, which must be set in the `.env` file to function. 🔑

## Prerequisites 🛠️

To use this project, ensure that you have the following installed:

- **Go (1.18 or newer)**: To run and build the project. 🏗️
- **Gemini API Key**: You must have an API key from Gemini to make requests. 🔑

## Setup 🏁

Follow these steps to set up and use **BashBrain**:

### 1. Clone the Repository 📥

Clone the project to your local machine using the following command:

```bash
git clone git@github.com:ItsArnavSh/BashBrain.git
cd BashBrain
```

### 2. Set Up the `.env` File 📝

In the root directory of the project, create a `.env` file to store your Gemini API key.

```bash
touch .env
```

Inside the `.env` file, add the following line:

```bash
GEMINI_KEY=your_gemini_api_key_here
```

Replace `your_gemini_api_key_here` with your actual Gemini API key. 🔑

### 3. Install Dependencies 📦

Install the necessary Go dependencies using the following command:

```bash
go get
```

### 4. Run the Project 🚀

Run the project by executing the following command:

```bash
go run .
```

Once the server is running, you can start interacting with the terminal through the Gemini-powered API. 💬

### 5. Usage 🔍

The Gemini model will process the request, generate the appropriate command or response, and display the result in the terminal. 🤖

### Example Command 🧑‍💻:

```bash
go run .
```

Prompt:

```
Create a directory named "Test" and list its contents
```

Response:

```
Command executed: mkdir Test && ls Test
Result: Directory created and contents listed.
```

## API Details 🔧

**BashBrain** interacts with the Gemini API using a simple REST API request to generate content. The project currently uses **Gemini** to handle command execution in the terminal. 🌐


## Contributing 🤝

Feel free to open issues or submit pull requests to contribute to **BashBrain**. Contributions are welcome, and we appreciate any improvements or new features that enhance the project! 🌱

### Steps to Contribute:

1. Fork the repository. 🍴
2. Create a new branch for your feature or fix. 🌿
3. Make your changes and test them thoroughly. 🧪
4. Submit a pull request with a description of the changes. 📤

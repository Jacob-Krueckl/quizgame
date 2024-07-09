# Quiz Game

This is a simple quiz game written in Go. It allows users to answer multiple-choice questions and keeps track of their score.

## Installation

1. Clone the repository:

    ```shell
    git clone https://github.com/Jacob-Krueckl/quizgame.git
    ```

2. Navigate to the project directory:

    ```shell
    cd quizgame
    ```

3. Build the executable:

    ```shell
    go build
    ```

4. Run the game:

    ```shell
    ./quizgame
    ```

## Usage

- The game will present a series of multiple-choice questions.
- Enter the number corresponding to your answer and press Enter.
- Your score will be displayed at the end of the quiz.

### Custom Questions

You can provide your own questions by creating a CSV file with the following format:

```csv
question,answer,choice1,choice2,choice3,choice4
```

- `question` is the text of the question.
- `answer` is the number corresponding to the correct answer.
- `choice1`, `choice2`, `choice3`, and `choice4` are the possible answers.

The program is looking for a CSV file named `problems.csv` in the same directory as the executable.

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvements, please open an issue or submit a pull request.

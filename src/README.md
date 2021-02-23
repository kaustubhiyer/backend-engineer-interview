# Code Structure

## Usage Instructions
    To use, you need datasets similar to the ones provided in the data directory. Navigate to the src directory and run the following:

        ```
        > go build -o output
        > ./output -dir=../data
        ```
    
    The files in data need to be `themes.json` and `f.json` for the **themes** and **feedback** respectively.

## Code workflow
    The design basically follows the following procedure:
        - Get data files and pull data
        - Give the user options to navigate as well as sort

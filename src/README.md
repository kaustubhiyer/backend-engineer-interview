# Code Structure

## Usage Instructions
    To use, you need datasets similar to the ones provided in the data directory. Navigate to the src directory and run the following:

        ```
        > go build -o output
        > ./output -dir=<dir/to/data>
        ```
    
    The files in <dir/to/data> need to be `themes.json` and `f.json` for the **themes** and **feedback** respectively.

## Main Topics Screen
    The main topics screen contains a list of topics with ids. The interface gives you options to enter as text one of the following:
        1. <ID> -  This expands the topic with the corresponding id and shows its underlying themes
        2. <ID>F-  This switches to the feedback view for that topic
        3. exit -  Allows the user to exit the program

## Feedback Screen
    The feedback screen contains a list of reviews, paginated with page size 5. The interface gives options to:
        1. sort:<sort type> - This option allows the user to sort the reviews by: date(default), highest score, lowest score.
        2. pg:<page number> - This allows the user to move to a page of his choice
        3. return           - Allows the user to return to the topics screen
        4. exit             - Exits the program

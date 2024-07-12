# ASCII ART WEB

## Description
This is a simple web-based application developed using Go and HTML. It takes input text from the user then converts it to the appropriate graphical representation using the chosen ascii art banner file (or font). The user interacts with it through a graphical interface.

## Usage
1. **Clone Repository**: Clone this repository to your local machine.
2. **Install go**: Ensure you have Go installed on your machine.
3. **Run Program**: To get the server up and running, navigate to the directory where you've cloned the repo then use the following command:
    ```bash
    go run . 
    ```
4.  **Launch GUI**: Open your web browser and on the address bar type this:
    ```bash
    http:localhost:8080 
    ```
If the GUI doesn't open check your firewall settings and grant access to the program.

5. **Choose banner** : Select your banner of choice in the list provided.

6. **Input text**: In the text area labeled input text, type the text that you would like to be converted.
7. **Execute**: Click on the 'produce art' button. The program will execute and the results returned. If for example your input text is 'hello' and your banner of choice is 'standard' then the output will be:
    ```bash
     _              _   _          
    | |            | | | |         
    | |__     ___  | | | |   ___   
    |  _ \   / _ \ | | | |  / _ \  
    | | | | |  __/ | | | | | (_) | 
    |_| |_|  \___| |_| |_|  \___/  
                                
                                                          
    ``` 
## Files
Within the repo there are 2 directories and other files.
1. **main**

    This is the entry point of the program. The file contains various functions for handling different requests. It also sets up the server. 
2. **ascii-art**

    This directory contains the banner files and functions used for converting text to ascii-art representation. 
-  **art_map_creator**: Creates a map of ascii art characters from a chosen banner file.
-  **get_banner_file**: Gets the appropriate banner file that matches the choice. Example; if choice is 'standard' the function returns 'standard.txt'
-  **file_downloader**: If the selected banner file is missing, the function downloads it.
-  **art_maker**: Converts the text and returns the output.
-  **integrity_checker**: Ensures that the banner files are available and in original condition.
-  **tab_backaspace**: Handles instances of '\b' and '\t' present in the user input text.

3. **templates**

This contains various html templates used by the GUI.
-  **test.html**:  This is the home page.
-  **result.html**: This is where the output is displayed.
-  **about.html**:  This page contains a brief description of the web-app.

## Contributors

-   **[Hezborn Shikuku](https://github.com/Mania124/ascii-art-web)**
-   **[Rogers Ragwel](https://github.com/oragwelr/ascii-art-web)**
-   **[Rabin Otieno](https://github.com/Rabinnnn/ascii-art-web)**
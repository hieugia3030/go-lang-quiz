# go-lang-quiz
A simple GoLang console application to show quiz from a csv file and take answer from console input

# How to use
- Download all files in this github respiratory
- Run quiz.exe file with terminal (Command Prompt in Windows, Terminal in MacOS)

# How to run a exe file in Windows
- Type in cmd in search section - near the Window icon in the bottom left of the screen
- Click on Command Prompt
- Go to the folder containing the exe file. (Use cd, to go to D:/ just type in D: and enter)
- Execute the exe file by: quiz -csv="problems.csv"
- problems.csv file is the default file

# How to use a new csv file
- Create a new csv file with the format:
  question1, ans1\n
  question2, ans2\n
  ...

  For example:
  1+1,2
  2+2,4
- To use this file, run quiz -csv="yourFileName.csv"

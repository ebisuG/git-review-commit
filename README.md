# git-review-commit
## Overview
As a non-native English speaker, I'm often unsure whether my commit messages are appropriate. Pasting a draft commit message into a chat with AI is a little inconvenient because I usually work with the Git CLI during development, so I have to switch to a browser.

Complete "Git Commit Message Review" is the purpose of this tool.

I know that these days, the need to write code manually is decreasing. However, I think it is good practice to use AI tools under my control and use them to amplify my abilities. That's why I developed this tool.


## How to Use
1. Download the executable file from **GitHub Releases**.
2. Place the executable file in your `PATH`.
3. Rename the file as you like, following the `git-<command>` format.
4. Set your API key and model name in `config.yaml`. Place the file next to the executable.
5. Run `git <command>` in your local Git repository. Git will automatically find the command using the name you specified in step 3.

See : https://git-scm.com/docs/git#Documentation/git.txt-PATH
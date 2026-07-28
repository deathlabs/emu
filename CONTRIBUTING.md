# Contributing

## Contribution Workflow
This document describes a general workflow that a developer contributing to EMU will follow. This is the basic template for creating and adding code to the EMU project.

### Preamble
While the following sections describe at a high-level the process an EMU developer can expect to go through, it assumes that you the developer are already tasked to do the work. However, the EMU team envisions this project be open-source for the community, and we welcome input and contribution from members of the community using EMU.

The core EMU team may not have the capacity to take on a particular feature, or may not be aware of the impacts of a bug. But, the community using EMU may. 

If you want a feature or find a bug and know how to solve it (or would like to learn!), please open a branch in our repository. Our team will gladly help you get started if it is your first time. Either way, please also add the feature or bug to our teams page so we can better track it and provide the community with updates.

### Code Lifecycle
A new branch should be created each time new code is written regardless of what the code is. Make sure you are on the relevant branch locally (`main` or the development branch) when you create the new branch. You should also run either `git pull` or `git fetch` first to get the latest updates prior to branching.

Use `git checkout` to create a new branch off of the current branch. Make sure you use the proper naming conventions.

Write and commit your code. You should commit often and provide descriptive commit messages that explain what was changed in the commit.

Once your changes are made and committed, ensure you follow all style requirements, pass all qa and tests, and have a successful pipeline run. At that point you can open a merge request for your branch. **Make sure you set up the merge request to merge to the correct branch.**

Wait for a reviewer to review your code. You may have to further iterate depending on the results of the code review. This is done to ensure functionality and long-term usability of EMU for both users and developers. When the reviewer signs off on your code and the CI pipeline passes, you will be able to merge your code into the selected branch.

For the most part, you will now have your changes in the development branch. You can use the new feature you created right away, but it will not be generally available to the community until that development branch is merged into the main branch as a new release.

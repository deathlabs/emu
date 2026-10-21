# Git Standards
This document covers specific standards and practices followed by the EMU team for git. 

## Git
Git is used for code management and code versioning. All EMU developers should be familiar and comfortable with `git` basics.

### Git Best Practices
For a more in depth git best practices discussion, see this [article](https://about.gitlab.com/topics/version-control/version-control-best-practices/). But, as a quick overview: On the EMU team, the `main` branch and the next release branch are locked which means you cannot commit changes directly on the branch. Instead the proper procedure is to make a new local branch and do your development (and make any changes you have) to the local branch. As you make changes, you should commit those changes. All commits should include a descriptive message that describes what changed in the commit. You should try to change one thing (accomplish one task) per commit. This makes it easier to do code reviews and fix problems. Branches should also be focused on fixing a single issue, and should be deleted once the branch gets merged with the `main` branch. This helps prevent branches from becoming stale and outdated. All code will be reviewed prior to being merged into the `main` or release branch. Be receptive to changes and comments to help ensure the final product is cohesive, readble, and bug-free.

### Useful Git commands
Below are a few git commands (not-exhaustive) that may be helpful. Get the latest update of the branch (useful to get the most recent version of `main` branch)
```bash
git pull
```

If you are changing anything in the cloned repo, it is best practice to do it on your own branch. To create a new local branch, enter the command below.
```bash
git checkout -b <insert name of new branch>
```

To check which branch you are on, enter the command below.
```bash
git branch
```

To delete a branch, enter the command below.
```bash
git branch -d  <insert name of branch to delete>
```

Once you make changes on your local branch, you may want these to propagate up to the remote branch and eventually get merged into the main branch. To see local changes, enter the command below.

```bash
git status
```

To add changes to the eventual commit, enter the command below.
```bash
# add all changes
git add --all
# add all changes in the current directory
git add .
# add a specific file
git add <insert file name>
```

To commit the changes (must be done before push) and add a message to the commit, enter the command below.
```bash
git commit -m <insert your notes about the commit about what you did>
```

To push the changes to remote, enter the command below.
```bash
git push
```

Note that the first time you push a branch and the changes on that branch you will need to use the command, enter the command below.
```bash
git push --set-upstream origin <your branch name>
```

## Branches
Branches are the method in git to allow a developer to work on a part of the code without causing problems on the core branches or with other developers. A branch gives you a new copy of the current state of the branch you are on. This new branch is now a sandbox where a developer can make changes and even break the code without impacting anyone else on the project.

### Branch Semantics
EMU developers perform the majority of their work on "feature branches" - these branches are not restricted to new features, however (see [branch naming conventions](#branch-naming-conventions) below). From there, updates to EMU progress following EMU conventions:
- Feature branches generally target the development branch. This branch contains source for the next major release for EMU. It is the first place where work from multiple development streams gets integrated. Like the `main` branch, the development branch will be locked and require code reviews for merges.
- Unit test, technical debt, and documenation branches also generally will target the development branch.
- Bug and hot fixes as well as refactor branches will likely target both teh development and `main` branch. For the `main` branch, these will update the release version. See [Versions](#versions) and [Releases](#releases) below for more information. 
- When a development branch is ready for release, it will be merged with the `main` branch. A new development branch will be made (see [Releases](#releases)) and the process will continue.

In practice, there are a few ways to handle a fix that will need to go to multiple branches. 
- The simple but most repetitive option would be to create two branches, one of the development branch and one off of the `main` branch. Then, write the same code to address the bug in both and complete two merge requests one for the development and one for the `main` branch.
- Alternatively, you can find the common ancestor commit between the `main` and development and branch off of that point. Then, make the change and merge back into both branches. This will require you to not delete your branch when you complete the first merge.
- Finally, there is the `git-cherry-pick` option. You make one branch, fix the bug (ideally in a single commit), then cherry pick that one commit into the relevant branches

The second two options may be a bit scary for new comers as they involve some more advanced git workflows, so choose what you as a develper are most comfortable with.

### Branch Naming Conventions
Git branches should adhere to the following naming conventions in EMU:
- Branches that introduce new features are named `feature/<description>`
- Branches that introduce new tests (and only new tests) are named `test/<description>`
- Branches that address a change in the eMASS API that impact existing code (and require updates) are named `refactor/<description>`
- Branches that address technical debt (and only address technical debt) are named `debt/<description>`
- Branches that introduce bug fixes (and only bug fixes) are named `bugfix/<description>`
- Branches that introduce bug fixes with short suspense are named `hotfix/<description>`
- Branches that update documentation (and only update documentation) are named `docs/<description>`

If you do not follow these conventions you will be asked. This is to help organize branches and prioritize for reviews and merges. The only exceptions to these branch naming standards are the `main` branch and the current development branch which will be named `v<next major EMU version>` such as `v2`.

## Versions
Versions help the EMU team track changes overtime and better manage releases.

### Version Semantic
We follow common versioning practices for EMU and define our versions in the 0.0.0 semantic where the first integer signifies the major version, the second signifies the minor version, and the third signifies a patch. Changes to the EMU code base on the `main` branch will increment one of the integers depending on the type of change as described in teh next [section](#version-conventions). Version numbers are only relevant to the `main` branch as that handles the releases for EMU. All other branches including the development branch do not need to worry about version numbers. EMU will not guarantee major versions will have backwards compatibility. Minor changes and patches should, and it should be considered a bug if this is not the case.

### Version Conventions
(TODO) We initially stated that we want to use the Semantic Release based on the code and structure described by CW3 Matt Lindholm in the I&P repository. However, this should have a follow up discussion because it will require EMU to bring in a collection of additional dependencies in JavaScript that are not required for functionality. Abby thinks it is reasonable to instead write the basic structure/requirements instead using `splinter` in just the pipeline to update versions/releases. The title of the merge request can indicate the version update which may be a variable to store in the CI/CD variables. 

The prefix format format should be added to the title of merge requests to the `main` branch:
 - fix: A bug fix or hotfix. This will increment the patch version number: 0.0.1 -> 0.0.2
 - feat: A new feature. This will increment the minor version number: 0.0.1 -> 0.1.0
    - This should be used when an eMASS update occurs that causes changes in the API
 - BREAKING CHANGE: A breaking change. This will increment the major version number: 0.0.1 -> 1.0.0
 - release: A major release that will also increment the major version number: 0.0.1 -> 1.0.0 
 - chore: A change that doesn't affect the code (e.g. updating dependencies). This will not increment the version number. If no prefix is provided, it will default to chore.

The prefix will be parsed and verison updated according to the prefix passed. Note, all version resets return the version number to `0` 

## Releases
Releases only happen on the `main` branch for the current version of `EMU`. Generally, only hot fixes (`fix`) will occur outside of release, but new features that are important to functionality or updates required for new eMASS api releases will be `feat` and update the minor versions. The current development branch for the next major version of EMU does not need to manage versioning.

### Release Branches
The `main` branch contains the code for the current release of EMU. It should have minimal merge requests and generally be limited to bug fixes. Most features and significant changes will occur on a different branch. Also refered to as the development branch elsewhere in this document, this branch will accept most of the merge requests made in the EMU project. It is where development is expected to occur. In EMU the development branch will be labelled as the next major version of EMU. For example, `v1` to indicate that the next release of EMU will be `version 1` and will result in `1.0.0` as the version number when this branch is merged into the main branch.

For most EMU development, developers will branch off of the current release branch. This will be the only other branch besides `main` that will be locked. It will also be the only other branch on the EMU project that will not follow the branch naming conventions explained in [Branches](#branches). The EMU team has a roadmap and release document that indicates what features are required or a part of the next few major versions of EMU. Further out versions of EMU will still be pliable as to what features are added, and depending on the demand signal, features can be shifted or added. Once all requirements are met for the next major version, the EMU team will finalize and release the new major version by merging the development branch into the `main` branch in the EMU repository.  

### Release Management
The EMU team has a roadmap and release document that indicates what features are required or a part of the next few major versions of EMU. Further out versions of EMU will still be pliable as to what features are added, and depending on the demand signal, features can be shifted or added. Once all requirements are met for the next major version, the EMU team will finalize and release the new major version by merging the development branch into the `main` branch in the EMU repository.  

As the finalizing of the release occurs, we will try to close out all outstanding active branches in the repository. We will also request minimizing the number of new branches and features that are started in this window. We will do our best to minimize this window to prevent stalling development. Once the development branch is merged into main, we will create a new, locked branch off of the main branch. This will become the new development branch for the next release. Once the branch is created, we will again open up EMU to new features. When a new major version is released, we will still have older versions accessible in our package and container registries, but we will no longer support them or provide updates to those versions.

### Accessing the Latest Release
With every merge to `main` we will run additional pipeline jobs that will push both the EMU binary and the EMU container to their respective package and container repositories on in the EMU Gitlab repository. This is where users should go to get the latest version of EMU. You will not need developer access to pull the binary or the container, but you will need access to R2D2 or wherever our Gitlab instance is hosted. We are also exploring how we can release both the binary and container in other locations to increase ease of access. As a developer, you can also clone the EMU repository and build either the binary or image locally off of the development branch. This will give you access to new features not yet generally available, but you may run into more bugs or errors.

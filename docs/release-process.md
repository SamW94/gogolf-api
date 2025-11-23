# 🏌🏼‍♂️ gogolf-api release process 

1. Pull the latest changes from the `dev` branch. 
		```
		git switch dev
		git pull
		```
2. Checkout the *release branch*. Use [semantic versioning](https://semver.org) based on the latest tag to decide what the new version should be. Name the branch `r-<major-version>-<minor-version>-<patch-version>` like this.

		git checkout -b r-0-0-1

3. Create or update the `CHANGELOG.md` file according to the [Keep a Changelog Style](https://keepachangelog.com/en/1.1.0/)

		git add CHANGELOG.md
		git commit -m "chore: release v0.0.1"

4. Push your changes to the remote repository on GitHub.

		git push --set-upstream origin r-0-0-1

5. Open a PR and review the changes. Paste the relevant section of the `CHANGELOG.md` file into the PR body instead of the PR template that's already there.

6. Merge the PR. **Do not squash the commits. Do not delete the release branch yet.**

7. Switch to the `main` branch and create a tag that matches the release version you decided to use above. 

		git switch main
		git tag v0.0.1
		git push origin --tag

8. Ensure the `Build API Documentation and Deploy to GitHub Pages` and `Build and push a Docker image (release)` jobs have run successfully in the GitHub Actions tab.

9. Create a second PR to merge any changes you made to the release branch/`CHANGELOG.md` file back to the `dev` branch. Paste the relevant section of the `CHANGELOG.md` file into the PR body instead of the PR template that's already there.

10. Merge the PR. You can squash the commits. You can now delete the release branch. 
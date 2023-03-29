# Personal Bot

This repository is an application that uses OpenAI GPT-3 to classify payments mails from your Gmail account and store them in a PostgreSQL database. The application is protected using Google's OAuth2.0 Service, ensuring the security of your information. You can use this bot for personal use, but can not be used for commercial purposes.

## Technologies Used

[![GitHub go.mod Go version of a Go module](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
[![Docker](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)
[![Postgres](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)

The following technologies were used to develop this application:

- Go 1.20
- Docker
- PostgreSQL

## Current State

The application is currently able to read your payment mails from your Gmail account, classify them using OpenAI GPT-3, and store them in a PostgreSQL database. It is also protected using Google's OAuth2.0 Service, ensuring that your data is kept safe and secure.

## What's Coming Next

In the near future, I plan to add some health features. It is already planned but don't want to make it public yet ;\)

Stay tuned for more updates!

## Setting Up Personal Bot with Docker

This tutorial will guide you through setting up the Payment Mail Classifier with OpenAI GPT-3 using Docker. This tutorial assumes that you have Docker and Docker Compose installed on your machine.

### Configuration Files Setup

- Clone the repository to your local machine.

`git clone git@github.com:JuanVF/personal_bot.git`

- Create a file called Dockerfile in the db folder, and set your Postgres password in this file. Do not change the username.

```docker
ENV POSTGRES_PASSWORD your_fancy_password
```

- Create a `local.yaml` file in the `common` folder, and replace all the variables with your own OpenAI API data, Google Console Data, your Postgres username and password, and in the container environment, set the Google Redirect to your prod URL if you are aiming for that.

### Docker Image and Docker Compose Setup

- Run the `update_bot_image` and `update_bot_db_image` targets in the Makefile and pass the `VERSION` parameter to create the Docker images.

```bash
make update_bot_image VERSION=1.0.0
make update_bot_db_image VERSION=1.0.0
```

- Modify the docker-compose.yaml file with the versions you just set.
- Start the Docker Compose in your prod environment.

And that's it! You should now have the Personal Bot up and running in your prod environment.

## Useful Resources

Here are some useful resources related to the Payment Mail Classifier with OpenAI GPT-3:

- Google OAuth2.0 Setup: https://developers.google.com/identity/protocols/oauth2/web-server
- OpenAI API: https://platform.openai.com/docs/api-reference

For the OpenAI API, you will need to create a custom model and train it yourself to use it with this project. You can find detailed instructions on how to do this in the OpenAI API documentation. Also, you can obtain your API keys to use with this project in your Open AI Platform Account.

## How to Contribute

We welcome contributions from the community to help improve the `Personal Bot`. Here are the steps to follow if you would like to contribute:

- Check the `issues` section to see if there are any open issues that you could work on. If there is an issue that you would like to work on, please leave a comment indicating that you are interested in working on it.

- If there are no open issues that you want to work on, you can create a new issue to propose a new feature or suggest an improvement to the existing codebase.

- If you are a collaborator, you can create a new branch to work on your feature, push your changes to your branch, and then create a Pull Request (PR) to merge your changes into the main branch. Please make sure to follow the `coding standards` and include unit tests for any new features or changes.

- If you are not a collaborator, you can fork the repository, create a new branch to work on your feature, push your changes to your branch in your forked repository, and then create a Pull Request to merge your changes into the main branch of the original repository. Please make sure to follow the `coding standards` and include unit tests for any new features or changes.

We appreciate any contributions to the project and will review PRs and issues as soon as possible.

## Coding Standards

To ensure consistency and maintainability of the Personal Bot, we follow the following coding standards:

### Go Version

The codebase is developed using `Go 1.20`.

### Code Formatting

All code should be formatted using gofmt.

### Variable Naming

Variable names should be descriptive and in camelCase. For example: `paymentAmount`.

### Function Naming

Function names should be descriptive and in camelCase, with the first letter of each word capitalized. For example: `ParsePaymentEmail`.

### Comments

Code should be well-documented with clear and concise comments. Comments should be written in English.

### Error Handling

All functions should return an error, and errors should be checked and handled properly.

### Unit Tests

All code changes should include unit tests to ensure that the changes work as expected and do not introduce any regressions. Unit tests should cover both positive and negative scenarios.

By following these coding standards, we can ensure that the codebase is consistent, easy to read and maintain, and that changes can be made safely without introducing bugs or breaking existing functionality.

## Tags

> `Go` `Docker` `PostgreSQL` `OAuth2.0` `OpenAI GPT-3`

## License

This repo uses (CC BY-NC)

[![Licencia Creative Commons](https://i.creativecommons.org/l/by-nc/4.0/88x31.png)](http://creativecommons.org/licenses/by-nc/4.0/)

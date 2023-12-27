# Distributed Build and Continuous Integration MVP

![gopher (1)](https://github.com/suyash0x/dbci/assets/121722187/69211596-6d15-44a1-8b6c-da92799cf5c5)


## Introduction

This project is a hands-on exploration of Go programming language concepts and practices, specifically tailored for learning purposes. The primary goal is to delve into automating and optimizing the software development process through continuous integration and distributed build strategies. DBCI involves automatically building code changes upon each commit, ensuring streamlined development workflow. In this MVP, the publisher-subscriber model forms the core, allowing build servers to connect to a central publisher for triggered build processes. Although the MVP currently simulates triggers at intervals and logs build processes, future iterations will integrate with real Version Control Systems (VCS) and enhance the build simulation, providing a solid foundation for scalable and efficient continuous integration practices. Happy coding!

## Setup instructions (Refer Makefile for commands)

1. clone the repo
2. execute the build command `make all`
3. run publisher `make publisher`
4. run build server `make buildServer`



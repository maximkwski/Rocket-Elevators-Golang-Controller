# Rocket-Elevators-Golang-Controller
This is the template to use for the golang commercial controller. You will find the classes that should be used along with some methods described in the requirements. The necessary files to run some tests are also included. 
Hello! Here you will find Elevator Controller Project that can be used for Commercial type Buildings. In Current project we create a 4 elevator columns, that hold 5 elevators per column. Building has 60 floors and 6 basement floors. Each column responsible for certain range of floors. Depends on scenario, the proper elevator will be sent. 
The project folder has a test file with multiple scenarios to see how code works with different inputs. 
The Classes are separated in diferent files and named accordingly for convenience. 

### Installation

With golang installed on your computer, all you need to do is initialize the module:

`go mod init Rocket-Elevators-Commercial-Controller`

The code to run the scenarios is included, and can be executed with:

`go run . <SCENARIO-NUMBER>`

### Running the tests

To launch the tests:

`go test`

With a fully completed project, you should get an output like:

![Screenshot from 2021-06-15 15-25-10](https://user-images.githubusercontent.com/28630658/122111573-e6ea7380-cded-11eb-95e3-95e0096a1b3a.png)

The test and scenarios files can be left in your final project. The grader will run tests similar to the ones provided.

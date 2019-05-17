# GO-JEK-CAR

Need to design a parking area that issues a ticket to driver printing registration number and color. The objective is to provide the nearest parking slot to the driver and in case no parking space is available then retrun a proper message to the drivers. Along with this provide additional funtionality to query following features.

● Registration numbers of all cars of a particular colour.
● Slot number in which a car with a given registration number is parked.
● Slot numbers of all slots where a car of a particular colour is parked.


It is necessary that the code should be placed in `$GOPATH/src/` folder.

#### Run

`./bin/setup` runs the tests, install dependencies and/or compile the code and then run unit test cases and outputs the test results to `output.log` file.


```
./bin/setup
```

Following command is to execute the input file which will output result to STDOUT.

```
./bin/parking_lot input_file.txt
```

#### Tests

Tests are run using below command

```
go test ./...
```
Written in Go
DBMS in Postgres
####
RESTful web application that satisfies the below user stories.

1. Let you download all the activities in the masterplan in a csv or excel file.
2. Let you view all the details like start and end dates of the activities in the csv file.
3. Let you download a list of activities sorted by WBS number as a csv or excel file.
4. Let you download a sorted csv or excel file based on the start date of activities followed by the Work Breakdown Structure number in the masterplan.


## Steps to Build the project

Run `go get` command to download all dependencies.
Run `go build` to create the binary.


### Project Setup Method

1) setup the Golang setup first and then clone the whole project repo.
2) setup a Postgres DB named "ConstructionPlan" and select the DB tar provided in the repo to restore the database.
3) download all the dependencies of the golang whichever has been used. (mentioned in the import section)
4) good to run the project.

### Swagger documentation
Fire the below command to Generate a spec from source
swagger generate spec -o ./swagger.json



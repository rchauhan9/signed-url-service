# Signed URL Service (AWS)

This repository generates signed urls to securely retrieve objects from a 
chosen S3 bucket. The service runs a gin-gonic web server so that a signed
URL can be requested with a simple HTTP GET request.

## Setup
* Ensure you have an AWS Access Key and Secret Access Key set up in the 
below file. It is possible to have multiple credentials set up in this file,
see [Specifying Profiles](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html)
```
~/.aws/credentials
``` 

* If you have multiple aws credentials stored in this file, set an environment
variable to specify which credentials to read from. The variable required is 
AWS_PROFILE.   
**Note: your AWS credentials must be from as user that has 
read-access to the S3 you plan to interact with**

* If using an IDE like GoLand you can set this variable in Run > Edit Configurations

* If you plan on running the app from the command line set your variable using
the command below, substituting `<example_profile>` for the profile of your choice.
```
export AWS_PROFILE=<example_profile>
```

## Running the app
The app can be run whilst specifying a port to expose of your choice.

##### From GoLand
* Head to Run > Edit Configurations and set program arguments to the below
(8081 can be substituted for a port of your choice).
```
-port 8081
```
Then click the Run button from the IDE.

##### From the Command Line
* Run the below cmd from the project's root directory. Again the port number 
can be changed to one you desire.
```
go run main.go -port 8081
```

## Retrieving the signed URL
The URL is retrieved via an HTTP GET request. If the server is now running 
locally, you can use curl as a quick proof of concept.
```
curl localhost:8081/<bucket_name>/<key>/<path>/<to>/<file on S3>
```
* The bucket name should be the text between `8081/` and the next `/` 
* Anything after `<bucket_name>/` is interpreted as the S3 key (filepath).
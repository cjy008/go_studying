## URL Shortener
The goal of this program is to the create an http.Handler that will determine if the path requested from the incoming web request should redirect the user to a new page, much like how bit.ly works.

For instance, if  there is a redirect setup for /dogs to https://www.somesite.com/a-story-about-dogs, the program will look for any incoming web requests with the path /dogs and redirect them to the specified URL.

Once you have that working, focus on parsing the YAML using the gopkg.in/yaml.v2 package. Note: You will need to go get this package if you don’t have it already.

Majority of the logic is handled in the `handler.go` file. The `main/main.go` file uses the package created by the `handler.go` to test the code and create a webserver in order to get an idea of what the program is doing.

The program relies on two types of mapping, a `MapHandler` and a `YAMLHandler` to create the mappings of the short URLs to the redirect locations.

### Build and Execute
```
# Create the local url_short package
go mod init url_short

# Execute the main go file
go run main/main.go
```

### Possible further enhancements
- Update the main/main.go source file to accept a YAML file as a flag and then load the YAML from a file rather than from a string.
- Build a JSONHandler that serves the same purpose, but reads from JSON data.
- Build a Handler that doesn’t read from a map but instead reads from a database. Whether you use BoltDB, SQL, or something else is entirely up to you.

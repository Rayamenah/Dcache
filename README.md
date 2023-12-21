# Database cache

This is a simple database cache implementation for caching user information.
Say we have a bunch of users information in our database and there are thousands of simultaeneous requests, implementing a caching mechanism allows us to cache the user information while simultaneously taking care of databse throughput

Example:
if within 1000 requests there are 100 unique user ids then there should be only a maximum of 100 requests into the database but all 1000 requests should get a response with a user data
##

A test file is avaliable to run the code 

run <code>go test</code>

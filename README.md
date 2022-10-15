Hello World!

This is a url shortener built using golang and redis.
Ideally, you should use a more traditional RDBMS for better scalability.

The primary objectives of this project are:

1. Build a simple URL shortener service that will accept a URL as an argument over a REST API and
return a shortened URL as a result.
3. If I again ask for the same URL, it should give me the same URL as it gave before instead of
generating a new one.
4. Put this application in a Docker image by writing a Dockerfile
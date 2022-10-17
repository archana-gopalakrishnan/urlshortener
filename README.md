Hello World!

This is a url shortener built using golang and redis.
Ideally, you should use a more traditional RDBMS for better scalability.

The primary objectives of this project are:

1. Build a simple URL shortener service that will accept a URL as an argument over a REST API and
return a shortened URL as a result.
3. If I again ask for the same URL, it should give me the same URL as it gave before instead of
generating a new one.
4. Put this application in a Docker image by writing a Dockerfile

To run the application:
1. Since this uses golang and redis, please make sure you have golang and redis installed on your system
2. Clone the directory locally and run: go run main.go. 
3. Following which you can test out if the program works by opening http://localhost:8080/ on your browser. You can run the curl command similar to the following one to test the generation of url. 

```
curl --request POST \
--data '{
    "long_url": "https://www.youtube.com/watch?v=ryp60Q397b4"
}' \
  http://localhost:8080/create-short-url

```

Additionally: You can try to paste the shorteneed url you get as output the above curl command in your browser to see that it gets directed back to the main url's website

4. To create docker image you need to first ensure you have Docker running locally. Then run docker-compose up -d. Following which you can check if containers are running using docker-compose ps.
After this you should be able to hit the same localhost:8080 again.
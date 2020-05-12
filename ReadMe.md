# FireGo

A little Golang API which allows to dynamically create JSON documents with a NoSQL Database perspective.

## What is it

Just an API running on port 5000, which allows you to dynamically create JSON documents in order to constitute kind of a NoSQL Database.

I wrote it to be quite similar to the database service that Google offers with Firebase.
It's not as sophiscated, for now.

## How do I use it

This system uses a logic of "resources", which are in fact separate JSON Documents.

To start the API, you can use :

``` go run main.go```


**or**


``` go build -o firego .``` and ```./firego```


**or**


```docker build -t firego .``` and ```docker run -d -p 5000:5000 firego```


### Read

In order to read from a resource, you can query this, by using the ```GET``` method :

``` curl localhost:5000/resource```

You can also read a specific ID from a resource.

``` curl localhost:5000/resource/id```

### Write

In order to insert an item in a resource, you can query this, by using the ```POST``` method :

``` curl -d '{"name":"t-shirt", "price":"8.80"}' -X POST localhost:5000/products```

In the body of your request, you can put the data to insert.
For now, values can only be strings.

But I'll fix this asap.

### Update 

In order to update an item in a resource you'll first need its ID, which is automatically generated at the insertion. Next you can query this, by using the ```PATCH``` method :

``` curl -d '{"name":"t-shirt", "price":"8.80"}' -X PATCH localhost:5000/products/productsID```

In the body of your request, you can put the data to insert.
For now, values can only be strings. 

But I'll fix this asap.

## Functionnalities Roadmap

- [x] F1 : Allow to insert all kind of values (bool, string, num ..)
- [ ] F2 : Deploy to Heroku
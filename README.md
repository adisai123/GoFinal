
***Part of Url***
https://abc.google.com:80/videopath?docid=234234234&hl=en#000hhsds

protocol = https
abc = sub domain
google.com = domain name
80 = port
videopath = path
? = query
docid=23..   = paramenters seperated by & and starts after ?
000hhsds = fragement


# ***cookies***: 
little file , server run on client , saved on clients machine , you can write unique id on to it of each id . if there is any cookies for specific domin , then it will return back to that domain whenever requested , by appended in the url.

#start mysql server : mysql.server start

# ***Context***: 
carries deadlines , cancelation signals , and other request scope values across API boundries and between processes. Incomming request should create context and outgoing calls to servers should accept a context. The chain of function calls between them must propogate the context, optionally replace it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue. When a context is canceled, all contexts derived from it are also canceled.

Context makes it possible to manage a chain of calls within the same call path by signaling context's Done channel.
context value scope is request scope

# ***Ajax***: 
Asynchronous Javascript and xml. In a nutshell, it is use of the XMLHTTPRequest object to communicate with server-side  scripts. It can send as well as receive information in a variety of formats , including JSON, XML, HTML and even text files. Ajax's most appealing characteristics, however its asynchronous nature, which means it can do all of this without having to refresh the page. This lets you update portions of a page based on user events

```javascript
var xhr = new XMLHttpRequest();
xhr.open('GET','XHTML.html',true);
xhr.onreadystatechange = function (){
    if(xhr.readState === XHMHttpRequest.DONE && xhr.status === 200){
        console.log(xhr.responseText)
    }
};
xhr.send();
```

#***Curl***:

```curl -X POST -H 'Content-Type: application/json' -d '{"Name":"aditya","Gender":"Male","Age":"29","Id":"123"}' http://localhost:8080/user/```

***Docker***
Docker file -> docker images -> docker container

Docker file must be named Dockerfile
Dockerfile always start with From : specify some image unless you're creating from scratch.

RUN executes commands in a new layer and creates a new image. E.g. its often used for installing software packages.

CMD sets default command and/or paramenters, which can be overwritten from command line when docker container runs.

ENTRYPOINT configures a container that will run as an executable.




# Some comment
```
FROM golang:1.8-onbuild
MAINTAINER youremail@gmail.com


docker build -t my-app .
docker run -d -p 80:80 my-app
07_push-pull
docker images
docker tag <image ID>  <docker hub username>/<image name>:<version label or tag>
docker login
docker push <docker hub username>/<image name>
docker images
docker rmi -f <image ID or image name>
docker --help
docker <COMMAND> --help
docker rmi --help
docker search <yourusername>
## run it this way if it's our go web app from previous step
docker run -d -p 80:80 <yourusername>/<app-name>
docker ps
docker stop <container id>
docker ps
docker images
08_aws-docker
sudo chmod 400 your.pem
ssh -i /path/to/[your].pem ec2-user@[public-DNS]
sudo yum update -y
sudo yum install -y docker
sudo service docker start
sudo usermod -a -G docker ec2-user
docker info
docker run -d -p 80:80 toddmcleod/golang-hello-world
docker ps
Use the IP address of your instance
```

To add user to docker group so that you can execute Docker commands without using sudo
```bash
sudo usermod -a -G docker ec2user
```
scopes : function , file , package , Universal scopes..




# ***Mongo*** -----
```
db.dropDatabase()
db.mycol.find()
db.mycol.remove({'title':'MongoDB Overview'})
db.COLLECTION_NAME.drop()
```
String − This is the most commonly used datatype to store the data. String in MongoDB must be UTF-8 valid.

Integer − This type is used to store a numerical value. Integer can be 32 bit or 64 bit depending upon your server.

Boolean − This type is used to store a boolean (true/ false) value.

Double − This type is used to store floating point values.

Min/ Max keys − This type is used to compare a value against the lowest and highest BSON elements.

Arrays − This type is used to store arrays or list or multiple values into one key.

Timestamp − ctimestamp. This can be handy for recording when a document has been modified or added.

Object − This datatype is used for embedded documents.

Null − This type is used to store a Null value.

Symbol − This datatype is used identically to a string; however, it's generally reserved for languages that use a specific symbol type.

Date − This datatype is used to store the current date or time in UNIX time format. You can specify your own date time by creating object of Date and passing day, month, year into it.

Object ID − This datatype is used to store the document’s ID.

Binary data − This datatype is used to store binary data.

Code − This datatype is used to store JavaScript code into the document.

Regular expression − This datatype is used to store regular expression.
databases -> collection -> documents

***create database*** : ```use <db_name>``` (either way to use it or to create and use it for the first time.)

***see current database*** : ```db```

***see all dbs***:  ```show dbs```

***View collections*** :   ```show collections ```

***insert document***:  
 ``` db.<collection_name>.insert([{}])```
 for single row just insert {}
(If collection does not exist, it is created.)


***View Documents***:
``` db.<collection_name>.find()```

### example
```db.cat.find().preety()```

***drop db***
```db.dropDatabase()```

***Query***
    
### find one
    ```db.cat.findOne()```

### find
    ```db.cat.find()``

### find specific
    ```db.cat.find({name:"sjdh"})```

### or
    ```db.cat.find({$or:[{name:"aditya"},{age:123}]})```

### and
    ```db.cat.find({$and:[{name:"aditya"},{age:{$lt:123}}]})```  //lt gt ,lte, gte , ne

### andor
    ```db.cat.find({$and:[{name:"aditya"},{age:"123"}]},{$or:[{name,"sai"}]})```

### regx
    ```db.cat.find({name:{$regex:'^M'}})```

### update 
    ``` db.<collection_name>.update(<select critera>,<update data> , <optional options>)```
    example
    ``` db.cat.update({name:"adita",_id:ObjectId("123123")},{$set:{age:12}})```
    ```db.cat.update({},{$set:{role:"abc"}},{multi:true})```    
    {multi:true} -- if not used only one record will be used.

### save // create or update record(if _id passed) 
    ```db.cat.save({"role":"villain","name":"Goldfinger","age":777})```

### remove 
    ```db.<collection_name>.remove(<selection criteria>)```
    ```db.customer.remove({role:"double-zero"}) or db.customer.remove({role:"double-zero"},1)```
    here 1 - to remove only one record but you can use only 1 and not other number.

### projection
    Retriving part of a document; only some of the fields.
    ```db.<collection name>.find(<selection critera>,<list of fields with toggle 0 or 1>)```
    ```db.customer.find({},{_id:0,name:1})```
    id is default on so marked it 0 to disable and only name will be fetched.

    ```db.custmer.find({},{_id:0,name:1,age:1})```
    it will fetch name and age

### limit
    ```db.<collection_name>.find(<selection critera>).limit(n)```

### sort
    ```db.<collection_name>.find(<selectioncritera>).sort(<field to sort on>:<1 for ascend, -1 descend>)```
    ```db.cat.find({}).limit(10).sort({name:1})```
### index
    concept of no sql index  is similar to sql index
    (used for quick search on regular data that are required to search)
   ``` db.createIndex({<field to index>:<1 ascend, -1 descend>})```

### see all the index:
    ```db.<collection_name>.getIndexes()```

### aggregate
    ```
    db.collection.count()
    db.collection.distinct()
    ```
    # count
    ```
    db.cat.find({name:"aditya"}).count()
    db.cat.count()
    ```

    #distinct
    ```
    db.cat.distinct("name")
    ```
   
   ```
   #aggregation:
   db.<collection name>.aggregate([{{match,sort}},{<group>}])
    ```
    db.order.aggregate([
        {$match:{status:"A"}},
        {$group:{
            _id:"$cust_id",
            total:{$sum:"$amount"}
        } }
    ])

todo: instead of mongo 
[]file based 
[]map based 
write to file (everytime modify map and call store which will write to file)
json.NewEncoder(f).Encode(&d)

read to file json.NewDecoder(f).Decode(&d)

refactor code model , controller , sessio (117)

### create regular user
while creating user to perticular database , you have to be there (using use keywoard).

```
use admin
db.createUser({
    user:"aditya",
    pwd:"password",
    roles:[{
        role:"readwrite",
        db:"store"
    }]
})
```
exit mongo & then start again 
```mongod```

```mongo -u "aditya" -p "password" --authenticationDatabase "admin"```

# see current user:

    ``` db.runCommand({connectionStatus:1})```

#drop user :
    ```db.dropUser("<user name>")```


#External links***:
https://golang.org/pkg/net/http/

third party mux:
https://godoc.org/github.com/julienschmidt/httprouter

https://github.com/GoesToEleven/golang-web-dev

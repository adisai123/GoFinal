https://golang.org/pkg/net/http/

third party mux:
https://godoc.org/github.com/julienschmidt/httprouter


https://abc.google.com:80/videopath?docid=234234234&hl=en#000hhsds

protocol = https
abc = sub domain
google.com = domain name
80 = port
videopath = path
? = query
docid=23..   = paramenters seperated by & and starts after ?
000hhsds = fragement


cookies: little file , server run on client , saved on clients machine , you can write unique id on to it of each id . if there is any cookies for specific domin , then it will return back to that domain whenever requested , by appended in the url.

start mysql server : mysql.server start

Context: carries deadlines , cancelation signals , and other request scope values across API boundries and between processes. Incomming request should create context and outgoing calls to servers should accept a context. The chain of function calls between them must propogate the context, optionally replace it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue. When a context is canceled, all contexts derived from it are also canceled.

Context makes it possible to manage a chain of calls within the same call path by signaling context's Done channel.
context value scope is request scope

Ajax: Asynchronous Javascript and xml. In a nutshell, it is use of the XMLHTTPRequest object to communicate with server-side  scripts. It can send as well as receive information in a variety of formats , including JSON, XML, HTML and even text files. Ajax's most appealing characteristics, however its asynchronous nature, which means it can do all of this without having to refresh the page. This lets you update portions of a page based on user events

var xhr = new XMLHttpRequest();
xhr.open('GET','XHTML.html',true);
xhr.onreadystatechange = function (){
    if(xhr.readState === XHMHttpRequest.DONE && xhr.status === 200){
        console.log(xhr.responseText)
    }
};
xhr.send();


curl -X POST -H 'Content-Type: application/json' -d '{"Name":"aditya","Gender":"Male","Age":"29","Id":"123"}' http://localhost:8080/user/


Mongo -----

MongoDB use DATABASE_NAME is used to create database
To check your currently selected database, use the command db

If you want to check your databases list, use the command show dbs.
Basic syntax of dropDatabase() command is as follows −
db.dropDatabase()


todo: instead of mongo 
file based 
map based 

write to file (everytime modify map and call store which will write to file)
json.NewEncoder(f).Encode(&d)

read to file json.NewDecoder(f).Decode(&d)

refactor code model , controller , sessio (117)



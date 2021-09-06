const http = require('http');

const hostname = '127.0.0.1';
const port = 3000;

const server = http.createServer((req, res) => {


    res.statusCode = 200;
    res.writeHead(200, {"Content-Type" : "text/plain; charset=UTF-8"});

    'use strict';

    const fs = require('fs');

    let rawdata = fs.readFileSync('data.json', "utf-8");
    let student = JSON.parse(rawdata);


    let studentLength = student.length




    for (let i = 0; i < studentLength ; i++){
        console.log(student[i])

        res.write(student[i].UUID.toString() + "\n", 'utf-8')
        res.write(student[i].EventName.toString()+ "\n", 'utf-8')
        res.write(student[i].EventTime.toString()+ "\n", "utf8")
        res.write(student[i].Floor.toString()+ "\n", "utf8")
        res.write(student[i].Room.toString()+ "\n", "utf8")

        if (student[i].Presentation.toString() == "false"){
            res.write("Cvičenie", "utf8")
        } else {
            res.write("Prednáška", "utf8")
        }

        res.write("\n------------------\n")

    }




    res.render('studentlist',{'studentlist' : result} );


    res.end("END OF VYPIS")
});


server.listen(port, hostname, () => {
    console.log(`Server running at http://${hostname}:${port}/`);
});
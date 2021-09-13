/*
    let data;
    let rawFile;

     */

let data = {
    "Floor"  :  "",
    "Room"   :  "",
    "EventName"   :  "",
    "EventTime"   :  "",
    "Presentation"      :  true,
    "DateStart"   :  "",
    "DateEnd"   :  "",
    "UUID"   :  "",

};
let rawFile = {
    "Floor"  :  "",
    "Room"   :  "",
    "EventName"   :  "",
    "EventTime"   :  "",
    "Presentation"      :  true,
    "DateStart"   :  "",
    "DateEnd"   :  "",
    "UUID"   :  "",

};

data.async = false;
rawFile.async = false;




function ReadTextFile(file, callback) {
    rawFile = new XMLHttpRequest();
    rawFile.overrideMimeType("application/json");
    rawFile.open("GET", file, false);
    rawFile.onreadystatechange = function() {
        if (rawFile.readyState === 4 && rawFile.status == "200") {
            callback(rawFile.responseText);
        }
    }
    rawFile.send(null);
}

function CreateTableFromJSON() {
    ReadTextFile("data.json", function(text) {
        data = JSON.parse(text);
        console.log(data);
    });

    // EXTRACT VALUE FOR HTML HEADER.
    var col = [];
    for (var i = 0; i < data.length; i++) {
        for (var key in data[i]) {
            if (col.indexOf(key) === -1) {
                col.push(key);
            }
        }
    }

    // CREATE DYNAMIC TABLE.
    var table = document.createElement("table");

    // CREATE HTML TABLE HEADER ROW USING THE EXTRACTED HEADERS ABOVE.

    var tr = table.insertRow(-1);                   // TABLE ROW.

    for (var i = 0; i < col.length-3 ; i++) {
        var th = document.createElement("th");      // TABLE HEADER.
        th.innerHTML = col[i];
        tr.appendChild(th);
    }

    // ADD JSON DATA TO THE TABLE AS ROWS.
    for (var i = 0; i < data.length-3; i++) {

        tr = table.insertRow(-1);

        for (var j = 0; j < col.length-3 ; j++) {
            var tabCell = tr.insertCell(-1);
            tabCell.innerHTML = data[i][col[j]];
        }
    }

    // FINALLY ADD THE NEWLY CREATED TABLE WITH JSON DATA TO A CONTAINER.
    let divContainer = document.getElementById("showData");
    divContainer.innerHTML = "";
    divContainer.appendChild(table);

}
var articleList = [];
//var articleList = [
    //{ FirstName: 'Byron', LastName: 'Seto' },
    //{ FirstName: 'Hi', LastName: 'Test' },
    //{ FirstName: 'Another', LastName: 'Test' }
//];

function updateList() {
    //$('#ArticleDiv').append("<ul id='article_list' class='article_list'></ul>");
    //for (cnt = 0; cnt < articleList.length; cnt++) {
        //$("#article_list").append("<li><a href='#'>"+articleList[cnt].title + ":" + articleList[cnt].create_date+"</li>");
    //}
    //
    //$('#ArticleDiv').append("<table id='article_list' class='article_list'></table>");
    //$("#article_list").append("<tr>");
    //$("#article_list").append("<th>Title</th>");
    //$("#article_list").append("<th>Date</th>");
    //for (cnt = 0; cnt < articleList.length; cnt++) {
        //$("#article_list").append("<tr>");
        //$("#article_list").append("<td>"+articleList[cnt].title + "</td>");
        //$("#article_list").append("<td>"+articleList[cnt].create_date + "</td>");
    //}
    
    //Create a HTML Table element.
    //var table = $('<table></table>').addClass('article_list');

    var table = $("<table/>").addClass('article_list');
    table.attr('id', 'article_list');

    var columnCount = 2

    var row = $("<tr/>");
    for (header in articleList[0]) {
        var headerCell = $("<th />");
        if (header !== "body") {
            console.log(articleList[0][header]);
            row.append($("<th/>").text(articleList[0][header]));
        }
        table.append(row);
    }

    for (var i = 1; i < articleList.length; i++) {
            var row = $("<tr/>");
        for (header in articleList[i]) {
            if (header !== "body") {
                console.log(articleList[i][header]);
                row.append($("<td/><a href=/home").text(articleList[i][header]));
            }
            table.append(row);
        }
    }


    //$.each(articleList, function(rowIndex, r) {
        //var row = $("<tr/>");
        //$.each(r, function(colIndex, c) { 
            //row.append($("<t"+(rowIndex == 0 ?  "h" : "d")+"/>").text(c));
        //});
        //table.append(row);
    //});
    //return container.append(table);
    
    //Get the count of columns.
    //console.log(articleList[0])
    //var columnCount = 2
    //var row = $(table[0].insertRow(-1));
    ////for (var i = 0; i < columnCount; i++) {
        ////var headerCell = $("<th />");
        ////headerCell.html(articleList[0][i]);
        ////row.append(headerCell);
    ////}

    //for (header in articleList[0]) {
        //var headerCell = $("<th />");
        //if (header !== "body") {
            //console.log(articleList[0][header]);
            //headerCell.html(articleList[0][header]);
            //console.log(headerCell);
        //}
        //row.append(headerCell);
    //}

    ////Add the data rows.
    //for (var i = 1; i < articleList.length; i++) {
        //row = $(table[0].insertRow(-1));
        //for (var j = 0; j < columnCount; j++) {
            //var cell = $("<td />");
            //cell.html(articleList[i][j]);
            //row.append(cell);
        //}
    //}

    console.log(table);
    $("#ArticleDiv").append(table);
}


function filterArticles() {
    var input, filter, ul, li, a, i;
    input = document.getElementById("article_search");
    filter = input.value.toUpperCase();
    ul = document.getElementById("article_list");
    li = ul.getElementsByTagName("li");
    for (i = 0; i < li.length; i++) {
        a = li[i].getElementsByTagName("a")[0];
        if (a.innerHTML.toUpperCase().indexOf(filter) > -1) {
            li[i].style.display = "";
        } else {
            li[i].style.display = "none";

        }
    }
}


function loadArticles() {
    var xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            console.log("Probably no return yet");
            console.log(this.responseText);
            articleList = (JSON.parse(this.responseText));
            updateList();
        }
    }
    xhttp.open("GET", "api/list_articles", true);
    xhttp.send();
}

loadArticles();

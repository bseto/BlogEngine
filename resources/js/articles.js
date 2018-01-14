var someList = [
    { FirstName: 'Byron', LastName: 'Seto' },
    { FirstName: 'Hi', LastName: 'Test' },
    { FirstName: 'Another', LastName: 'Test' }
];

function updateList() {
    $('#ArticleDiv').append("<ul id='article_list' class='article_list'></ul>");
    for (cnt = 0; cnt < someList.length; cnt++) {
        $("#article_list").append("<li><a href='#'>"+someList[cnt].FirstName + ":" + someList[cnt].LastName+"</li>");
    }
}


function addToList() {
    someList.push({FirstName: 'AnotherOne!', LastName: 'LastNameAnother!'});
    console.log("Hi");
    console.log(someList)
    updateList()
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

function testAPI() {
    var xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            console.log("Probably no return yet");
            console.log(this.responseText);
        }
    }
    xhttp.open("GET", "api/list_articles", true);
    xhttp.send();
}

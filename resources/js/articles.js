var articleList = [];
//var articleList = [
    //{ FirstName: 'Byron', LastName: 'Seto' },
    //{ FirstName: 'Hi', LastName: 'Test' },
    //{ FirstName: 'Another', LastName: 'Test' }
//];

function updateList() {
    $('#ArticleDiv').append("<ul id='article_list' class='article_list'></ul>");
    for (cnt = 0; cnt < articleList.length; cnt++) {
        $("#article_list").append("<li><a href='#'>"+articleList[cnt].title + ":" + articleList[cnt].create_date+"</li>");
    }
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
            console.log(articleList);
            updateList();
        }
    }
    xhttp.open("GET", "api/list_articles", true);
    xhttp.send();
}

loadArticles();

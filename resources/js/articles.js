var articleList = [];
//var articleList = [
    //{ FirstName: 'Byron', LastName: 'Seto' },
    //{ FirstName: 'Hi', LastName: 'Test' },
    //{ FirstName: 'Another', LastName: 'Test' }
//];

function updateList() {
    for (cnt = 0; cnt < articleList.length; cnt++) {
        var row = $("<a/>").addClass('list-group-item');
        row.attr('href', articleList[cnt].title);
        var header = $("<h4/>").addClass('list-group-item-heading');
        header.append(articleList[cnt].title);
        var date = $("<p/>").addClass('list-group-item-text');
        date.append(articleList[cnt].create_date);
        row.append(header);
        row.append(date);
        $('#ArticleDiv').append(row);
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
            updateList();
        }
    }
    xhttp.open("GET", "api/list_articles", true);
    xhttp.send();
}


loadArticles();

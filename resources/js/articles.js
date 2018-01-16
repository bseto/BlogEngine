var articleList = [];
//var articleList = [
    //{ FirstName: 'Byron', LastName: 'Seto' },
    //{ FirstName: 'Hi', LastName: 'Test' },
    //{ FirstName: 'Another', LastName: 'Test' }
//];

function updateList() {
    for (cnt = 0; cnt < articleList.length; cnt++) {
        var row = $("<a/>").addClass('list-group-item');
        row.attr('href', "article/" + articleList[cnt].path);
        var header = $("<h3/>").addClass('list-group-item-heading');
        header.append(articleList[cnt].title);
        var date = $("<p/>").addClass('list-group-item-text');
        date.append(articleList[cnt].create_date);

        var tags = $("<p/>").addClass('list-group-item-heading');
        for (tagn = 0; tagn < articleList[cnt].tags.length; tagn++) {
            var span = $("<span/>").addClass('label label-primary')
            span.append(articleList[cnt].tags[tagn]);
            tags.append(span);
            tags.append(" ");
        }
        row.append(header);
        row.append(tags);
        row.append(date);
        $('#ArticleDiv').append(row);
    }
}

function filterArticles() {
    var input, filter, div, a, p, i;
    input = document.getElementById("article_search");
    filter = input.value.toUpperCase();
    div = document.getElementById("ArticleDiv");
    a = div.getElementsByTagName("a");
    for (i = 0; i < a.length; i++) {
        p = a[i].getElementsByTagName("p")[0];
        if (p.innerHTML.toUpperCase().indexOf(filter) > -1) {
            a[i].style.display = "";
        } else {
            a[i].style.display = "none";
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

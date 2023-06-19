# Library-Manegement-System

{{ define "auth_reg" }}

<body>

    <div class="container">
        <main>
            <div class="category">
                <div class="list">
                     {{ if .Auth }}
                    <span><a href="/user/profile">Мой профиль</a></span>
                    <span><a href="/user/library">Отложенные</a></span>
                    <span><a href="/user/favourites">Закладки</a></span>
                    <hr>
                    {{ end }}
 
                    <span><a href="/books/genre/new">Новиники</a></span>
                    <span><a href="/books/genre/all">Все жанры</a></span>
                    <span><a href="/books/authors">Авторы</a></span>
                    {{ range $_, $g := .Genres }}
                    <span><a href="/books/genre/{{ $g.Link }}">{{ $g.Name }}</a></span>
                    {{ end }}
                </div>
            </div>
            <div class="main">
                <div class="link">
                    <div>
                        {{ range $_, $u := .URL_NAME }}
                        <a href='{{ index $.URLs $u }}'>{{ $u }}</a>
                        {{ end }}
                    </div>
                    <hr>
                </div>
                <h2>{{ .Auth-title }}</h2>
                <div class="books">
                </div>      
                <h2>{{ .Reg-title }}</h2>
                <div class="books">
                </div>          
            </div>
        </main>
    </div>
</body>
</html>

{{ end }}

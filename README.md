

# hashtagScraper 


This app allows you to search by some hashtag, and it returns, images and descriptions from different social media

On current version we only support instagram and twitter.

 - Clone project 
 - Import libraries and stuff (no clue on how
   :p) 
 - Run `go run main.go`

Then, you can just API call it doing a

    > GET http://localhost:8080/search/[YOUR SEARCH]

for example `http://localhost:8080/search/pokemon`

    {
      "result": [
        {
          "Name": "instagram",
          "Posts": [
            {
              "URL": "https://instagram.faep9-2.fna.fbcdn.net/vp/8ecde5141188b9a512b8211e3980218d/5C6C5BD6/t51.2885-15/e35/c30.0.480.480/50822493_119806635774214_5175769181322525196_n.jpg?_nc_ht=instagram.faep9-2.fna.fbcdn.net",
              "Text": "anyone know the song name?\n⚫\n⚫\n⚫\n#tiktok #tiktokmemes #epic #stolenmemes #meme #hitormiss #pewdiepie #christmas #tseries #furry #dankmemes #dank #funny #iwannabetracer #naruto #thanos #rainbowsixsiege #pokemon #sal #rockefellerstreet",
              "Hash": "ba10bc66b83850454b348dc0203a71ea9a76d751772841db04e1a3d2a53d92b3"
            }
	        ... <-aprox 100 latests/more trendy with this hashtag
          ]
        },
        {
          "Name": "twitter",
          "Posts": [
            {
              "URL": "",
              "Text": "I did watch sailor moon but I wasnt as into it as I was, like, Pokemon and Cardcaptor Sakura lol",
              "Hash": "6deea4131b5a1e83c7c96a43545d26f162df7bf6aca2546ebea60559f46df610"
            }
	        ... <-aprox 100 latests/more trendy with this hashtag
          ]
        }
      ]
    }
    



#### Local load tests

Series

	[GIN] 2019/02/17 - 01:18:47 | 200 |  3.731495938s |             ::1 | GET      /search/pokemon
	[GIN] 2019/02/17 - 01:20:11 | 200 |  3.382655464s |             ::1 | GET      /search/pokemon
	[GIN] 2019/02/17 - 01:20:26 | 200 |  2.349863251s |             ::1 | GET      /search/pokemon

Parallel

	[GIN] 2019/02/17 - 01:51:51 | 200 |  2.112491575s |             ::1 | GET      /search/pokemon
	[GIN] 2019/02/17 - 01:51:54 | 200 |  1.638974595s |             ::1 | GET      /search/pokemon
	[GIN] 2019/02/17 - 01:51:56 | 200 |  1.731282873s |             ::1 | GET      /search/pokemon

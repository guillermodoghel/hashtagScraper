


<p align="center">
<img 
    src="https://i.imgur.com/LAz2B0z.png" 
     height="120">
<br>
</p>

![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

# hashtagScraper 


This app allows you to search by some hashtag, and it returns, images and descriptions from different social media

On current version we only support instagram and twitter.

 - Clone project 
 - Import libraries and stuff (no clue on how,
    try following this steps https://golang.org/doc/install and may the force be with you :p)
 - Edit **settings.properties** , this file will contain your twitter api credentials and your instagram account login data. 
 - Now you are ready to go, just run `go run main.go`

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
		"Name": "instagram_stories",
		"Posts": [
		  {
			"URL": "https://instagram.faep8-1.fna.fbcdn.net/vp/332497378f2c3ce5cc115eb55bbb37b2/5C6E0E7A/t51.12442-15/e35/51808015_392768001517608_7771288113720066240_n.jpg?_nc_ht=instagram.faep8-1.fna.fbcdn.net",
			"Text": "",
			"Hash": "d7a0b26d0159a1b09730b00385d2b2ba99c1c659ea180b48fec1e97f82ccc6e5"
		  },
		  {
			"URL": "https://instagram.faep8-1.fna.fbcdn.net/vp/99b46b6323cc56328ee7a6455a8150f6/5C6D9931/t51.12442-15/e35/50107738_309557969700767_2586428198555859722_n.jpg?_nc_ht=instagram.faep8-1.fna.fbcdn.net",
			"Text": "",
			"Hash": "7853e56c81224f709590547c2ee4c42099631bbeff8f7f079ec780f30b07f880"
		  },
		  ... <-aprox 20 latests/more trendy with this hashtag
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
    


#### Now with full parallel scrapping

	Twitter took 1.981979526s
	Instagram took 4.928551827s
	Instagram stories took 6.15028761s
	Scrapping took 6.151213822s
	[GIN] 2019/02/19 - 00:48:00 | 200 |  6.151229098s |             ::1 | GET      /search/pokemon



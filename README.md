


<p align="center">
<img 
    src="https://i.imgur.com/LAz2B0z.png" 
     height="120">
<br>
</p>

# hashtagScraper 


This app allows you to search by some hashtag, and it returns, images and descriptions from different social media

On current version we only support instagram and twitter.

 - Clone project 
 - Import libraries and stuff (no clue on how,
    try following this steps https://golang.org/doc/install and may the force be with you :p)
 - Create at project root level a text file named **settings.properties** , this file will contain your twitter api credentials. The .gitignore is configured to avoid the commit of this file for obvious reasons. (If you search previous commits on this repo you will find one example with voided keys)  It must have the following format:
 

>      twitterConsumerKey = 123asdf
>      twitterConsumerSecret = 123asdf
>      twitterAccessToken = 123asdf 
>      twitterAccessSecret = 123asdf

 
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

      Twitter took 2.468745525s
      Instagram took 2.551599415s
      Scrapping took 2.552257732s
      [GIN] 2019/02/18 - 16:14:33 | 200 |   2.55226692s |             ::1 | GET      /search/pokemon



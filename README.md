

# hashtagScraper 


This app allows you to search by some hashtag, and it returns, images and descriptions from different social media

On current version we only support instagram. Next: twitter, facebook, and linkedin

 - Clone project 
 - Import libraries and stuff (no clue on how
   :p) 
   Run `go run main.go`

Then, you can just API call it doing a

    > GET http://localhost:8080/search/[YOUR SEARCH]

for example `http://localhost:8080/search/pokemon`


    {
	    "result": [
		    {
				"Name": "instagram",
				"Posts": [
				    {
					    "URL": "https://instagram.faep8-2.fna.fbcdn.net/vp/3ffbed111ee726ea92c2e02addf765b8/5CE226C7/t51.2885-15/sh0.08/e35/s640x640/50618298_353111202197343_1404184570313895034_n.jpg?_nc_ht=instagram.faep8-2.fna.fbcdn.net",
					    "Text": "Even though we are all sold out of the 1V1 and the 2V2 tournament entries. There are still a few Spectator tickets available, which come with an unlimited Gameplay card to the Dave and Buster's Arcade and also the chance to get added at random to the tournament! These are going fast too! Don’t miss out!\n.\n.\n______________________________________\n@wyd_gaming @wyd_gaming\n.\n.\n______________________________________\n\n#esportsla #laesports #losangelesgamers\n#losangelesesports #hollywoodesports\n#esports\n#tournamentkingz\n#wyd\n#wydgaming\n#gamersonly #gamer\n#gaming #instagaming #instagamer\n #photooftheday\n#onlinegaming\n#videogameaddict\n#instagame\n#gamerlife\n#smashbros\n#ssb5\n#smashbrosultimate #smash5 \n#pokemon #fortnite #videogamer #inkling #inklinggirl #2days #fun"
				    }
				]
			},
		    {...},
   		    {...}
	   	]		
  	}
